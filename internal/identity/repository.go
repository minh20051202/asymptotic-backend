package identity

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type IdentityRepository interface {
	Init() error

	CreateUserWithBalance(user *User) error
	UpdateUser(user *User) error
	GetAllUsers() ([]*User, error)
	GetUserById(uuid uuid.UUID) (*User, error)
	GetUserByUsername(username string) (*User, error)

	CreateApiKey(apiKey *ApiKey) error
	GetUserIdByApiKeyHash(apiKeyHash string) (uuid.UUID, error)
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
	if err := r.createUserTable(); err != nil {
		return err
	}
	if err := r.createApiKeyTable(); err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
        user_id UUID PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
		role VARCHAR(20) CHECK (role IN ('USER', 'ADMIN')) DEFAULT 'USER',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) createApiKeyTable() error {
	query := `CREATE TABLE IF NOT EXISTS api_keys (
        api_key VARCHAR(255) PRIMARY KEY,
        user_id UUID NOT NULL,
        name VARCHAR(50) NOT NULL, 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(user_id, name),
        CONSTRAINT fk_apikey_user
            FOREIGN KEY (user_id)
                REFERENCES users(user_id)
                    ON DELETE RESTRICT
    )`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) CreateUserWithBalance(user *User) error {
	tx, err := r.db.Begin()
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

func (r *PostgresRepository) UpdateUser(user *User) error {
	return nil
}

func (r *PostgresRepository) GetAllUsers() ([]*User, error) {
	rows, err := r.db.Query("SELECT user_id, username, email, password, created_at FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		user, err := scanIntoUsers(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *PostgresRepository) GetUserById(uuid uuid.UUID) (*User, error) {
	user := new(User)

	query := `SELECT user_id, username, email, password, created_at FROM users WHERE user_id = $1`

	err := r.db.QueryRow(query, uuid).Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user %s not found", uuid)
		}
		return nil, err
	}

	return user, nil
}

func (r *PostgresRepository) GetUserByUsername(username string) (*User, error) {
	user := new(User)

	query := `SELECT user_id, username, email, password, created_at FROM users WHERE username = $1`

	err := r.db.QueryRow(query, username).Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user %s not found", username)
		}
		return nil, err
	}

	return user, nil
}

func (r *PostgresRepository) CreateApiKey(apiKey *ApiKey) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	queryApiKey := `
		INSERT INTO api_keys(api_key, user_id, name, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err = tx.Exec(queryApiKey, apiKey.ApiKey, apiKey.UserId, apiKey.Name, apiKey.CreatedAt)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresRepository) GetUserIdByApiKeyHash(apiKeyHash string) (uuid.UUID, error) {
	var userId uuid.UUID

	query := `SELECT user_id FROM api_keys WHERE api_key = $1`

	err := r.db.QueryRow(query, apiKeyHash).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return uuid.Nil, fmt.Errorf("invalid API key")
		}
		return uuid.Nil, err
	}

	return userId, nil
}

func scanIntoUsers(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
