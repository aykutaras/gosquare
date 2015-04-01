package gosquare

import (
	"os"
	"testing"
)

func TestGetFriends(t *testing.T) {
	api := new(Friends)
	api.Get(os.Getenv("FOURSQUARE_TOKEN"))

	if api.Meta.Code != 200 {
		t.Errorf("Has errors %v", api)
	}
}
