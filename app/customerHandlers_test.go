package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/izaakdale/goBank/dto"
	"github.com/izaakdale/goBank/errs"
	"github.com/izaakdale/goBank/mocks/service"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)

	ch = CustomerHandlers{service: mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getCustomers)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func TestReturnCustomerListWithStatus200(t *testing.T) {
	//arrange
	teardown := setup(t)
	defer teardown()
	dummyCustomers := []dto.CustomerResponse{
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
	mockService.EXPECT().GetAllCustomersByStatus("").Return(dummyCustomers, nil)
	req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	//act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//assert
	if recorder.Code != http.StatusOK {
		t.Error("Incorrect status code for customers request")
	}
}

func TestReturnErrorWithStatus500(t *testing.T) {
	//arrange
	teardown := setup(t)
	defer teardown()
	mockService.EXPECT().GetAllCustomersByStatus("").Return(nil, errs.NewUnexpectedError("Something"))
	req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	//act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Incorrect status code for failed customers request")
	}
}
