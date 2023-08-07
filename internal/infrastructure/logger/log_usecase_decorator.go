package logger

import (
	"fmt"
	"payments/internal/application/dto"
	"payments/internal/domain/entities"
)

type paymentService interface {
	OnSuccessfulPayment(dto.CratePaymentDTO) (entities.Payment, error)
}

type logPaymentServiceDecorator struct {
	s paymentService
	l logger
}

func NewlogPaymentServiceDecorator(s paymentService, l logger) logPaymentServiceDecorator {
	return logPaymentServiceDecorator{s: s, l: l}
}

func (s logPaymentServiceDecorator) OnSuccessfulPayment(d dto.CratePaymentDTO) (entities.Payment, error) {
	payment, err := s.s.OnSuccessfulPayment(d)
	if err != nil {
		s.l.Error(fmt.Sprintf("OnSuccessfulPayment failed with route %s: %v", d.RouteID, err))
		return payment, err
	}

	return payment, nil
}
