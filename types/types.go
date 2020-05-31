package types

import (
	"html/template"
)

type Task struct {
	Id           int
	Title        string
	Content      string
	ContentHTML  template.HTML
	Created      string
	Priority     string
	Category     string
	Referer      string
	Comments     []Comment
	IsOverdue    bool
	IsHidden     int
	CompletedMsg string
}

type Tasks []Task

type Context struct {
	Tasks      []Task
	Navigation string
	Search     string
	Message    string
	CSRFToken  string
	Categories []CategoryCount
	Referer    string
}

type Comment struct {
	ID       int
	Content  string
	Created  string
	Username string
}

type CategoryCount struct {
	Name  string
	Count int
}

type Status struct {
	StatusCode int
	Message    string
}

type Category struct {
	ID      int
	Name    string
	Created string
}

type Categories []Category
