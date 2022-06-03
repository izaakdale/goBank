package service

import (
	"time"

	"github.com/izaakdale/goBank/domain"
	"github.com/izaakdale/goBank/dto"
	"github.com/izaakdale/goBank/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}
type DefaultAccountService struct {
	repo domain.AccountRepo
}

func NewAccountService(repository domain.AccountRepo) DefaultAccountService {
	return DefaultAccountService{repository}
}

func (as DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}
	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Balance)

	newAccount, err := as.repo.SaveAccount(account)
	if err != nil {
		return nil, err
	}

	return newAccount.ToNewAccountResponseDto(), nil
}

func (as DefaultAccountService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account, err := as.repo.FindById(req.AccountId)
	if err != nil {
		return nil, err
	}
	if !account.CanWithdraw(req.Amount) && req.TransactionType == "withdrawal" {
		return nil, errs.NewValidationError("Insufficient Funds")
	}

	transaction := domain.Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	transactionResponse, err := as.repo.SaveTransaction(transaction)
	if err != nil {
		return nil, err
	}

	newTransactionResponse := transactionResponse.ToNewTransactionResponseDto()
	return &newTransactionResponse, nil
}
