package event

import "time"

type ListOrders struct {
	Name    string
	Payload interface{}
}

func NewListOrders() *ListOrders {
	return &ListOrders{
		Name: "ListOrders",
	}
}

func (e *ListOrders) GetName() string {
	return e.Name
}

func (e *ListOrders) GetPayload() interface{} {
	return e.Payload
}

func (e *ListOrders) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *ListOrders) GetDateTime() time.Time {
	return time.Now()
}
