package usecase

import (
	"testing"

	"itmx_test/domain"
	"itmx_test/service/entity"

	"github.com/stretchr/testify/assert"
)

type mockCustomerRepo struct {
	CreateFunc     func(customer *entity.Customer) error
	FindByIDFunc   func(id string) (*entity.Customer, error)
	UpdateFunc     func(customer *entity.Customer) error
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
		}
		usecase := NewCustomerUsecase(repo)

		err := usecase.CreateCustomer(&entity.Customer{Name: "test", Age: 11})
		assert.NoError(t, err)
	})
}

func TestGetCustomerByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedCustomer := &entity.Customer{Name: "test", Age: 11}
		repo := &mockCustomerRepo{
			FindByIDFunc: func(id string) (*entity.Customer, error) {
				return expectedCustomer, nil
			},
		}
		usecase := NewCustomerUsecase(repo)

		customer, err := usecase.GetCustomerByID("123")
		assert.NoError(t, err)
		assert.Equal(t, expectedCustomer, customer)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := domain.ErrNotFound
		repo := &mockCustomerRepo{
			FindByIDFunc: func(id string) (*entity.Customer, error) {
				return nil, expectedErr
			},
		}
		usecase := NewCustomerUsecase(repo)

		customer, err := usecase.GetCustomerByID("123")
		assert.Error(t, err)
		assert.Nil(t, customer)
		assert.Equal(t, expectedErr, err)
	})
}

func TestUpdateCustomerByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedCustomer := &entity.Customer{Name: "updated", Age: 30}
		repo := &mockCustomerRepo{
			FindByIDFunc: func(id string) (*entity.Customer, error) {
				return &entity.Customer{Name: "test", Age: 11}, nil
			},
			UpdateFunc: func(customer *entity.Customer) error {
				assert.Equal(t, expectedCustomer.Name, customer.Name)
				assert.Equal(t, expectedCustomer.Age, customer.Age)
				return nil
			},
		}
		usecase := NewCustomerUsecase(repo)

		err := usecase.UpdateCustomerByID(expectedCustomer, "123")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := domain.ErrNotFound
		updateCustomer := &entity.Customer{Name: "updated", Age: 30}
		repo := &mockCustomerRepo{
			FindByIDFunc: func(id string) (*entity.Customer, error) {
				return nil, expectedErr
			},
			UpdateFunc: func(customer *entity.Customer) error {
				return nil
			},
		}
		usecase := NewCustomerUsecase(repo)

		err := usecase.UpdateCustomerByID(updateCustomer, "123")
		assert.Equal(t, expectedErr, err)
	})
}

func TestDelCustomerByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedCustomer := &entity.Customer{Name: "test", Age: 11}
		repo := &mockCustomerRepo{
			FindByIDFunc: func(id string) (*entity.Customer, error) {
				return expectedCustomer, nil
			},
			DeleteByIDFunc: func(id string) error {
				return nil
			},
		}
		usecase := NewCustomerUsecase(repo)

		err := usecase.DelCustomerByID("123")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := domain.ErrNotFound
		repo := &mockCustomerRepo{
			FindByIDFunc: func(id string) (*entity.Customer, error) {
				return nil, domain.ErrNotFound
			},
			DeleteByIDFunc: func(id string) error {
				return nil
			},
		}
		usecase := NewCustomerUsecase(repo)

		err := usecase.DelCustomerByID("123")
		assert.Equal(t, expectedErr, err)
	})
}
