package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseUri = "https://api.foursquare.com/v2/users"
const version = "20150309"

func (users *Users) GetCheckIns(accessToken string) {
	res, err := http.Get(fmt.Sprintf("%s/self/checkins?oauth_token=%s&v=%s&m=swarm", baseUri, accessToken, version))
	if err != nil {
		log.Fatal(err)
	}

	byteCheckIns, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	//checkIns := &UserCheckIns{}
	if err := json.Unmarshal(byteCheckIns, &users.CheckIns); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", users.CheckIns)
}
