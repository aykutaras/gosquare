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

func (checkIns *UserCheckIns) Get(accessToken string) {
	res, err := http.Get(fmt.Sprintf("%s/self/checkins?oauth_token=%s&v=%s&m=swarm", baseUri, accessToken, version))
	if err != nil {
		log.Fatal(err)
	}

	byteCheckIns, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	if err := json.Unmarshal(byteCheckIns, &checkIns); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", checkIns)
}

func (profile *Profile) Get(accessToken string) {
	res, err := http.Get(fmt.Sprintf("%s/self?oauth_token=%s&v=%s&m=foursquare", baseUri, accessToken, version))
	if err != nil {
		log.Fatal(err)
	}

	byteProfile, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	if err := json.Unmarshal(byteProfile, &profile); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", profile)
}

func (friends *Friends) Get(accessToken string) {
	res, err := http.Get(fmt.Sprintf("%s/self/friends?oauth_token=%s&v=%s&m=swarm", baseUri, accessToken, version))
	if err != nil {
		log.Fatal(err)
	}

	byteFriends, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	if err := json.Unmarshal(byteFriends, &friends); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", friends)
}
