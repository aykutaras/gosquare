package gosquare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UserCheckIns struct {
	Meta          Meta             `json:"meta"`
	Notifications Notifications    `json:"notifications"`
	Response      CheckInsResponse `json:"response"`
}

type CheckInsResponse struct {
	CheckIns CheckIns `json:"checkins"`
}

func (checkIns *UserCheckIns) Get(accessToken string) {
	res, err := http.Get(fmt.Sprintf("%s/self/checkins?oauth_token=%s&v=%s&m=swarm", usersBaseUri, accessToken, version))
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
}
