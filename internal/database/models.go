package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"userId"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Event struct {
	EventID      uuid.UUID `json:"eventId"`
	EventName    string    `json:"eventName"`
	TotalQuota   uint64    `json:"totalQuota"`
	AvailableQty uint64    `json:"availableQty"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Order struct {
	OrderID   uuid.UUID `json:"orderId"`
	EventID   uuid.UUID `json:"eventId"`
	UserID    uuid.UUID `json:"userId"`
	Amount    uint16    `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type Ticket struct {
	TicketID  uuid.UUID `json:"ticketId"`
	OrderID   uuid.UUID `json:"orderId"`
	UserID    uuid.UUID `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
