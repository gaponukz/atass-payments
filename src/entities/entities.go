package entities

type RouteID string

type Passanger struct {
	ID          string `json:"id"`
	Gmail       string `json:"gmail"`
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
	MoveFromID  string `json:"movingFromId"`
	MoveToID    string `json:"movingTowardsID"`
}

type Payment struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	RouteID   RouteID   `json:"routeId"`
	Passanger Passanger `json:"passanger"`
}
