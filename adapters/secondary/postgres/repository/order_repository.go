package repository

import (
	"context"
	"delivery/adapters/secondary/postgres/common"
	"delivery/core/domain/model/order"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	otgorm "github.com/smacker/opentracing-gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(conn *gorm.DB) *OrderRepository {
	return &OrderRepository{db: conn}
}

func (r *OrderRepository) Get(ctx context.Context, id uuid.UUID) (*order.Order, error) {
	db := otgorm.SetSpanToGorm(ctx, r.db)
	orderDb := &order.Order{}

	err := db.
		Preload("Transport").
		Where("id = ?", id).
		First(orderDb).
		Error

	if err != nil {
		return nil, err
	}

	return orderDb, err
}

func (r *OrderRepository) Create(order order.Order) (order.Order, error) {
	err := common.NewUnitOfWork(r.db).ExecuteInTransaction(func(tx *gorm.DB) error {
		var err error
		if err != nil {
			return err
		}

		tx.Create(order)

		return nil
	})

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *OrderRepository) Update(ctx context.Context, order *order.Order) (*order.Order, error) {
	err := common.NewUnitOfWork(r.db).ExecuteInTransaction(func(tx *gorm.DB) error {
		var err error
		if err != nil {
			return err
		}

		err = r.db.Save(order).Error

		return nil
	})

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *OrderRepository) GetListWithStatusCreated() (orders []*order.Order, err error) {
	err = r.db.
		Where("status = ? ", order.Assigned).
		Find(&orders).
		Error

	return orders, err
}

func (r *OrderRepository) GetListWithStatusAssigned() (orders []*order.Order, err error) {
	err = r.db.
		Where("status = ? ", order.Assigned).
		Find(&orders).
		Error

	return orders, err
}
