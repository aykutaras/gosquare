package gosquare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Profile struct {
	Meta          Meta          `json:"meta"`
	Notifications Notifications `json:"notifications"`
	Response      UserResponse  `json:"response"`
}

type UserResponse struct {
	User User `json:"user"`
}

func (profile *Profile) Get(accessToken string) {
	res, err := http.Get(fmt.Sprintf("%s/self?oauth_token=%s&v=%s&m=foursquare", usersBaseUri, accessToken, version))
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
}
