package repository

import (
	"delivery/adapters/secondary/postgres/common"
	"delivery/core/domain/model/courier"
	"delivery/core/domain/model/sharedkernel"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourierRepository struct {
	db *gorm.DB
}

func NewCourierRepository(conn *gorm.DB) *CourierRepository {
	return &CourierRepository{db: conn}
}

func (r *CourierRepository) Get(id uuid.UUID) courier.Courier {

	var result courier.Courier
	r.db.Model(courier.Courier{Id: id}).First(&result)

	return result
}

func (r *CourierRepository) Create(name string, transport courier.Transport, location sharedkernel.Location) (courier.Courier, error) {
	newCourier := courier.NewCourier(name, transport, location)
	err := common.NewUnitOfWork(r.db).ExecuteInTransaction(func(tx *gorm.DB) error {
		var err error
		if err != nil {
			return err
		}

		tx.Create(newCourier)

		return nil
	})

	if err != nil {
		return newCourier, err
	}

	return newCourier, nil
}

func (r *CourierRepository) Update(id uuid.UUID) (courier.Courier, error) {
	//TODO:поиск курьера по id, затем его обновление
	var newCourier courier.Courier //этот код заглушка, потом удалить
	err := common.NewUnitOfWork(r.db).ExecuteInTransaction(func(tx *gorm.DB) error {
		//TODO:Реализация обновления курьера
		return nil
	})

	if err != nil {
		return newCourier, err
	}

	return newCourier, nil
}

func (r *CourierRepository) GetListWithStatusReady() ([]courier.Courier, error) {
	var couriers []courier.Courier

	r.db.Where(&courier.Courier{Status: courier.Ready}).Find(&couriers)

	return couriers, nil
}
