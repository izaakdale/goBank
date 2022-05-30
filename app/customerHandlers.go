package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/izaakdale/goBank/service"
	"github.com/izaakdale/utils-go/logger"
	"github.com/izaakdale/utils-go/response"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch CustomerHandlers) getCustomers(writer http.ResponseWriter, request *http.Request) {

	status := request.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomersByStatus(status)

	if err != nil {
		response.WriteJson(writer, err.Code, err.AsMessage())
	} else {
		response.WriteJson(writer, http.StatusOK, customers)
	}
}
func (ch CustomerHandlers) getCustomer(writer http.ResponseWriter, request *http.Request) {

	logger.Debug("Hitting getCustomer in handler")
	vars := mux.Vars(request)
	id := vars["id"]
	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customer)
	}
}

func writeResponse(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}
}
