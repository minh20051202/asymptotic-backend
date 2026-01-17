package server

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateEventRequest struct {
	EventName    string `json:"eventName"`
	TotalQuota   uint64 `json:"totalQuota"`
	AvailableQty uint64 `json:"availableQty"`
}

type CreateOrderRequest struct {
	EventID uuid.UUID `json:"eventId"`
	UserID  uuid.UUID `json:"userId"`
	Amount  uint16    `json:"amount"`
}
