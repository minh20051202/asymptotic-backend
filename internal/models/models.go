package models

import "github.com/google/uuid"

type User struct {
	UserID   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:user_id"`
	Username string    `gorm:"type:varchar(50);unique;not null;column:username"`
	Password string    `gorm:"type:varchar(255);not null;column:password"`
}

type Event struct {
	EventID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:event_id"`
	TotalQuota   int       `gorm:"type:int;not null;column:total_quota"`
	AvailableQty int       `gorm:"type:int;not null;check:available_qty >= 0;column:available_qty"`
}

type Order struct {
	OrderID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:order_id"`
	EventID uuid.UUID `gorm:"type:uuid;not null;column:event_id"`
	UserID  uuid.UUID `gorm:"type:uuid;not null;column:user_id"`
	Amount  int16     `gorm:"type:smallint;not null;column:amount"`

	Event Event `gorm:"foreignKey:EventID;references:EventID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	User  User  `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type Ticket struct {
	TicketID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:ticket_id"`
	UserID   uuid.UUID `gorm:"type:uuid;not null;column:user_id"`
	OrderID  uuid.UUID `gorm:"type:uuid;not null;column:order_id"`

	User  User  `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Order Order `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
