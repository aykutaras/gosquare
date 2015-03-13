package users

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

const BaseUri = "https://api.foursquare.com/v2/users"
const Version = "20150309"

func GetCheckIns(accessToken string) string {
	res, err := http.Get(fmt.Sprintf("%s/self/checkins?oauth_token=%s&v=%s&m=swarm", BaseUri, accessToken, Version))
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
