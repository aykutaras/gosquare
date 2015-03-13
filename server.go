package gosquare

import (
	"fmt"
	"github.com/aykutaras/gosquare/oauth2"
	"github.com/aykutaras/gosquare/users"
	"log"
	"net/http"
)

const Uri = "localhost:4001"

func SendLinkHandler(w http.ResponseWriter, r *http.Request) {
	authUri := oauth2.Authenticate(Uri)
	fmt.Fprintf(w, "<a href='%s'>Connect to Foursquare</a>", authUri)
}

func ListenForCodeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := fmt.Sprintf("%s", query.Get("code"))

	accessToken := oauth2.AccessToken(Uri, code)
	checkIns := users.GetCheckIns(accessToken)

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
