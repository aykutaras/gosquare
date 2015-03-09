package foursquash

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const ApiOwner = "Aykut Aras"
const ClientId = "32R11TKNCLDMOZQQXLDGKX1J3OCTDJTPI21HWL35V4LF1NCC"
const ClientSecret = "SUSVS34WO3ANGHYSYOSA5HOPKWB0M0E1LN3SICZKFCF1ZZIZ"

const GetUserCheckIns = ""

func SendLink() {
	fmt.Printf("https://foursquare.com/oauth2/authenticate?client_id=%s&response_type=code&redirect_uri=http://localhost:4001/code", ClientId)
}

func ListenForCode() {
	http.HandleFunc("/code", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Query())
	})
}

func GetAccessToken(code string) string {
	res, err := http.Get(fmt.Sprintf("https://foursquare.com/oauth2/access_token?client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s&code=%s", ClientId, ClientSecret, "http://localhost:4001/code", code))

	if err != nil {
		log.Fatal(err)
	}

	token, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", token)
	return fmt.Sprintf("%s", token)
}

func RetrieveFromFoursquare(accessToken string) {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get(fmt.Sprintf("https://api.foursquare.com/v2/users/self/checkins?oauth_token=%s&v=20150309&m=swarm", accessToken))

		if err != nil {
			log.Fatal(err)
		}

		checkIns, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s", checkIns)
	})
}

func InitHttpService() {
	SendLink()
	//ListenForCode()
	//accessToken := GetAccessToken("SR3U0YXBRH52GLLOPKZRVPRYOVS5RPTDXX54DXPLW3KBNFPK")
	RetrieveFromFoursquare("DKXXGSXF35LRBERMJ0XYOBCASR3JSECERVQ4ELJTYVMG0NW1")

	err := http.ListenAndServe("localhost:4001", nil)
	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
}
