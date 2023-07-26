package usecase

import (
	"payments/src/dto"
	"payments/src/entities"

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

func (s service) OnSuccessfulPayment(info dto.CratePaymentDTO) (entities.Payment, error) {
	payment := entities.Payment{
		ID:        uuid.New().String(),
		Amount:    info.Amount,
		RouteID:   info.RouteID,
		Passanger: info.Passanger,
	}

	err := s.db.Create(payment)
	if err != nil {
		return entities.Payment{}, err
	}

	return payment, nil
}
