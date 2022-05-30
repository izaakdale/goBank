package service

import (
	"github.com/izaakdale/goBank/domain"
	"github.com/izaakdale/goBank/dto"
	"github.com/izaakdale/goBank/errs"
)

type CustomerService interface {
	GetAllCustomersByStatus(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepo
}

func (service DefaultCustomerService) GetAllCustomersByStatus(status string) ([]dto.CustomerResponse, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else if status != "" {
		return nil, errs.NewNotFoundError("Status invalid")
	}

	customers, err := service.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	var customerResponse []dto.CustomerResponse
	for _, customer := range customers {
		customerResponse = append(customerResponse, customer.ToDto())
	}
	return customerResponse, nil
}
func (service DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := service.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	customerResponse := customer.ToDto()
	return &customerResponse, nil
}
func NewCustomerRepoService(repository domain.CustomerRepo) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
