package repository

import (
	"delivery/core/domain/model/order"
	"delivery/core/domain/model/sharedkernel"
	"github.com/google/uuid"
)

type OrderRepository interface {
	Get(id uuid.UUID) order.Order
	Create(id uuid.UUID, location sharedkernel.Location) (order.Order, error)
	Update(id uuid.UUID) (order.Order, error)
	GetListWithStatusCreated() ([]order.Order, error)
	GetListWithStatusAssigned() ([]order.Order, error)
}
