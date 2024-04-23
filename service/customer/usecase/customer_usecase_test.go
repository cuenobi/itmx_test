package usecase

import (
	"testing"

	"itmx_test/service/entity"

	"github.com/stretchr/testify/assert"
)

type mockCustomerRepo struct {
	CreateFunc     func(customer *entity.Customer) error
	FindByIDFunc   func(id string) (*entity.Customer, error)
	UpdateFunc    func(customer *entity.Customer) error
	DeleteByIDFunc func(id string) error
}

func (m *mockCustomerRepo) Create(customer *entity.Customer) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(customer)
	}
	return nil
}

func (m *mockCustomerRepo) FindByID(id string) (*entity.Customer, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return nil, nil
}

func (m *mockCustomerRepo) Update(customer *entity.Customer) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(customer)
	}
	return nil
}

func (m *mockCustomerRepo) DeleteByID(id string) error {
	if m.DeleteByIDFunc != nil {
		return m.DeleteByIDFunc(id)
	}
	return nil
}

func TestCreateCustomer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockCustomerRepo{
			CreateFunc: func(customer *entity.Customer) error {
				return nil
			},
			FindByIDFunc: func(id string) (*entity.Customer, error) {
				return &entity.Customer{Name: "test", Age: 11}, nil
			},
			UpdateFunc: func(customer *entity.Customer) error {
				return nil
			},
			DeleteByIDFunc: func(id string) error {
				return nil
			},
		}
		service := NewCustomerUsecase(repo)

		err := service.CreateCustomer(&entity.Customer{Name: "test", Age: 11})
		assert.NoError(t, err)
	})
}
