package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"uuid"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Event struct {
	EventID      uuid.UUID
	TotalQuota   uint64
	AvailableQty uint64
	CreatedAt    time.Time
}

type Order struct {
	OrderID   uuid.UUID
	EventID   uuid.UUID
	UserID    uuid.UUID
	Amount    uint16
	CreatedAt time.Time
}

type Ticket struct {
	TicketID  uuid.UUID
	OrderID   uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
}
