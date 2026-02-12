package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"userId"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Wallet struct {
	WalletID  uuid.UUID `json:"walletId"`
	UserId    uuid.UUID `json:"userId"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type Transaction struct {
	TransactionID  uuid.UUID `json:"transactionId"`
	WalletID       uuid.UUID `json:"walletId"`
	UserID         uuid.UUID `json:"userId"`
	IdempotencyKey string    `json:"idempotencyKey"`
	Amount         int64     `json:"amount"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
}
