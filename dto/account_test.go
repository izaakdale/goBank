package dto

import "testing"

func TestAccoutBalanceLimit(t *testing.T) {
	req := NewAccountRequest{
		CustomerId:  "1",
		Balance:     4999,
		AccountType: "checking",
	}

	err := req.Validate()

	if err != nil {
		if err.Message != BalanceTooLowErrorMessage {
			t.Error("Failed to identify that balance was too low")
		}
	} else {
		t.Error("Failed to throw error when validating invalid data")
	}
}
func TestAccoutInvalidType(t *testing.T) {
	req := NewAccountRequest{
		CustomerId:  "1",
		Balance:     6000,
		AccountType: "checkng",
	}

	err := req.Validate()

	if err != nil {
		if err.Message != InvalidTypeErrorMessage {
			t.Error("Failed to identify incorrect type")
		}
	} else {
		t.Error("Failed to throw error when validating invalid data")
	}
}
