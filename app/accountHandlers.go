package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/izaakdale/goBank/dto"
	"github.com/izaakdale/goBank/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) NewAccount(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	customerId := vars["id"]

	var accountRequest dto.NewAccountRequest
	err := json.NewDecoder(request.Body).Decode(&accountRequest)
	if err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		accountRequest.CustomerId = customerId
		account, appError := ah.service.NewAccount(accountRequest)
		if appError != nil {
			writeResponse(writer, appError.Code, appError.Message)
		} else {
			writeResponse(writer, http.StatusCreated, account)
		}
	}
}

func (ah AccountHandler) Transaction(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]

	var transactionRequest dto.NewTransactionRequest
	err := json.NewDecoder(request.Body).Decode(&transactionRequest)
	if err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		transactionRequest.CustomerId = customerId
		transactionRequest.AccountId = accountId
		account, appError := ah.service.NewTransaction(transactionRequest)
		if appError != nil {
			writeResponse(writer, appError.Code, appError.Message)
		} else {
			writeResponse(writer, http.StatusCreated, account)
		}
	}

}
