package main

import (
	"log"
	"net/http"
	"tasks/config"
	"tasks/views"
)

func main() {

	conf := config.ReadConfigFile("config.json")

	PORT := conf.ServerPort
	log.Println("listening on ", PORT)
	views.PopulateTemplates()
	//router := httprouter.New()
	//router.POST("/login", views.Login)


	http.HandleFunc("/complete/", views.CompleteTaskFunc)
	http.HandleFunc("/delete/", views.DeleteTaskFunc)
	http.HandleFunc("/deleted/", views.ShowTrashTaskFunc)
	http.HandleFunc("/trash/", views.TrashTaskFunc)
	http.HandleFunc("/edit/", views.EditTaskFunc)
	http.HandleFunc("/completed/", views.ShowCompletedTaskFunc)
	http.HandleFunc("/add/", views.AddTaskFunc)
	http.HandleFunc("/restore/", views.RestoreTaskFunc)
	http.HandleFunc("/update/", views.UpdateTaskFunc)
	http.HandleFunc("/search/", views.SearchTaskFunc)
	http.HandleFunc("/login/", views.Login)
	http.HandleFunc("/logout/", views.RequiresLogin(views.Logout))

	http.HandleFunc("/register", views.PostRegister)
	http.HandleFunc("/admin", views.HandleAdmin)
	http.HandleFunc("/add_user", views.PostAddUser)
	http.HandleFunc("/change", views.PostChange)

	http.HandleFunc("/", views.RequiresLogin(views.ShowAllTasksFunc))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/Users/mmt6198/go/src/tasks/public"))))
	log.Fatal(http.ListenAndServe(PORT, nil))

}
