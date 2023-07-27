package mocks

import (
	"payments/src/entities"
	"payments/src/errors"
)

type queueMockStorage struct {
	payments []entities.Payment
}

func NewQueueMockStorage() *queueMockStorage {
	return &queueMockStorage{}
}

func (m *queueMockStorage) Pop() (entities.Payment, error) {
	if len(m.payments) == 0 {
		return entities.Payment{}, errors.ErrStorageEmpty
	}

	lastIndex := len(m.payments) - 1
	payment := m.payments[lastIndex]
	m.payments = m.payments[:lastIndex]

	return payment, nil
}

func (m *queueMockStorage) Rollback(payment entities.Payment) error {
	m.payments = append(m.payments, payment)
	return nil
}
