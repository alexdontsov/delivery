package repository

import (
    "context"
    "delivery/core/domain/model/order"
    "github.com/google/uuid"
)

type OrderRepository interface {
    Get(ctx context.Context, id uuid.UUID) (*order.Order, error)
    Create(order order.Order) (order.Order, error)
    Update(ctx context.Context, order *order.Order) (*order.Order, error)
    GetListWithStatusCreated() ([]*order.Order, error)
    GetListWithStatusAssigned() ([]*order.Order, error)
}
