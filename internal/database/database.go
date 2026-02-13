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

type Storage interface {
	CreateUser(*shared.User) error
	DeleteUserById(uuid.UUID) error
	UpdateUser(*shared.User) error
	GetAllUsers() ([]*shared.User, error)
	GetUserById(uuid.UUID) (*shared.User, error)

	CreateWallet(*shared.Wallet) error
	DeleteWalletById(uuid.UUID) error
	UpdateWallet(*shared.Wallet) error
	GetAllWallets() ([]*shared.Wallet, error)
	GetWalletById(uuid.UUID) (*shared.Wallet, error)

	Charge(*shared.Transaction) (*shared.Transaction, error)
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
	if err := ps.createWalletTable(); err != nil {
		return err
	}
	if err := ps.createTransactionTable(); err != nil {
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

func (ps *PostgresStore) createWalletTable() error {
	query := `CREATE TABLE IF NOT EXISTS wallets (
        wallet_id UUID PRIMARY KEY,
        user_id UUID NOT NULL,
        balance BIGINT DEFAULT 0 CHECK(balance >= 0),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_wallet_user
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
        wallet_id UUID NOT NULL,
        user_id UUID NOT NULL,
        idempotency_key VARCHAR(255) UNIQUE NOT NULL,
        amount BIGINT NOT NULL,
        status VARCHAR(20) NOT NULL CHECK (status IN ('PENDING', 'FAILED', 'SUCCEEDED')) DEFAULT 'PENDING', 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

        CONSTRAINT fk_transaction_wallet
            FOREIGN KEY (wallet_id)
                REFERENCES wallets(wallet_id)
                    ON DELETE RESTRICT,
        CONSTRAINT fk_transaction_user
            FOREIGN KEY (user_id)
                REFERENCES users(user_id)
                    ON DELETE RESTRICT
    )`
	_, err := ps.db.Exec(query)
	return err
}

func (ps *PostgresStore) CreateUser(user *shared.User) error {
	query := `INSERT INTO users
	(user_id, username, email, password, created_at)
	values($1, $2, $3, $4, $5)`

	_, err := ps.db.Exec(
		query,
		user.UserID,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ps *PostgresStore) UpdateUser(user *shared.User) error {
	return nil
}

func (ps *PostgresStore) GetAllUsers() ([]*shared.User, error) {
	rows, err := ps.db.Query("SELECT * FROM users")

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
	rows, err := ps.db.Query("SELECT * FROM users WHERE user_id = $1", uuid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanIntoUsers(rows)
	}

	return nil, fmt.Errorf("User %d not found", uuid)
}

func (ps *PostgresStore) DeleteUserById(uuid uuid.UUID) error {
	_, err := ps.db.Exec("DELETE FROM users WHERE user_id = $1", uuid)

	if err != nil {
		return err
	}

	return nil
}

func scanIntoUsers(rows *sql.Rows) (*shared.User, error) {
	user := new(shared.User)
	err := rows.Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (ps *PostgresStore) CreateWallet(wallet *shared.Wallet) error {
	query := `INSERT INTO wallets
	(wallet_id, user_id, created_at)
	values($1, $2, $3)`

	_, err := ps.db.Exec(
		query,
		wallet.WalletID,
		wallet.UserId,
		wallet.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ps *PostgresStore) GetAllWallets() ([]*shared.Wallet, error) {
	rows, err := ps.db.Query("SELECT * FROM wallets")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	wallets := []*shared.Wallet{}
	for rows.Next() {
		wallet, err := scanIntoWallets(rows)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}
func (ps *PostgresStore) GetWalletById(uuid uuid.UUID) (*shared.Wallet, error) {
	rows, err := ps.db.Query("SELECT * FROM wallets WHERE wallet_id = $1", uuid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanIntoWallets(rows)
	}

	return nil, fmt.Errorf("Wallet %d not found", uuid)
}

func (ps *PostgresStore) UpdateWallet(wallet *shared.Wallet) error {
	return nil
}

func (ps *PostgresStore) DeleteWalletById(uuid uuid.UUID) error {
	return nil
}

func scanIntoWallets(rows *sql.Rows) (*shared.Wallet, error) {
	wallet := new(shared.Wallet)
	err := rows.Scan(
		&wallet.WalletID,
		&wallet.UserId,
		&wallet.Balance,
		&wallet.CreatedAt,
	)
	return wallet, err
}

func (ps *PostgresStore) Charge(transaction *shared.Transaction) (*shared.Transaction, error) {
	tx, err := ps.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	queryTransaction := `
		INSERT INTO transactions (transaction_id, wallet_id, user_id, idempotency_key, amount, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (idempotency_key) DO NOTHING
	`

	result, err := tx.Exec(queryTransaction, transaction.TransactionID, transaction.WalletID, transaction.UserID, transaction.IdempotencyKey, transaction.Amount, transaction.CreatedAt)
	if err != nil {
		return nil, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowAffected == 0 {
		oldTransaction := &shared.Transaction{}
		queryRead := `SELECT transaction_id, wallet_id, user_id, idempotency_key, amount, status, created_at FROM transactions WHERE idempotency_key = $1`
		err = tx.QueryRow(queryRead, transaction.IdempotencyKey).Scan(&oldTransaction.TransactionID, &oldTransaction.WalletID, &oldTransaction.UserID, &oldTransaction.IdempotencyKey, &oldTransaction.Amount, &oldTransaction.Status, &oldTransaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		return oldTransaction, nil
	}

	var balance int64
	queryRead := `SELECT balance FROM wallets WHERE wallet_id = $1 FOR UPDATE`

	err = tx.QueryRow(queryRead, transaction.WalletID).Scan(&balance)
	if err != nil {
		return nil, err
	}

	if balance < int64(transaction.Amount) {
		return nil, ErrInsufficientFunds
	}

	newBalance := balance - int64(transaction.Amount)

	queryUpdate := `
        UPDATE wallets 
        SET balance = $1
        WHERE wallet_id = $2
    `
	_, err = tx.Exec(queryUpdate, newBalance, transaction.WalletID)
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
		&transaction.TransactionID,
		&transaction.WalletID,
		&transaction.UserID,
		&transaction.IdempotencyKey,
		&transaction.Amount,
		&transaction.Status,
		&transaction.CreatedAt)
	return transaction, err
}
