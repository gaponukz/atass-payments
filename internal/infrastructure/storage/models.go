package storage

import (
	"payments/internal/domain/entities"
	"time"
)

type passengerModel struct {
	ID          string `json:"id" gorm:"primaryKey;not null"`
	Gmail       string `json:"gmail" gorm:"not null"`
	FullName    string `json:"fullName" gorm:"not null"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null"`
	MoveFromID  string `json:"movingFromId" gorm:"not null"`
	MoveToID    string `json:"movingTowardsId" gorm:"not null"`
}

type paymentModel struct {
	ID          string           `json:"id" gorm:"primaryKey;not null;unique"`
	Amount      float64          `json:"amount" gorm:"not null"`
	RouteID     entities.RouteID `json:"routeId" gorm:"not null"`
	PassengerID string           `json:"passengerId" gorm:"not null"`
	Passenger   passengerModel   `json:"passenger" gorm:"foreignKey:PassengerID"`
}

type outboxDataModel struct {
	CreatedAt   time.Time        `json:"createdAt" gorm:"column:created_at;not null"`
	PaymentID   string           `json:"paymentId" gorm:"primaryKey;not null;unique"`
	RouteID     entities.RouteID `json:"routeId" gorm:"not null"`
	PassengerID string           `json:"passengerId" gorm:"not null"`
	Passenger   passengerModel   `json:"passenger" gorm:"foreignKey:PassengerID"`
}
