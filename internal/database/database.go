package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/minh20051202/ticket-system-backend/internal/shared"
)

var ErrInsufficientFunds = errors.New("insufficient funds")
var ErrAmountNotGreaterThanZero = errors.New("amount not greater than 0")

type Storage interface {
	CreateUserWithBalance(*shared.User) error
	UpdateUser(*shared.User) error
	GetAllUsers() ([]*shared.User, error)
	GetUserById(uuid.UUID) (*shared.User, error)

	GetBalanceById(uuid.UUID) (*shared.Balance, error)

	Charge(*shared.Transaction) (*shared.Transaction, error)
	Deposit(*shared.Transaction) (*shared.Transaction, error)
	GetAllTransactions() ([]*shared.Transaction, error)
}

type PostgresStore struct {
	db *sql.DB
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() (*PostgresStore, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh", host, username, password, database, port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return &PostgresStore{
		db: db,
	}, nil
}

func (ps *PostgresStore) Init() error {
	if err := ps.createUserTable(); err != nil {
		return err
	}
	if err := ps.createBalanceTable(); err != nil {
		return err
	}
	if err := ps.createTransactionTable(); err != nil {
		return err
	}
	if err := ps.createApiKeyTable(); err != nil {
		return err
	}
	return nil
}

func (ps *PostgresStore) Close() error {
	return ps.db.Close()
}

func (ps *PostgresStore) createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
        user_id UUID PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`
	_, err := ps.db.Exec(query)
	return err
}

func (ps *PostgresStore) createBalanceTable() error {
	query := `CREATE TABLE IF NOT EXISTS balances (
        user_id UUID PRIMARY KEY,
        balance BIGINT DEFAULT 0 CHECK(balance >= 0),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_balance_user
            FOREIGN KEY (user_id)
                REFERENCES users(user_id)
                    ON DELETE RESTRICT
    )`
	_, err := ps.db.Exec(query)
	return err
}

func (ps *PostgresStore) createTransactionTable() error {
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
	_, err := ps.db.Exec(query)
	return err
}

func (ps *PostgresStore) createApiKeyTable() error {
	query := `CREATE TABLE IF NOT EXISTS api_keys (
        api_key VARCHAR(255) PRIMARY KEY,
        user_id UUID NOT NULL,
        name VARCHAR(50) NOT NULL, 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_apikey_user
            FOREIGN KEY (user_id)
                REFERENCES users(user_id)
                    ON DELETE RESTRICT
    )`
	_, err := ps.db.Exec(query)
	return err
}

func (ps *PostgresStore) CreateUserWithBalance(user *shared.User) error {
	tx, err := ps.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	userQuery := `
		INSERT INTO users (user_id, username, email, password, created_at) 
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = tx.Exec(userQuery, user.UserId, user.Username, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}

	balanceQuery := `
		INSERT INTO balances (user_id, balance, created_at) 
		VALUES ($1, $2, $3)
	`

	_, err = tx.Exec(balanceQuery, user.UserId, 0, user.CreatedAt)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (ps *PostgresStore) UpdateUser(user *shared.User) error {
	return nil
}

func (ps *PostgresStore) GetAllUsers() ([]*shared.User, error) {
	rows, err := ps.db.Query("SELECT user_id, username, email, password, created_at FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*shared.User{}
	for rows.Next() {
		user, err := scanIntoUsers(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ps *PostgresStore) GetUserById(uuid uuid.UUID) (*shared.User, error) {
	rows, err := ps.db.Query("SELECT user_id, username, email, password, created_at FROM users WHERE user_id = $1", uuid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanIntoUsers(rows)
	}

	return nil, fmt.Errorf("User %v not found", uuid)
}

func scanIntoUsers(rows *sql.Rows) (*shared.User, error) {
	user := new(shared.User)
	err := rows.Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (ps *PostgresStore) GetBalanceById(uuid uuid.UUID) (*shared.Balance, error) {
	rows, err := ps.db.Query("SELECT * FROM balances WHERE user_id = $1", uuid)

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

func (ps *PostgresStore) Charge(transaction *shared.Transaction) (*shared.Transaction, error) {
	tx, err := ps.db.Begin()
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

func (ps *PostgresStore) Deposit(transaction *shared.Transaction) (*shared.Transaction, error) {
	tx, err := ps.db.Begin()
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

func (ps *PostgresStore) GetAllTransactions() ([]*shared.Transaction, error) {
	rows, err := ps.db.Query("SELECT * FROM transactions")

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

func (ps *PostgresStore) GetUserIdByApiKey(apiKey string) (uuid.UUID, error) {
	var userId uuid.UUID

	query := `SELECT user_id FROM api_keys WHERE api_key = $1`

	err := ps.db.QueryRow(query, apiKey).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return uuid.Nil, fmt.Errorf("invalid API key")
		}
		return uuid.Nil, err
	}

	return userId, nil
}
