package dto

import "payments/src/entities"

type CratePaymentDTO struct {
	Amount    float64            `json:"amount"`
	RouteID   entities.RouteID   `json:"routeId"`
	Passanger entities.Passanger `json:"passanger"`
}
