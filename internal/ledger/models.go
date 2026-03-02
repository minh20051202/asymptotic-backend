package ledger

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	UserId    uuid.UUID `json:"userId"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type Transaction struct {
	TransactionId  uuid.UUID `json:"transactionId"`
	UserId         uuid.UUID `json:"userId"`
	IdempotencyKey string    `json:"idempotencyKey"`
	Amount         int64     `json:"amount"`
	Type           string    `json:"type"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
}

type CreateTransactionRequest struct {
	UserId         uuid.UUID `json:"userId"`
	IdempotencyKey string    `json:"idempotencyKey"`
	Amount         int64     `json:"amount"`
	Type           string    `json:"type"`
}
