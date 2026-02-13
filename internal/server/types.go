package server

import (
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateWalletRequest struct {
	UserId  uuid.UUID `json:"userId"`
	Balance int64     `json:"Balance"`
}

type CreateTransactionRequest struct {
	WalletID       uuid.UUID `json:"walletId"`
	UserID         uuid.UUID `json:"userId"`
	IdempotencyKey string    `json:"idempotencyKey"`
	Amount         int64     `json:"amount"`
}
