package outbox

import (
	"fmt"
	"payments/src/entities"
	"payments/src/errors"
)

type errorLogger interface {
	Error(message string)
}

type popStorageLogger struct {
	storage popAbleStorage
	logger  errorLogger
}

func NewPopStorageLogger(storage popAbleStorage, logger errorLogger) popStorageLogger {
	return popStorageLogger{storage: storage, logger: logger}
}

func (l popStorageLogger) Pop() (entities.Payment, error) {
	payment, err := l.storage.Pop()
	if err != nil {
		if err != errors.ErrStorageEmpty {
			l.logger.Error(fmt.Sprintf("Can not pop payment from storage: %v", err))
		}
	}

	return payment, err
}

func (l popStorageLogger) Rollback(payment entities.Payment) error {
	err := l.storage.Rollback(payment)
	if err != nil {
		l.logger.Error(fmt.Sprintf("Can not rollback payment %s: %v", payment.ID, err))
	}

	return err
}
