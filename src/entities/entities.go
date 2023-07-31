package entities

type RouteID string

type Passenger struct {
	ID          string `json:"id" gorm:"primaryKey;not null"`
	Gmail       string `json:"gmail" gorm:"not null"`
	FullName    string `json:"fullName" gorm:"not null"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null"`
	MoveFromID  string `json:"movingFromId" gorm:"not null"`
	MoveToID    string `json:"movingTowardsId" gorm:"not null"`
}

type Payment struct {
	ID          string    `json:"id" gorm:"primaryKey;not null;unique"`
	Amount      float64   `json:"amount" gorm:"not null"`
	RouteID     RouteID   `json:"routeId" gorm:"not null"`
	PassengerID string    `json:"passengerId" gorm:"not null"`
	Passenger   Passenger `json:"passenger" gorm:"foreignKey:PassengerID"`
}

type OutboxData struct {
	PaymentID   string    `json:"paymentId" gorm:"primaryKey;not null;unique"`
	RouteID     RouteID   `json:"routeId" gorm:"not null"`
	PassengerID string    `json:"passengerId" gorm:"not null"`
	Passenger   Passenger `json:"passenger" gorm:"foreignKey:PassengerID"`
}
