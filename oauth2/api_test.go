package oauth2

import (
	"testing"
)

const (
	ClientId     = "32R11TKNCLDMOZQQXLDGKX1J3OCTDJTPI21HWL35V4LF1NCC"
	ClientSecret = "SUSVS34WO3ANGHYSYOSA5HOPKWB0M0E1LN3SICZKFCF1ZZIZ"
)

func TestAuthenticate(t *testing.T) {
	api := &Auth{ClientId: ClientId, ClientSecret: ClientSecret}

	authUrl := api.Authenticate("localhost")
	correctUrl := "https://foursquare.com/oauth2/authenticate?client_id=32R11TKNCLDMOZQQXLDGKX1J3OCTDJTPI21HWL35V4LF1NCC&response_type=code&redirect_uri=http://localhost/code"

	if authUrl != correctUrl {
		t.Errorf("Invalid url: %v", authUrl)
	}
}

func TestAccessToken(t *testing.T) {
	api := &Auth{ClientId: ClientId, ClientSecret: ClientSecret}

	// Find a way to automatically generate this code
	code := "OKP5EONIRKOCVWVJ0BHIAWFVBJMV35LICGAGKPEDXPW54ITI#_=_"
	api.GetAccessToken("localhost", code)
	if api.AccessToken != "DKXXGSXF35LRBERMJ0XYOBCASR3JSECERVQ4ELJTYVMG0NW1" {
		t.Errorf("%v invalid access token", api.AccessToken)
	}
}
