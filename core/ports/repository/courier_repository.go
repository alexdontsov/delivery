package repository

import (
	"delivery/core/domain/model/courier"
	"delivery/core/domain/model/sharedkernel"
	"github.com/google/uuid"
)

type CourierRepository interface {
	Get(id uuid.UUID) courier.Courier
	Create(name string, transport courier.Transport, location sharedkernel.Location) (courier.Courier, error)
	Update(id uuid.UUID) (courier.Courier, error)
	GetListWithStatusReady() ([]courier.Courier, error)
}
