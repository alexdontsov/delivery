package repository

import (
	"delivery/adapters/secondary/postgres/common"
	"delivery/core/domain/model/order"
	"delivery/core/domain/model/sharedkernel"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(conn *gorm.DB) *OrderRepository {
	return &OrderRepository{db: conn}
}

func (r *OrderRepository) Get(id uuid.UUID) order.Order {

	var result order.Order
	r.db.Model(order.Order{Id: id}).First(&result)

	return result
}

func (r *OrderRepository) Create(id uuid.UUID, location sharedkernel.Location) (order.Order, error) {
	newOrder, _ := order.NewOrder(id, location)
	err := common.NewUnitOfWork(r.db).ExecuteInTransaction(func(tx *gorm.DB) error {
		var err error
		if err != nil {
			return err
		}

		tx.Create(newOrder)

		return nil
	})

	if err != nil {
		return newOrder, err
	}

	return newOrder, nil
}

func (r *OrderRepository) Update(id uuid.UUID) (order.Order, error) {
	//TODO: поиск по idи потом обновление
	var newOrder order.Order //этот код заглушка
	err := common.NewUnitOfWork(r.db).ExecuteInTransaction(func(tx *gorm.DB) error {
		//TODO:Реализация обновления заказа
		return nil
	})

	if err != nil {
		return newOrder, err
	}

	return newOrder, nil
}

func (r *OrderRepository) GetListWithStatusCreated() ([]order.Order, error) {
	var orders []order.Order

	r.db.Where(&order.Order{Status: order.Created}).Find(&orders)

	return orders, nil
}

func (r *OrderRepository) GetListWithStatusAssigned() ([]order.Order, error) {
	var orders []order.Order

	r.db.Where(&order.Order{Status: order.Assigned}).Find(&orders)

	return orders, nil
}
