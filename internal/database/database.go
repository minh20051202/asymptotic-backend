package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(*User) error
	DeleteUserById(uuid.UUID) error
	UpdateUser(*User) error
	GetAllUsers() ([]*User, error)
	GetUserById(uuid.UUID) (*User, error)

	CreateEvent(*Event) error
	DeleteEventById(uuid.UUID) error
	UpdateEvent(*Event) error
	GetAllEvents() ([]*Event, error)
	GetEventById(uuid.UUID) (*Event, error)

	CreateOrder(*Order) error
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
	return &PostgresStore{
		db: db,
	}, nil
}

func (ps *PostgresStore) Init() error {
	if err := ps.createUserTable(); err != nil {
		return err
	}
	if err := ps.createEventTable(); err != nil {
		return err
	}
	if err := ps.createOrderTable(); err != nil {
		return err
	}
	if err := ps.createTicketTable(); err != nil {
		return err
	}
	return nil
}
func (ps *PostgresStore) createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		user_id UUID DEFAULT gen_random_uuid(),
		username VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (user_id)
	)`
	_, err := ps.db.Exec(query)
	return err
}

func (ps *PostgresStore) createEventTable() error {
	query := `CREATE TABLE IF NOT EXISTS events (
		event_id UUID DEFAULT gen_random_uuid(),
		event_name VARCHAR(255) NOT NULL,
		total_quota INT NOT NULL,
		available_qty INT NOT NULL CHECK(available_qty >= 0),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (event_id)
	)`
	_, err := ps.db.Exec(query)
	return err
}
func (ps *PostgresStore) createOrderTable() error {
	query := `CREATE TABLE IF NOT EXISTS orders (
		order_id UUID DEFAULT gen_random_uuid(),
		event_id UUID,
		user_id UUID,
		amount SMALLINT NOT NULL CHECK(amount > 0),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY(order_id),
		CONSTRAINT fk_order_event
			FOREIGN KEY (event_id)
				REFERENCES events(event_id)
					ON DELETE RESTRICT,
		CONSTRAINT fk_order_user
			FOREIGN KEY (user_id)
				REFERENCES users(user_id)
					ON DELETE RESTRICT
					
	)`
	_, err := ps.db.Exec(query)
	return err
}

func (ps *PostgresStore) createTicketTable() error {
	query := `CREATE TABLE IF NOT EXISTS tickets (
		ticket_id UUID DEFAULT gen_random_uuid(),
		user_id UUID,
		order_id UUID,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(ticket_id),

		CONSTRAINT fk_ticket_user
			FOREIGN KEY (user_id)
				REFERENCES users(user_id)
					ON DELETE RESTRICT,
		CONSTRAINT fk_ticket_order
			FOREIGN KEY (order_id)
				REFERENCES orders(order_id)
					ON DELETE RESTRICT
	)`
	_, err := ps.db.Exec(query)
	return err
}

func (ps *PostgresStore) CreateUser(user *User) error {
	query := `INSERT INTO users
	(user_id, username, password, created_at)
	values($1, $2, $3, $4)`

	_, err := ps.db.Query(
		query,
		user.UserID,
		user.Username,
		user.Password,
		user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ps *PostgresStore) UpdateUser(user *User) error {
	return nil
}
func (ps *PostgresStore) GetAllUsers() ([]*User, error) {
	rows, err := ps.db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

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
func (ps *PostgresStore) GetUserById(uuid uuid.UUID) (*User, error) {
	rows, err := ps.db.Query("SELECT * FROM users WHERE user_id = $1", uuid)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUsers(rows)
	}

	return nil, fmt.Errorf("User %d not found", uuid)
}

func (ps *PostgresStore) DeleteUserById(uuid uuid.UUID) error {
	return nil
}

func scanIntoUsers(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(
		&user.UserID,
		&user.Username,
		&user.Password,
		&user.CreatedAt)
	return user, err
}

func (ps *PostgresStore) CreateEvent(event *Event) error {
	query := `INSERT INTO events
	(event_id,event_name, total_quota, available_qty, created_at)
	values($1, $2, $3, $4, $5)`

	_, err := ps.db.Query(
		query,
		event.EventID,
		event.EventName,
		event.TotalQuota,
		event.AvailableQty,
		event.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ps *PostgresStore) GetAllEvents() ([]*Event, error) {
	rows, err := ps.db.Query("SELECT * FROM events")

	if err != nil {
		return nil, err
	}

	events := []*Event{}
	for rows.Next() {
		event, err := scanIntoEvents(rows)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
func (ps *PostgresStore) GetEventById(uuid uuid.UUID) (*Event, error) {
	rows, err := ps.db.Query("SELECT * FROM events WHERE event_id = $1", uuid)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoEvents(rows)
	}

	return nil, fmt.Errorf("Event %d not found", uuid)
}

func (ps *PostgresStore) UpdateEvent(event *Event) error {
	return nil
}

func (ps *PostgresStore) DeleteEventById(uuid uuid.UUID) error {
	return nil
}

func scanIntoEvents(rows *sql.Rows) (*Event, error) {
	event := new(Event)
	err := rows.Scan(
		&event.EventID,
		&event.EventName,
		&event.TotalQuota,
		&event.AvailableQty,
		&event.CreatedAt)
	return event, err
}

func (ps *PostgresStore) CreateOrder(order *Order) error {
	tx, err := ps.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	queryEvent := `
        UPDATE events 
        SET available_qty = available_qty - $1 
        WHERE event_id = $2
    `
	_, err = tx.Exec(queryEvent, order.Amount, order.EventID)
	if err != nil {
		return err
	}

	queryOrder := `
		INSERT INTO orders (order_id, event_id, user_id, amount, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = tx.Exec(queryOrder, order.OrderID, order.EventID, order.UserID, order.Amount, order.CreatedAt)
	if err != nil {
		return err
	}

	queryTicket := `
        INSERT INTO tickets (ticket_id, order_id, user_id) 
        VALUES ($1, $2, $3)
    `
	for range order.Amount {
		ticketID := uuid.New()
		_, err = tx.Exec(queryTicket, ticketID, order.OrderID, order.UserID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
