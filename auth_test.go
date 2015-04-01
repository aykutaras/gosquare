package gosquare

import (
	"fmt"
	"os"
	"testing"
)

const (
	clientIdEnvKey     = "FOURSQUARE_CLIENTID"
	clientSecretEnvKey = "FOURSQUARE_CLIENTSECRET"
	redirectUri        = "http://localhost/"
)

func TestAuthenticate(t *testing.T) {
	api := &Auth{ClientId: os.Getenv(clientIdEnvKey), ClientSecret: os.Getenv(clientSecretEnvKey)}

	authUrl := api.Authenticate(redirectUri)
	correctUrl := fmt.Sprintf("https://foursquare.com/oauth2/authenticate?client_id=%s&response_type=code&redirect_uri=%s", api.ClientId, redirectUri)

	if authUrl != correctUrl {
		t.Errorf("Invalid url: %v", authUrl)
	}
}
