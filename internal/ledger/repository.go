package ledger

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/minh20051202/ticket-system-backend/internal/shared"
)

var ErrInsufficientFunds = errors.New("insufficient funds")
var ErrAmountNotGreaterThanZero = errors.New("amount not greater than 0")

type LedgerRepository interface {
	Init() error

	GetBalanceById(uuid.UUID) (*shared.Balance, error)

	Charge(*shared.Transaction) (*shared.Transaction, error)
	Deposit(*shared.Transaction) (*shared.Transaction, error)
	UpdateTransactionStatus(uuid.UUID, string) error
	GetAllTransactions() ([]*shared.Transaction, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Init() error {
	if err := r.createBalanceTable(); err != nil {
		return err
	}
	if err := r.createTransactionTable(); err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) createBalanceTable() error {
	query := `CREATE TABLE IF NOT EXISTS balances (
        user_id UUID PRIMARY KEY,
        balance BIGINT DEFAULT 0 CHECK(balance >= 0),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_balance_user
            FOREIGN KEY (user_id)
                REFERENCES users(user_id)
                    ON DELETE RESTRICT
    )`
	_, err := r.db.Exec(query)
	return err
}

func (r *PostgresRepository) createTransactionTable() error {
	query := `CREATE TABLE IF NOT EXISTS transactions (
        transaction_id UUID PRIMARY KEY,
        user_id UUID NOT NULL,
        idempotency_key VARCHAR(255) UNIQUE NOT NULL,
        amount BIGINT NOT NULL,
		type VARCHAR(20) NOT NULL CHECK (type IN ('CHARGE', 'DEPOSIT')),
        status VARCHAR(20) NOT NULL CHECK (status IN ('PENDING', 'FAILED', 'SUCCEEDED')) DEFAULT 'PENDING', 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

        CONSTRAINT fk_transaction_user
            FOREIGN KEY (user_id)
                REFERENCES users(user_id)
                    ON DELETE RESTRICT
    )`
	_, err := r.db.Exec(query)
	return err
}

func (r *PostgresRepository) GetBalanceById(uuid uuid.UUID) (*shared.Balance, error) {
	rows, err := r.db.Query("SELECT * FROM balances WHERE user_id = $1", uuid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanIntoBalances(rows)
	}

	return nil, fmt.Errorf("User %v not found", uuid)
}

func scanIntoBalances(rows *sql.Rows) (*shared.Balance, error) {
	balance := new(shared.Balance)
	err := rows.Scan(
		&balance.UserId,
		&balance.Balance,
		&balance.CreatedAt,
	)
	return balance, err
}

func (r *PostgresRepository) Charge(transaction *shared.Transaction) (*shared.Transaction, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if transaction.Amount <= 0 {
		return nil, ErrAmountNotGreaterThanZero
	}

	queryTransaction := `
		INSERT INTO transactions (transaction_id, user_id, idempotency_key, amount, type, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (idempotency_key) DO NOTHING
	`

	result, err := tx.Exec(queryTransaction, transaction.TransactionId, transaction.UserId, transaction.IdempotencyKey, transaction.Amount, transaction.Type, transaction.CreatedAt)
	if err != nil {
		return nil, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowAffected == 0 {
		oldTransaction := &shared.Transaction{}
		queryRead := `SELECT transaction_id, user_id, idempotency_key, amount, type, status, created_at FROM transactions WHERE idempotency_key = $1`
		err = tx.QueryRow(queryRead, transaction.IdempotencyKey).Scan(&oldTransaction.TransactionId, &oldTransaction.UserId, &oldTransaction.IdempotencyKey, &oldTransaction.Amount, &oldTransaction.Type, &oldTransaction.Status, &oldTransaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		return oldTransaction, nil
	}

	var balance int64
	queryRead := `SELECT balance FROM balances WHERE user_id = $1 FOR UPDATE`

	err = tx.QueryRow(queryRead, transaction.UserId).Scan(&balance)
	if err != nil {
		return nil, err
	}

	if balance < int64(transaction.Amount) {
		return nil, ErrInsufficientFunds
	}

	newBalance := balance - int64(transaction.Amount)

	queryUpdate := `
        UPDATE balances 
        SET balance = $1
        WHERE user_id = $2
    `
	_, err = tx.Exec(queryUpdate, newBalance, transaction.UserId)
	if err != nil {
		return nil, err
	}

	transaction.Status = "PENDING"

	return transaction, tx.Commit()
}

func (r *PostgresRepository) Deposit(transaction *shared.Transaction) (*shared.Transaction, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if transaction.Amount <= 0 {
		return nil, ErrAmountNotGreaterThanZero
	}

	queryTransaction := `
		INSERT INTO transactions (transaction_id, user_id, idempotency_key, amount, type, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (idempotency_key) DO NOTHING
	`

	result, err := tx.Exec(queryTransaction, transaction.TransactionId, transaction.UserId, transaction.IdempotencyKey, transaction.Amount, transaction.Type, transaction.CreatedAt)
	if err != nil {
		return nil, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowAffected == 0 {
		oldTransaction := &shared.Transaction{}
		queryRead := `SELECT transaction_id, user_id, idempotency_key, amount, type, status, created_at FROM transactions WHERE idempotency_key = $1`
		err = tx.QueryRow(queryRead, transaction.IdempotencyKey).Scan(&oldTransaction.TransactionId, &oldTransaction.UserId, &oldTransaction.IdempotencyKey, &oldTransaction.Amount, &oldTransaction.Type, &oldTransaction.Status, &oldTransaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		return oldTransaction, nil
	}

	var balance int64
	queryRead := `SELECT balance FROM balances WHERE user_id = $1 FOR UPDATE`

	err = tx.QueryRow(queryRead, transaction.UserId).Scan(&balance)
	if err != nil {
		return nil, err
	}

	newBalance := balance + int64(transaction.Amount)

	queryUpdate := `
        UPDATE balances 
        SET balance = $1
        WHERE user_id = $2
    `
	_, err = tx.Exec(queryUpdate, newBalance, transaction.UserId)
	if err != nil {
		return nil, err
	}

	transaction.Status = "PENDING"

	return transaction, tx.Commit()
}

func (r *PostgresRepository) UpdateTransactionStatus(txId uuid.UUID, status string) error {
	query := `UPDATE transactions SET status = $1 WHERE transaction_id = $2`
	_, err := r.db.Exec(query, status, txId)
	return err
}

func (r *PostgresRepository) GetAllTransactions() ([]*shared.Transaction, error) {
	rows, err := r.db.Query("SELECT * FROM transactions")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []*shared.Transaction{}
	for rows.Next() {
		transaction, err := scanIntoTransactions(rows)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func scanIntoTransactions(rows *sql.Rows) (*shared.Transaction, error) {
	transaction := new(shared.Transaction)
	err := rows.Scan(
		&transaction.TransactionId,
		&transaction.UserId,
		&transaction.IdempotencyKey,
		&transaction.Amount,
		&transaction.Status,
		&transaction.CreatedAt)
	return transaction, err
}
