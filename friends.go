package gosquare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Friends struct {
	Meta          Meta            `json:"meta"`
	Notifications Notifications   `json:"notifications"`
	Response      FriendsResponse `json:"response"`
}

type FriendsResponse struct {
	Friends UserGroup `json:"friends"`
}

func (friends *Friends) Get(accessToken string) {
	res, err := http.Get(fmt.Sprintf("%s/self/friends?oauth_token=%s&v=%s&m=swarm", usersBaseUri, accessToken, version))
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
