package gosquare

import (
	"fmt"
	"log"
	"net/http"
)

const Uri = "localhost:4001"

func SendLinkHandler(w http.ResponseWriter, r *http.Request) {
	authUri := GetAuthenticationUrl(Uri)
	fmt.Fprintf(w, "<a href='%s'>Connect to Foursquare</a>", authUri)
}

func ListenForCodeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := fmt.Sprintf("%s", query.Get("code"))

	accessToken := GetAccessToken(Uri, code)
	checkIns := RetrieveFromFoursquare(accessToken)

	fmt.Fprintf(w, "%s", checkIns)
}

func InitHttpService() {
	http.HandleFunc("/", SendLinkHandler)
	http.HandleFunc("/code", ListenForCodeHandler)
	err := http.ListenAndServe("localhost:4001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
