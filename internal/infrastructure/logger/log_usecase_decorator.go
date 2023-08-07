package logger

import (
	"fmt"
	"payments/internal/application/dto"
	"payments/internal/domain/entities"
)

type paymentService interface {
	IsPaymentValid(dto.CratePaymentDTO) bool
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

func (s logPaymentServiceDecorator) IsPaymentValid(d dto.CratePaymentDTO) bool {
	isValid := s.s.IsPaymentValid(d)
	if !isValid {
		strPassenger := fmt.Sprintf(
			"{Gmail: %s, FullName: %s, PhoneNumber: %s, MoveFromID: %s, MoveToID: %s}",
			d.Passenger.Gmail, d.Passenger.FullName, d.Passenger.PhoneNumber, d.Passenger.MoveFromID, d.Passenger.MoveToID,
		)
		strPayment := fmt.Sprintf("{Amount: %f, RouteID: %s, Passenger: %s}", d.Amount, d.RouteID, strPassenger)

		s.l.Info(fmt.Sprintf("Get invalid payment %s", strPayment))
	}

	return isValid
}
