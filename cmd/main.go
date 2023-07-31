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

	paymentsDB, err := storage.NewPostgresUserStorage(storage.PostgresCredentials{
		Host:     config.PostgresHost,
		User:     config.PostgresUser,
		Password: config.PostgresPassword,
		Dbname:   config.PostgresDbname,
		Port:     config.PostgresPort,
		Sslmode:  config.PostgresSslmode,
	})
	if err != nil {
		panic(err.Error())
	}

	logger := logger.NewConsoleLogger()
	loggeredPaymentsDB := outbox.NewStorageLoggerDecorator(paymentsDB, logger)
	paymentService := usecase.NewPaymentService(loggeredPaymentsDB)
	controller := controller.NewController(paymentService)
	sendEventsService := outbox.NewSendEventsService(loggeredPaymentsDB, rabbitMQNotifier)

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
