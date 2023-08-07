package storage

import (
	"payments/internal/domain/entities"
	"time"
)

type passengerModel struct {
	ID          string `gorm:"primaryKey;not null"`
	Gmail       string `gorm:"not null"`
	FullName    string `gorm:"not null"`
	PhoneNumber string `gorm:"not null"`
	MoveFromID  string `gorm:"not null"`
	MoveToID    string `gorm:"not null"`
}

type paymentModel struct {
	ID          string           `gorm:"primaryKey;not null;unique"`
	Amount      float64          `gorm:"not null"`
	RouteID     entities.RouteID `gorm:"not null"`
	PassengerID string           `gorm:"not null"`
	Passenger   passengerModel   `gorm:"foreignKey:PassengerID"`
}

type outboxDataModel struct {
	CreatedAt   time.Time        `gorm:"column:created_at;not null"`
	PaymentID   string           `gorm:"primaryKey;not null;unique"`
	RouteID     entities.RouteID `gorm:"not null"`
	PassengerID string           `gorm:"not null"`
	Passenger   passengerModel   `gorm:"foreignKey:PassengerID"`
}
