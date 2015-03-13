package gosquare

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

func GetAuthenticationUrl(uri string) string {
	authUrl := fmt.Sprintf("https://foursquare.com/oauth2/authenticate?client_id=%s&response_type=code&redirect_uri=http://%s/code", ClientId, uri)
	return authUrl
}

func GetAccessToken(uri, code string) string {
	res, err := http.Get(fmt.Sprintf("https://foursquare.com/oauth2/access_token?client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=http://%s/code&code=%s", ClientId, ClientSecret, uri, code))

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

func RetrieveFromFoursquare(accessToken string) string {
	res, err := http.Get(fmt.Sprintf("https://api.foursquare.com/v2/users/self/checkins?oauth_token=%s&v=20150309&m=swarm", accessToken))

	if err != nil {
		log.Fatal(err)
	}

	byteCheckIns, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	checkIns := &UserCheckIns{}
	if err := json.Unmarshal(byteCheckIns, &checkIns); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s", byteCheckIns)
}
