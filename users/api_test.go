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
	api := &Users{}

	api.GetCheckIns(AccessToken)

	if api.CheckIns.Meta.Code != 200 {
		t.Errorf("Has errors %v", api)
	}
}
