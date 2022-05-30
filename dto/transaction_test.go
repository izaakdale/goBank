package dto

import (
	"net/http"
	"testing"
)

func TestTransactionErrorForInvalidType(t *testing.T) {
	// Arrange
	request := NewTransactionRequest{
		TransactionType: "invalid",
	}

	// Act
	err := request.Validate()

	// Assert
	if err != nil {
		if err.Message != TransactionTypeError {
			t.Error("Invalid message while testing transaction request")
		}
		if err.Code != http.StatusUnprocessableEntity {
			t.Error("Invalid http code while testing transaction request")
		}
	} else {
		t.Error("Error incorrectly assigned nil in fail case transaction request test")
	}
}

func TestTransactionErrorForInvalidAmount(t *testing.T) {
	// Arrange
	request := NewTransactionRequest{
		TransactionType: "withdrawal",
		Amount:          0,
	}

	// Act
	err := request.Validate()

	// Assert
	if err != nil {
		if err.Message != TransactionAmountError {
			t.Error("Invalid message while testing transaction request")
		}
		if err.Code != http.StatusUnprocessableEntity {
			t.Error("Invalid http code while testing transaction request")
		}
	} else {
		t.Error("Error incorrectly assigned nil in fail case transaction request test")
	}
}
