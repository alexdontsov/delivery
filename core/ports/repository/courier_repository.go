package repository

import (
    "context"
    "delivery/core/domain/model/courier"
    "github.com/google/uuid"
)

type CourierRepository interface {
    Get(ctx context.Context, id uuid.UUID) (*courier.Courier, error)
    Create(courier courier.Courier) (courier.Courier, error)
    Update(ctx context.Context, courier *courier.Courier) (*courier.Courier, error)
    GetListWithStatusReady() (couriers []*courier.Courier, err error)
}
