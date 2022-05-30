package dto

import (
	"strings"

	"github.com/izaakdale/goBank/errs"
)

type NewTransactionRequest struct {
	CustomerId      string  `json:"customer_id"`
	AccountId       string  `json:"account_type"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

const TransactionTypeError = "Transaction type should be withdrawal or deposit"
const TransactionAmountError = "Amount should be greater than 0"

func (req NewTransactionRequest) isTransactionWithdrawal() bool {
	return strings.ToLower(req.TransactionType) == "withdrawal"
}
func (req NewTransactionRequest) isTransactionDeposit() bool {
	return strings.ToLower(req.TransactionType) == "deposit"
}

func (req NewTransactionRequest) Validate() *errs.AppError {
	if !req.isTransactionDeposit() && !req.isTransactionWithdrawal() {
		return errs.NewValidationError(TransactionTypeError)
	}
	if req.Amount <= 0 {
		return errs.NewValidationError(TransactionAmountError)
	}
	return nil
}
