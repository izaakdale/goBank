package domain

import (
	"github.com/izaakdale/goBank/dto"
	"github.com/izaakdale/goBank/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Balance     float64 `db:"amount"`
	Status      string  `db:"status"`
}

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepo.go -package=domain github.com/izaakdale/goBank/domain AccountRepo
type AccountRepo interface {
	FindById(string) (*Account, *errs.AppError)
	SaveAccount(Account) (*Account, *errs.AppError)
	// UpdateAccount(Account, *Transaction) *errs.AppError
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (account Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: account.AccountId,
	}
}

func (transaction Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:   transaction.TransactionId,
		AccountId:       transaction.AccountId,
		NewBalance:      transaction.Amount,
		TransactionType: transaction.TransactionType,
		TransactionDate: transaction.TransactionDate,
	}
}

func (account Account) CanWithdraw(amount float64) bool {
	return account.Balance >= amount
}

func (transaction Transaction) IsWithdrawal() bool {
	return transaction.TransactionType == "withdrawal"
}
