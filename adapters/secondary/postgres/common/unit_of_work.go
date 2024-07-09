package common

import "gorm.io/gorm"

type UnitOfWork interface {
	ExecuteInTransaction(fn func(tx *gorm.DB) error) error
}

type unitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

func (u *unitOfWork) ExecuteInTransaction(fn func(tx *gorm.DB) error) error {
	tx := u.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
