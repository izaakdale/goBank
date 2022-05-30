package dto

import (
	"strings"

	"github.com/izaakdale/goBank/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Balance     float64 `json:"balance"`
}

const BalanceTooLowErrorMessage = "Balance too low for new account"
const InvalidTypeErrorMessage = "Account type should be checking or saving"

func (req NewAccountRequest) Validate() *errs.AppError {
	if req.Balance < 5000 {
		return errs.NewValidationError(BalanceTooLowErrorMessage)
	}
	if strings.ToLower(req.AccountType) != "saving" && strings.ToLower(req.AccountType) != "checking" {
		return errs.NewValidationError(InvalidTypeErrorMessage)
	}
	return nil
}
