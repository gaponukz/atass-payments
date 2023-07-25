package usecase

import "payments/src/entities"

type service struct{}

func (s service) ProcessPayment(data interface{}) error {
	return nil
}

func (s service) GeneratePayment(payment entities.Payment) error {
	return nil
}
