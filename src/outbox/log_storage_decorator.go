package outbox

import (
	"fmt"
	"payments/src/entities"
	"payments/src/errors"
)

type errorLogger interface {
	Info(message string)
	Error(message string)
}

type paymentStorage interface {
	Create(entities.Payment) error
	PopPayment() (entities.OutboxData, error)
	PushBack(payment entities.OutboxData) error
}

type storageLogger struct {
	storage paymentStorage
	logger  errorLogger
}

func NewStorageLoggerDecorator(storage paymentStorage, logger errorLogger) storageLogger {
	return storageLogger{storage: storage, logger: logger}
}

func (l storageLogger) Create(payment entities.Payment) error {
	err := l.storage.Create(payment)
	if err != nil {
		l.logger.Error(fmt.Sprintf("Can not crate payment: %v", err))
		return err
	}

	l.logger.Info(fmt.Sprintf("Create payment: %s", payment.ID))
	return nil
}

func (l storageLogger) PopPayment() (entities.OutboxData, error) {
	payment, err := l.storage.PopPayment()
	if err != nil {
		if err != errors.ErrStorageEmpty {
			l.logger.Error(fmt.Sprintf("Can not PopPayment from storage: %v", err))
		}

		return payment, err
	}

	l.logger.Info(fmt.Sprintf("PopPayment: %s", payment.PaymentID))
	return payment, err
}

func (l storageLogger) PushBack(payment entities.OutboxData) error {
	err := l.storage.PushBack(payment)
	if err != nil {
		l.logger.Error(fmt.Sprintf("Can not PushBack payment %s: %v", payment.PaymentID, err))
		return err
	}

	l.logger.Info(fmt.Sprintf("PushBack payment: %s", payment.PaymentID))
	return err
}
