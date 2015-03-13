package oauth2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const ApiOwner = "Aykut Aras"
const ClientId = "32R11TKNCLDMOZQQXLDGKX1J3OCTDJTPI21HWL35V4LF1NCC"
const ClientSecret = "SUSVS34WO3ANGHYSYOSA5HOPKWB0M0E1LN3SICZKFCF1ZZIZ"

const BaseUri = "https://foursquare.com/oauth2"

func Authenticate(redirectUri string) string {
	authUrl := fmt.Sprintf("%s/authenticate?client_id=%s&response_type=code&redirect_uri=http://%s/code", BaseUri, ClientId, redirectUri)
	return authUrl
}

func AccessToken(redirectUri, code string) string {
	res, err := http.Get(fmt.Sprintf("%s/access_token?client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=http://%s/code&code=%s", BaseUri, ClientId, ClientSecret, redirectUri, code))
	if err != nil {
		log.Fatal(err)
	}

	accessToken := &AccessTokenResponse{}
	byteToken, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	if err := json.Unmarshal(byteToken, &accessToken); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprint(accessToken.AccessToken)
}
