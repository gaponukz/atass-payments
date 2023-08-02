package outbox

import (
	"payments/src/dto"
	"payments/src/entities"
)

type paymentService interface {
	OnSuccessfulPayment(dto.CratePaymentDTO) (entities.Payment, error)
}

type outboxService interface {
	SendNewEvent() error
}

type triggerOutboxOnSuccessfulPayment struct {
	ps paymentService
	os outboxService
}

func NewTriggerOutboxDecorator(ps paymentService, os outboxService) triggerOutboxOnSuccessfulPayment {
	return triggerOutboxOnSuccessfulPayment{ps: ps, os: os}
}

func (s triggerOutboxOnSuccessfulPayment) OnSuccessfulPayment(d dto.CratePaymentDTO) (entities.Payment, error) {
	p, err := s.ps.OnSuccessfulPayment(d)
	if err == nil {
		_ = s.os.SendNewEvent()
	}

	return p, err
}
