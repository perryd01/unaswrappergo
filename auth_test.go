package unaswrappergo

import (
	"os"
	"testing"
)

var apiKey = os.Getenv("apiKey")

func TestAuthwithAPIKey(t *testing.T) {
	uo, err := AuthwithAPIKey(apiKey)
	if err != nil {
		t.Error(err)
	}
	if uo.Login.Status != "ok" {
		t.Error("status not okay by response")
	}
}
