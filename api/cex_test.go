package api_test

import (
	"testing"

	"frontendmasters.com/go/cryptomasters/api"
)

// test for an empty currency string
func TestAPICall_EmptyCurrency(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Errorf("expected an error for empty currency, but got nil")
	}
}

// test for a currency string with invalid length
func TestAPICall_InvalidCurrencyLength(t *testing.T) {
	_, err := api.GetRate("bt")
	if err == nil {
		t.Errorf("expected an error for invalid currency length, but got nil")
	}
}

// test for a valid currency string
func TestAPICall_ValidCurrency(t *testing.T) {
	_, err := api.GetRate("btc")
	if err != nil {
		t.Errorf("expected no error for valid currency, but got: %v", err)
	}
}
