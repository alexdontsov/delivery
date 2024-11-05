package services

import (
    "delivery/core/domain/model/courier"
    "delivery/core/domain/model/order"
)

type DispatchServiceInterface interface {
    Dispatch(order order.Order, couriers []courier.Courier)
}

type DispatchService struct{}

func (DispatchService) Dispatch(order order.Order, couriers []courier.Courier) {

}
