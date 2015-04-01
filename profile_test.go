package gosquare

import (
	"os"
	"testing"
)

func TestGetProfile(t *testing.T) {
	api := new(Profile)

	api.Get(os.Getenv("FOURSQUARE_TOKEN"))

	if api.Meta.Code != 200 {
		t.Errorf("Has errors %v", api)
	}
}
