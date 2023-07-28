package main

import (
	"fmt"
	"net/http"
	"payments/src/cdc_service"
	"payments/src/controller"
	"payments/src/logger"
	"payments/src/notifier"
	"payments/src/outbox"
	"payments/src/settings"
	"payments/src/storage"
	"payments/src/usecase"
)

func main() {
	config := settings.NewDotEnvSettings().Load()
	rabbitMQNotifier, err := notifier.NewRabbitMQNotifier(config.RabbitmqUrl)
	if err != nil {
		panic(err.Error())
	}

	defer rabbitMQNotifier.Close()

	logger := logger.NewConsoleLogger()
	paymentsDB := storage.NewJsonPaymentsStorage("payments.json")
	outboxDB := storage.NewJsonPaymentsStorage("outbox.json")
	loggeredOutboxDB := outbox.NewPopStorageLogger(outboxDB, logger)
	paymentService := usecase.NewPaymentService(paymentsDB)
	serviceWithOutbox := outbox.NewSaveToOutboxDecorator(paymentService, outboxDB)
	controller := controller.NewController(serviceWithOutbox)
	sendEventsService := outbox.NewSendEventsService(loggeredOutboxDB, rabbitMQNotifier)

	handler := http.NewServeMux()
	handler.HandleFunc("/processPayment", controller.ProcessPayment)

	server := http.Server{
		Addr:    ":9090",
		Handler: handler,
	}

	go cdc_service.NewCDCservice(sendEventsService).Serve()

	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
