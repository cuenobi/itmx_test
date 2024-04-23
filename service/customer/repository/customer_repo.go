package repository

import (
	"itmx_test/service/entity"

	"itmx_test/domain"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	// Create
	Create(customer *entity.Customer) error

	// Read
	FindByID(id string) (*entity.Customer, error)

	// Update
	Update(customer *entity.Customer) error

	// Delete
	DeleteByID(id string) error
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepo{db}
}

func (cr *customerRepo) Create(customer *entity.Customer) error {
	tx := cr.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Create user
	if err := tx.Create(customer).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (cr *customerRepo) FindByID(id string) (*entity.Customer, error) {
	customer := &entity.Customer{}
	if err := cr.db.Order("created_at desc").Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, domain.ErrNotFound
	}
	return customer, nil
}

func (cr *customerRepo) Update(customer *entity.Customer) error {
	tx := cr.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// update
	if err := tx.Save(customer).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (cr *customerRepo) DeleteByID(id string) error {
	customer := &entity.Customer{}
	if err := cr.db.Where("id = ?", id).Delete(&customer).Error; err != nil {
		return err
	}
	return nil
}