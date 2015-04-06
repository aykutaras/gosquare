package gosquare

import (
	"os"
	"testing"
)

func TestGetCheckIns(t *testing.T) {
	api := new(UserCheckIns)

	token := os.Getenv("FOURSQUARE_TOKEN")

	if token == "" {
		t.Errorf("Invalid token %v", token)
	}

	api.Get(token)

	if api.Meta.Code != 200 {
		t.Errorf("Has errors %v", api.Meta.ErrorType)
	}
}
