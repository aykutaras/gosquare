package oauth2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseUri = "https://foursquare.com/oauth2"

func (api *Auth) Authenticate(redirectUri string) string {
	return fmt.Sprintf("%s/authenticate?client_id=%s&response_type=code&redirect_uri=http://%s/code", baseUri, api.ClientId, redirectUri)
}

func (api *Auth) GetAccessToken(redirectUri, code string) {
	res, err := http.Get(fmt.Sprintf("%s/access_token?client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=http://%s/code&code=%s", baseUri, api.ClientId, api.ClientSecret, redirectUri, code))
	if err != nil {
		log.Fatal(err)
	}

	byteToken, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	var accessToken map[string]string
	if err := json.Unmarshal(byteToken, &accessToken); err != nil {
		log.Fatal(err)
	}

	api.AccessToken = accessToken["access_token"]
}
