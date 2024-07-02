package order

import (
	"delivery/core/domain/model/sharedkernel"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrFailedCreateOrder   = errors.New("failed to create order")
	ErrFailedCompleteOrder = errors.New("failed to complete order")
)

type Order struct {
	id        uuid.UUID
	status    Status
	location  sharedkernel.Location
	courierId uuid.UUID
}

func NewOrder(id uuid.UUID, location sharedkernel.Location) (Order, error) {
	return Order{
		id:       id,
		status:   Status{Value: Created.Value},
		location: location,
	}, ErrFailedCreateOrder
}

func AssignOrderToCourier(order Order, courierId uuid.UUID) (Order, error) {
	order.courierId = courierId
	order.status = Status{Value: Assigned.Value}

	return order, nil
}

func CompleteOrder(order Order) (Order, error) {
	if order.status.Value != Assigned.Value {
		return order, ErrFailedCompleteOrder
	}

	order.status = Status{Value: Completed.Value}

	return order, ErrFailedCompleteOrder
}
