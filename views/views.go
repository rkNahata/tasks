package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"tasks/db"
	"tasks/sessions"
	"time"
)

var homeTemplate *template.Template
var deletedTemplate *template.Template
var completedTemplate *template.Template
var editTemplate *template.Template
var searchTemplate *template.Template
var templates *template.Template
var loginTemplate *template.Template

var message string

func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		username := sessions.GetCurrentUserName(r)
		ctx, err := db.GetTasks(username, "pending", "")
		categories := db.GetCategories(username)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusInternalServerError)
		} else {
			if message != "" {
				ctx.Message = message
			}
			ctx.CSRFToken = "tasks"
			ctx.Categories = categories
			message = ""
			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := http.Cookie{Name: "csrf", Value: "tasks", Expires: expiration,}
			http.SetCookie(w, &cookie)
			homeTemplate.Execute(w, ctx)
		}
	}

}


func PostChange(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("will allow changing password"))

}

func PostAddUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("will allow creating new user"))

}

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("will allow creating new user"))

}

func PostRegister(w http.ResponseWriter, r *http.Request) {

}

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	a := fmt.Sprintf("URL string : %v\n", r.URL.String())
	w.Write([]byte(a))
}

func DeleteTaskFunc(w http.ResponseWriter, r *http.Request) {

}

func ShowTrashTaskFunc(w http.ResponseWriter, r *http.Request) {

}

func TrashTaskFunc(w http.ResponseWriter, r *http.Request) {

}

func EditTaskFunc(w http.ResponseWriter, r *http.Request) {

}

func ShowCompletedTaskFunc(w http.ResponseWriter, r *http.Request) {

}

func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	title := "some title"
	content := "some content"
	err := db.AddTask(title, content)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("added task"))
}

func RestoreTaskFunc(w http.ResponseWriter, r *http.Request) {

}

func UpdateTaskFunc(w http.ResponseWriter, r *http.Request) {

}
func SearchTaskFunc(w http.ResponseWriter, r *http.Request) {

}
