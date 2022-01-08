package unaswrappergo

import (
	"os"
	"testing"
)

func TestAuthwithAPIKey(t *testing.T) {
	var apiKey = os.Getenv("apiKey")
	uo, err := AuthwithAPIKey(apiKey)
	if err != nil {
		t.Error(err)
	}
	if uo.Login.Status != "ok" {
		t.Error("status not okay by response")
	}
}
