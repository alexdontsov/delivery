package create_order

import (
	"context"
	"delivery/core/ports/repository"
)

type CreateOrderCommandHandler struct {
	orderRepository *repository.OrderRepository
}

func NewCreateOrderCommandHandler(orderRepository *repository.OrderRepository) *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{orderRepository: orderRepository}
}

func (c *CreateOrderCommandHandler) Handle(ctx context.Context, command *CreateOrderCommand) (*CreateOrderCommand, error) {

	createdProduct, err := c.orderRepository.Create(ctx, product)
	if err != nil {
		return nil, err
	}
	//TODO:тут надо подумать так как по идее надо отдавать response Command-ы
	return command, nil
}
