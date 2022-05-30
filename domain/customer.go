package domain

import (
	"github.com/izaakdale/goBank/dto"
	"github.com/izaakdale/goBank/errs"
)

type Customer struct {
	Id      string `db:"customer_id"`
	Name    string
	City    string
	Zipcode string
	Dob     string `db:"date_of_birth"`
	Status  string
}

func (customer Customer) ToStatusAsText() string {
	statusAsText := "active"
	if customer.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (customer Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:      customer.Id,
		Name:    customer.Name,
		City:    customer.City,
		Zipcode: customer.Zipcode,
		Dob:     customer.Dob,
		Status:  customer.ToStatusAsText(),
	}
}

type CustomerRepo interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
