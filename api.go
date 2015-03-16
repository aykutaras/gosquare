package gosquare

import (
	"github.com/aykutaras/gosquare/oauth2"
	"github.com/aykutaras/gosquare/users"
)

type Api struct {
	Auth  *oauth2.Auth
	Users *users.Users
}
