package shared

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId    uuid.UUID `json:"userId"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

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

type ApiKey struct {
	ApiKey    string    `json:"apiKey"`
	UserId    uuid.UUID `json:"userId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
