package usecase

import (
	"payments/internal/application/dto"
	"payments/internal/domain/entities"

	"github.com/google/uuid"
)

type createAbleStorages interface {
	Create(payment entities.Payment) error
}

type service struct {
	db createAbleStorages
}

func NewPaymentService(db createAbleStorages) service {
	return service{db: db}
}

func (s service) IsPaymentValid(info dto.CratePaymentDTO) bool {
	if info.Amount < 0 {
		return false
	}

	if info.RouteID == "" {
		return false
	}

	if info.Passenger.FullName == "" {
		return false
	}

	if info.Passenger.MoveFromID == "" {
		return false
	}

	if info.Passenger.MoveToID == "" {
		return false
	}

	if info.Passenger.PhoneNumber == "" {
		return false
	}

	return true
}

func (s service) OnSuccessfulPayment(info dto.CratePaymentDTO) (entities.Payment, error) {
	if info.Passenger.ID == "" {
		info.Passenger.ID = uuid.New().String()
	}

	payment := entities.Payment{
		ID:        uuid.New().String(),
		Amount:    info.Amount,
		RouteID:   info.RouteID,
		Passenger: info.Passenger,
	}

	err := s.db.Create(payment)
	if err != nil {
		return entities.Payment{}, err
	}

	return payment, nil
}
