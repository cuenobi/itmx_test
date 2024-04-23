package usecase

import (
	"itmx_test/service/entity"
	"itmx_test/service/customer/repository"
	"itmx_test/util"
)

type CustomerUsecase interface {
	CreateCustomer(customer *entity.Customer) error
	GetCustomerByID(id string) (*entity.Customer, error)
	UpdateCustomerByID(customer *entity.Customer, id string) error
	DelCustomerByID(id string) error
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerUsecase(customerRepo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{customerRepo}
}

func (cu *customerUsecase) CreateCustomer(customer *entity.Customer) error {
	// generate uuid
	uuid := util.GenerateUuid()
	customer.ID = uuid

	if err := cu.customerRepo.Create(customer); err != nil {
		return err
	}

	return nil
}

func (cu *customerUsecase) GetCustomerByID(id string) (*entity.Customer, error) {
	customerExist, err := cu.customerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return customerExist, nil
}

func (cu *customerUsecase) UpdateCustomerByID(customer *entity.Customer, id string) error {
	customerExist, err := cu.customerRepo.FindByID(id)
	if err != nil {
		return err
	}

	customerExist.Name = customer.Name
	customerExist.Age = customer.Age

	if err := cu.customerRepo.Update(customerExist); err != nil {
		return err
	}

	return nil
}

func (cu *customerUsecase) DelCustomerByID(id string) error {
	customerExist, err := cu.customerRepo.FindByID(id)
	if err != nil {
		return err
	}

	if err := cu.customerRepo.DeleteByID(customerExist.ID); err != nil {
		return err
	}

	return nil
}
