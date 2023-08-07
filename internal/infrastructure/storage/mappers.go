package storage

import (
	"payments/internal/domain/entities"
	"time"
)

func passengerToModel(passenger entities.Passenger) passengerModel {
	return passengerModel{
		ID:          passenger.ID,
		Gmail:       passenger.Gmail,
		FullName:    passenger.FullName,
		PhoneNumber: passenger.PhoneNumber,
		MoveFromID:  passenger.MoveFromID,
		MoveToID:    passenger.MoveToID,
	}
}

func passengerFromModel(passenger passengerModel) entities.Passenger {
	return entities.Passenger{
		ID:          passenger.ID,
		Gmail:       passenger.Gmail,
		FullName:    passenger.FullName,
		PhoneNumber: passenger.PhoneNumber,
		MoveFromID:  passenger.MoveFromID,
		MoveToID:    passenger.MoveToID,
	}
}

func paymentToModel(payment entities.Payment) paymentModel {
	return paymentModel{
		ID:          payment.ID,
		Amount:      payment.Amount,
		RouteID:     payment.RouteID,
		PassengerID: payment.Passenger.ID,
		Passenger:   passengerToModel(payment.Passenger),
	}
}

func outboxDataToModel(outboxData entities.OutboxData) outboxDataModel {
	return outboxDataModel{
		CreatedAt:   time.Now(),
		PaymentID:   outboxData.PaymentID,
		RouteID:     outboxData.RouteID,
		PassengerID: outboxData.Passenger.ID,
		Passenger:   passengerToModel(outboxData.Passenger),
	}
}

func outboxDataFromModel(outboxData outboxDataModel) entities.OutboxData {
	return entities.OutboxData{
		PaymentID: outboxData.PaymentID,
		RouteID:   outboxData.RouteID,
		Passenger: passengerFromModel(outboxData.Passenger),
	}
}
