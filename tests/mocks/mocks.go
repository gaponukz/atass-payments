package mocks

import (
	"fmt"
	"payments/src/entities"
	"payments/src/errors"
)

type queueMockStorage struct {
	payments []entities.OutboxData
}

type paymentStorageMock struct {
	payments []entities.Payment
}

func NewPaymentStorageMock() *paymentStorageMock {
	return &paymentStorageMock{}
}

func (m *paymentStorageMock) Create(payment entities.Payment) error {
	m.payments = append(m.payments, payment)
	return nil
}

func (m *paymentStorageMock) First() (entities.Payment, error) {
	if len(m.payments) == 0 {
		return entities.Payment{}, fmt.Errorf("blabla")
	}

	return m.payments[0], nil
}

func NewQueueMockStorage() *queueMockStorage {
	return &queueMockStorage{}
}

func (m *queueMockStorage) PopPayment() (entities.OutboxData, error) {
	if len(m.payments) == 0 {
		return entities.OutboxData{}, errors.ErrStorageEmpty
	}

	lastIndex := len(m.payments) - 1
	payment := m.payments[lastIndex]
	m.payments = m.payments[:lastIndex]

	return entities.OutboxData{
		PaymentID: payment.PaymentID,
		RouteID:   payment.RouteID,
		Passenger: payment.Passenger,
	}, nil
}

func (m *queueMockStorage) Create(payment entities.OutboxData) error {
	return m.PushBack(payment)
}

func (m *queueMockStorage) PushBack(payment entities.OutboxData) error {
	m.payments = append(m.payments, payment)
	return nil
}
