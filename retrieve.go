package foursquash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type UserCheckIns struct {
	Meta          Meta          `json:"meta"`
	Notifications Notifications `json:"notifications"`
	Response      Response      `json:"response"`
}

type Meta struct {
	Code         int    `json:"code"`
	ErrorType    string `json:"errorType"`
	ErrorDetails string `json:"errorDetails"`
}

type Notifications []Notification

type Notification struct {
	Type string      `json:"type"`
	Item interface{} `json:"item"`
}

type Response struct {
	CheckIns CheckIns `json:"checkins"`
}

type CheckIns struct {
	Count int           `json:"count"`
	Items []CheckInItem `json:"items"`
}

type CheckInItem struct {
	Id               string   `json:"id"`
	CreatedAt        int      `json:"createdAt"`
	Type             string   `json:"type"`
	Shout            string   `json:"shout"`
	Entities         []Entity `json:"entities"`
	TimeZoneOffset   int      `json:"timeZoneOffset"`
	DisplayGeo       Geo      `json:"displayGeo"`
	ExactContextLine string   `json:"exactContextLine"`
	With             []User   `json:"with"`
	Venue            Venue    `json:"venue"`
	Likes            Likes    `json:"likes"`
	Like             bool     `json:"like"`
	Sticker          Sticker  `json:"sticker"`
	Photos           Photos   `json:"photos"`
	Posts            Posts    `json:"posts"`
	Comments         Comments `json:"comments"`
	Source           Source   `json:"source"`
}

type Entity struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Indices []int  `json:indices"`
}

type Geo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id           string    `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Gender       string    `json:"gender"`
	Relationship string    `json:"relationship"`
	Photo        UserPhoto `json:"photo"`
}

type Venue interface {
}

type Likes struct {
	Count  int         `json:"count"`
	Groups []UserGroup `json:"groups"`
}

type UserGroup struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
	Items []User `json:"items"`
}

type UserPhoto struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Sticker interface {
}

type Photos struct {
	Count int     `json:"count"`
	Items []Photo `json:"items"`
}

type Photo struct {
	Id         string `json:"id"`
	CreatedAt  int    `json:"createdAt"`
	Source     Source `json:"source"`
	Prefix     string `json:"prefix"`
	Suffix     string `json:"suffix"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	User       User   `json:"user"`
	Visibility string `json:"visibility"`
}

type Posts struct {
	Count     int `json:"count"`
	TextCount int `json:"textCount"`
}

type Comments struct {
	Count int `json:"count"`
}

type Source struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

const Uri = "http://localhost:4001"

const ApiOwner = "Aykut Aras"
const ClientId = "32R11TKNCLDMOZQQXLDGKX1J3OCTDJTPI21HWL35V4LF1NCC"
const ClientSecret = "SUSVS34WO3ANGHYSYOSA5HOPKWB0M0E1LN3SICZKFCF1ZZIZ"

func SendLinkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href='https://foursquare.com/oauth2/authenticate?client_id=%s&response_type=code&redirect_uri=%s/code'>Connect</a>", ClientId, Uri)
}

func ListenForCodeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := fmt.Sprintf("%s", query.Get("code"))

	accessToken := GetAccessToken(code)
	checkIns := RetrieveFromFoursquare(accessToken)

	fmt.Fprintf(w, "%s", checkIns)
}

func GetAccessToken(code string) string {
	res, err := http.Get(fmt.Sprintf("https://foursquare.com/oauth2/access_token?client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s/code&code=%s", ClientId, ClientSecret, Uri, code))

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
		fmt.Println("checkins")
		log.Fatal(err)
	}

	fmt.Println(checkIns)
	return fmt.Sprintf("%s", byteCheckIns)
}

func InitHttpService() {
	http.HandleFunc("/", SendLinkHandler)
	http.HandleFunc("/code", ListenForCodeHandler)
	err := http.ListenAndServe("localhost:4001", nil)
	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
}
