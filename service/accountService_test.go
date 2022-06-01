package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	realDomain "github.com/izaakdale/goBank/domain"
	"github.com/izaakdale/goBank/dto"
	"github.com/izaakdale/goBank/errs"
	"github.com/izaakdale/goBank/mocks/domain"
)

func TestReturnValidationErrorForInvalidRequest(t *testing.T) {
	//arrange
	req := dto.NewAccountRequest{
		CustomerId:  "100",
		Balance:     0,
		AccountType: "saving",
	}

	// does not need access to db for validating request
	service := NewAccountService(nil)

	//act
	_, err := service.NewAccount(req)

	if err == nil {
		t.Error("Should return error since request was invalid")
	}
}

var req dto.NewAccountRequest
var mockRepo *domain.MockAccountRepo
var service AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepo(ctrl)
	service = NewAccountService(mockRepo)

	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func TestReturnErrorFromServerForAccountCreationFailure(t *testing.T) {
	//arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		Balance:     6000,
		AccountType: "saving",
	}

	account := realDomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Balance:     req.Balance,
		Status:      "1",
	}

	//act
	mockRepo.EXPECT().SaveAccount(account).Return(nil, errs.NewUnexpectedError("Unexpected DB Error"))
	_, err := service.NewAccount(req)

	//assert
	if err == nil {
		t.Error("Test failed while validating legitimate data")
	}
}

func TestReturnNewAccountForSuccessfulAccountCreation(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		Balance:     6000,
		AccountType: "saving",
	}

	account := realDomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Balance:     req.Balance,
		Status:      "1",
	}
	accountWithId := account
	accountWithId.AccountId = "201"
	mockRepo.EXPECT().SaveAccount(account).Return(&accountWithId, nil)

	//act
	newAccount, err := service.NewAccount(req)

	//assert
	if err != nil {
		t.Error("Test failed to create new account")
	}
	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("Failed to create account with correct ID")
	}
}
