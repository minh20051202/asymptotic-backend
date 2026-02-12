package server

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateWalletRequest struct {
	WalletName   string `json:"walletName"`
	TotalQuota   uint64 `json:"totalQuota"`
	AvailableQty uint64 `json:"availableQty"`
}

type CreateTransactionRequest struct {
	WalletID uuid.UUID `json:"walletId"`
	UserID   uuid.UUID `json:"userId"`
	Amount   uint16    `json:"amount"`
}

type User struct {
	UserID    uuid.UUID `json:"userId"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Wallet struct {
	WalletID     uuid.UUID `json:"walletId"`
	WalletName   string    `json:"walletName"`
	TotalQuota   uint64    `json:"totalQuota"`
	AvailableQty uint64    `json:"availableQty"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Transaction struct {
	TransactionID uuid.UUID `json:"transactionId"`
	WalletID      uuid.UUID `json:"walletId"`
	UserID        uuid.UUID `json:"userId"`
	Amount        uint16    `json:"amount"`
	CreatedAt     time.Time `json:"createdAt"`
}

type Ticket struct {
	TicketID      uuid.UUID `json:"ticketId"`
	TransactionID uuid.UUID `json:"transactionId"`
	UserID        uuid.UUID `json:"userId"`
	CreatedAt     time.Time `json:"createdAt"`
}
