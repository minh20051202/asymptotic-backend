package server

import (
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	JWT      string `json:"jwt"`
}

type CreateTransactionRequest struct {
	UserId         uuid.UUID `json:"userId"`
	IdempotencyKey string    `json:"idempotencyKey"`
	Amount         int64     `json:"amount"`
	Type           string    `json:"type"`
}

type CreateApiKeyRequest struct {
	Name string `json:"name"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateApiKeyResponse struct {
	ApiKey string `json:"apiKey"`
}
