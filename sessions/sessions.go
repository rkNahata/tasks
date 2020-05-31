package sessions

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("secret-password"))
var session *sessions.Session

func IsLoggedIn(r *http.Request) bool {
	session, err := Store.Get(r, "session")
	if err != nil && session.Values["loggedin"] == "true" {
		return true
	}
	return false
}

func GetCurrentUserName(r *http.Request) string {
	session, err := Store.Get(r, "session")
	if err != nil {
		return session.Values["username"].(string)
	}
	return ""
}
