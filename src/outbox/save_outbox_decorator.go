package outbox

import (
	"payments/src/dto"
	"payments/src/entities"
)

type paymentService interface {
	OnSuccessfulPayment(dto.CratePaymentDTO) (entities.Payment, error)
}

type outboxStorage interface {
	Create(payment entities.Payment) error
}

type saveToOutboxDecorator struct {
	service paymentService
	storage outboxStorage
}

func NewSaveToOutboxDecorator(service paymentService, storage outboxStorage) saveToOutboxDecorator {
	return saveToOutboxDecorator{
		service: service,
		storage: storage,
	}
}

func (s saveToOutboxDecorator) OnSuccessfulPayment(info dto.CratePaymentDTO) (entities.Payment, error) {
	payment, err := s.service.OnSuccessfulPayment(info)
	if err != nil {
		return payment, err
	}

	err = s.storage.Create(payment)

	return payment, err
}
