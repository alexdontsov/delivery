package repository

import (
	"context"
	"delivery/adapters/secondary/postgres/common"
	"delivery/core/domain/model/courier"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	otgorm "github.com/smacker/opentracing-gorm"
)

type CourierRepository struct {
	db *gorm.DB
}

func NewCourierRepository(conn *gorm.DB) *CourierRepository {
	return &CourierRepository{db: conn}
}

func (r *CourierRepository) Get(ctx context.Context, id uuid.UUID) (*courier.Courier, error) {
	db := otgorm.SetSpanToGorm(ctx, r.db)
	courierDb := &courier.Courier{}

	err := db.
		Preload("Transport").
		Where("id = ?", id).
		First(courierDb).
		Error

	if err != nil {
		return nil, err
	}

	return courierDb, err
}

func (r *CourierRepository) Create(courier courier.Courier) (courier.Courier, error) {
	err := common.NewUnitOfWork(r.db).ExecuteInTransaction(func(tx *gorm.DB) error {
		var err error
		if err != nil {
			return err
		}

		tx.Create(courier)

		return nil
	})

	if err != nil {
		return courier, err
	}

	return courier, nil
}

func (r *CourierRepository) Update(ctx context.Context, courier *courier.Courier) (*courier.Courier, error) {
	err := common.NewUnitOfWork(r.db).ExecuteInTransaction(func(tx *gorm.DB) error {
		var err error
		if err != nil {
			return err
		}

		tx.Save(courier)

		return nil
	})

	return courier, err
}

func (r *CourierRepository) GetListWithStatusReady() (couriers []*courier.Courier, err error) {
	err = r.db.
		Preload("Transport").
		Where("status = ? ", courier.Ready).
		Find(&couriers).
		Error

	return couriers, err
}
