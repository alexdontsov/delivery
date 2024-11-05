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
    Id        uuid.UUID
    Status    Status
    location  sharedkernel.Location
    courierId uuid.UUID
}

func NewOrder(id uuid.UUID, location sharedkernel.Location) (Order, error) {
    return Order{
        Id:       id,
        Status:   Status{Value: Created.Value},
        location: location,
    }, ErrFailedCreateOrder
}

func AssignOrderToCourier(order Order, courierId uuid.UUID) (Order, error) {
    order.courierId = courierId
    order.Status = Status{Value: Assigned.Value}

    return order, nil
}

func CompleteOrder(order Order) (Order, error) {
    if order.Status.Value != Assigned.Value {
        return order, ErrFailedCompleteOrder
    }

    order.Status = Status{Value: Completed.Value}

    return order, ErrFailedCompleteOrder
}
