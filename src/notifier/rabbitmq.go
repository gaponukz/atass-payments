package notifier

import (
	"context"
	"encoding/json"
	"payments/src/entities"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMQNotifier struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewRabbitMQNotifier(url string) (*rabbitMQNotifier, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"payments_exchange", // exchange name
		amqp.ExchangeTopic,  // exchange type
		false,               // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"payments", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, err
	}

	return &rabbitMQNotifier{
		conn: conn,
		ch:   ch,
		q:    q,
	}, nil
}

func (r *rabbitMQNotifier) Notify(payment entities.Payment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	return r.ch.PublishWithContext(ctx,
		"payments_exchange", // exchange
		"",                  // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})

}

func (r *rabbitMQNotifier) Close() {
	_ = r.ch.Close()
	_ = r.conn.Close()
}
