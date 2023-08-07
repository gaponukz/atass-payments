package logger

import (
	"fmt"
	"payments/internal/domain/entities"
)

type notifier interface {
	Notify(entities.OutboxData) error
}

type logger interface {
	Error(string)
	Info(string)
}

type logNotifier struct {
	n notifier
	l logger
}

func NewLoggingNotifierDecorator(n notifier, l logger) logNotifier {
	return logNotifier{n: n, l: l}
}

func (n logNotifier) Notify(payment entities.OutboxData) error {
	err := n.n.Notify(payment)
	if err != nil {
		n.l.Error(fmt.Sprintf("Can not send event %s: %v", payment.PaymentID, err))
		return err
	}

	n.l.Info(fmt.Sprintf("Send event: %s", payment.PaymentID))
	return nil
}
