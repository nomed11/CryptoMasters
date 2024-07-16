package api_test

import (
	"testing"

	"frontendmasters.com/go/cryptomasters/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Errorf("error was not found!")
	}
}