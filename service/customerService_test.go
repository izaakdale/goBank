package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	realDomain "github.com/izaakdale/goBank/domain"
	"github.com/izaakdale/goBank/errs"
	"github.com/izaakdale/goBank/mocks/domain"
)

func TestReturnErrorForInvalidStatus(t *testing.T) {

	// invalid status
	req := "10"
	// not testing actual db here so don't need repo
	service := NewCustomerRepoService(nil)
	_, err := service.GetAllCustomersByStatus(req)

	if err == nil {
		t.Error("Test failed to throw error for invalid status")
	}

}

var mockCustomerRepo *domain.MockCustomerRepo
var customerService CustomerService

func setupCustomerTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockCustomerRepo = domain.NewMockCustomerRepo(ctrl)
	customerService = NewCustomerRepoService(mockCustomerRepo)

	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func TestCustomersReturned(t *testing.T) {
	teardown := setupCustomerTest(t)
	defer teardown()

	dummyCustomers := []realDomain.Customer{
		{
			Id:      "1001",
			Name:    "Izaak",
			City:    "Vancouver",
			Zipcode: "V5L4R2",
			Dob:     "1993-08-28",
			Status:  "1",
		},
		{
			Id:      "1002",
			Name:    "Mahtab",
			City:    "Vancouver",
			Zipcode: "V5L4R2",
			Dob:     "1996-07-07",
			Status:  "1",
		},
	}

	mockCustomerRepo.EXPECT().FindAll("").Return(dummyCustomers, nil)
	_, err := customerService.GetAllCustomersByStatus("")

	if err != nil {
		t.Error("Error returned for valid response")
	}
}

func TestFindById(t *testing.T) {
	teardown := setupCustomerTest(t)
	defer teardown()

	customer := realDomain.Customer{
		Id:      "100",
		Name:    "Test",
		City:    "Test",
		Zipcode: "10Test",
		Dob:     "1992-10-10",
		Status:  "1",
	}
	mockCustomerRepo.EXPECT().FindById(customer.Id).Return(&customer, nil)
	response, err := customerService.GetCustomer(customer.Id)

	if err != nil {
		t.Error("Get customer returned error for valid request")
	}
	if customer.Id != response.Id {
		t.Error("Returned incorrect customer ID")
	}
}

func TestFindByIdReturnsNoCustomer(t *testing.T) {

	teardown := setupCustomerTest(t)
	defer teardown()

	dbError := errs.NewNotFoundError("Something")
	missingId := "101"

	mockCustomerRepo.EXPECT().FindById(missingId).Return(nil, dbError)
	response, err := customerService.GetCustomer(missingId)

	if response != nil {
		t.Error("Service providing customer when it wasn't provided by DB")
	}
	if err == nil {
		t.Error("Service failed to pass error back from DB")
	}
}
