package users

import (
	"testing"
)

const (
	ClientId     = "32R11TKNCLDMOZQQXLDGKX1J3OCTDJTPI21HWL35V4LF1NCC"
	ClientSecret = "SUSVS34WO3ANGHYSYOSA5HOPKWB0M0E1LN3SICZKFCF1ZZIZ"
	AccessToken  = "DKXXGSXF35LRBERMJ0XYOBCASR3JSECERVQ4ELJTYVMG0NW1"
)

func TestGetCheckIns(t *testing.T) {
	api := new(UserCheckIns)

	api.Get(AccessToken)

	if api.Meta.Code != 200 {
		t.Errorf("Has errors %v", api)
	}
}

func TestGetProfile(t *testing.T) {
	api := new(Profile)

	api.Get(AccessToken)

	if api.Meta.Code != 200 {
		t.Errorf("Has errors %v", api)
	}
}

func TestGetFriends(t *testing.T) {
	api := new(Friends)

	api.Get(AccessToken)

	if api.Meta.Code != 200 {
		t.Errorf("Has errors %v", api)
	}
}
