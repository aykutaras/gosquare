package gosquare

import (
	"os"
	"testing"
)

func TestGetCheckIns(t *testing.T) {
	api := new(UserCheckIns)

	api.Get(os.Getenv("FOURSQUARE_TOKEN"))

	if api.Meta.Code != 200 {
		t.Errorf("Has errors %v", api)
	}
}
