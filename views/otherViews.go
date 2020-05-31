package views

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func PopulateTemplates(){
	var allFiles []string
	tempatesDir := "./templates/"
	files,err := ioutil.ReadDir(tempatesDir)
	if err !=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	for _,file := range files{
		fileName := file.Name()
		if strings.HasSuffix(fileName,".html"){
			allFiles = append(allFiles,tempatesDir+fileName)
		}
	}

	templates,err := template.ParseFiles(allFiles...)
	if err !=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	homeTemplate = templates.Lookup("home.html")
	loginTemplate = templates.Lookup("login.html")
	editTemplate = templates.Lookup("edit.html")
	searchTemplate = templates.Lookup("search.html")
	completedTemplate = templates.Lookup("completed.html")
	deletedTemplate = templates.Lookup("deleted.html")
}
