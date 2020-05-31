package views

import (
	"log"
	"net/http"
	"tasks/db"
	"tasks/sessions"
)

func RequiresLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !sessions.IsLoggedIn(r) {
			http.Redirect(w, r, "/login/", 302)
			return
		}
		handler(w, r)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "session")
	if err == nil {
		if session.Values["loggedin"] != "false" {
			session.Values["loggedin"] = "false"
			session.Save(r, w)
		}
		http.Redirect(w, r, "/login", 302)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	switch r.Method {
	case "GET":
		loginTemplate.Execute(w, nil)
	case "POST":
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if (username != "" && password != "") && (db.IsValidUser(username, password)) {
			session.Values["loggedin"] = "true"
			session.Values["username"] = username
			session.Save(r, w)
			log.Println("user ",username, " is authenticated")
			http.Redirect(w, r, "/", 302)
			return
		}
		log.Println("invalid user ", username)
		loginTemplate.Execute(w, nil)
	default:
		http.Redirect(w, r, "/login/", http.StatusUnauthorized)

	}

}
