package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	md "github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"log"
	"strconv"
	"strings"
	"tasks/types"
	"time"
)

type Database struct {
	database *sql.DB
}

var db Database
var err error
var taskStatus map[string]int

func (db Database) begin() (tx *sql.Tx) {
	tx, err = db.database.Begin()
	if err != nil {
		log.Println(err)
	}
	return tx
}

func (db Database) prepare(query string) (stmt *sql.Stmt) {
	stmt, err = db.database.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	return stmt
}

func (db Database) query(q string, args ...interface{}) (rows *sql.Rows) {
	rows, err = db.database.Query(q, args...)
	if err != nil {
		log.Println(err)
	}
	return rows

}

func init() {
	db.database, err = sql.Open("mysql", "root@/tasks?parseTime=True&loc=Local")
	taskStatus = map[string]int{"COMPLETE": 1, "PENDING": 2, "DELETE": 3}
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	db.database.Close()
}

func GetTasks(username, status, category string) (types.Context, error) {
	var (
		ctx         types.Context
		tasks       []types.Task
		task        types.Task
		TaskCreated time.Time
		getTaskSQL  string
		rows        *sql.Rows
		err         error
	)

	comments, err := GetComments(username)
	if err != nil {
		return ctx, err
	}

	basicSQL := "select t.id, title, content, created_date, priority, case when c.name is null then 'NA' else c.name end from task t," +
		" status s, user u left outer join  category c on c.id=t.cat_id where u.username=? and s.id=t.task_status_id and u.id=t.user_id"
	if category == "" {
		switch status {
		case "pending":
			getTaskSQL = basicSQL + " and s.status='PENDING' and t.hide!=1"
		case "deleted":
			getTaskSQL = basicSQL + " and s.status='DELETED' and t.hide!=1"
		case "completed":
			getTaskSQL = basicSQL + " and s.status='COMPLETE' and t.hide!=1"
		}
		getTaskSQL += " t.created_date ASC"
		rows = db.query(getTaskSQL, username, username)
	} else {
		status = category
		if category == "UNCATEGORIZED" {
			getTaskSQL = "select t.id, title, content, created_date, priority, 'UNCATEGORIZED' from task t, status s, " +
				"user u where u.username=? and s.id=t.task_status_id and u.id=t.user_id and t.cat_id=0  and  s.status='PENDING'  " +
				"order by priority desc, created_date asc, finish_date asc"
			rows = db.query(getTaskSQL, username)
		} else {
			getTaskSQL = basicSQL + " and name = ?  and  s.status='PENDING'  order by priority desc, created_date asc, finish_date asc"
			rows = db.query(getTaskSQL, username, category)
		}
	}
	getTaskSQL = "select id,title,content,created_date from task;"
	rows = db.query(getTaskSQL)

	defer rows.Close()
	for rows.Next() {
		task = types.Task{}

		err = rows.Scan(&task.Id, &task.Title, &task.Content, &TaskCreated, &task.Priority, &task.Category)

		taskCompleted := 0
		totalTasks := 0

		if strings.HasPrefix(task.Content, "- [") {
			for _, value := range strings.Split(task.Content, "\n") {
				if strings.HasPrefix(value, "- [x]") {
					taskCompleted += 1
				}
				totalTasks += 1
			}
			task.CompletedMsg = strconv.Itoa(taskCompleted) + " complete out of " + strconv.Itoa(totalTasks)
		}

		task.ContentHTML = template.HTML(md.Markdown([]byte(task.Content)))
		// TaskContent = strings.Replace(TaskContent, "\n", "<br>", -1)
		if err != nil {
			log.Println(err)
		}

		if comments[task.Id] != nil {
			task.Comments = comments[task.Id]
		}

		TaskCreated = TaskCreated.Local()
		task.Created = TaskCreated.Format("Jan 2 2006")

		tasks = append(tasks, task)
	}

	ctx = types.Context{Tasks: tasks, Navigation: status}
	return ctx, nil
}

func GetComments(username string) (map[int][]types.Comment, error) {
	commentMap := make(map[int][]types.Comment)

	var taskID int
	var comment types.Comment
	var created time.Time

	userID, err := GetUserId(username)
	if err != nil {
		return commentMap, err
	}
	stmt := "select c.id, c.taskID, c.content, c.created, u.username from comments c, task t, user u where t.id=c.taskID and c.user_id=t.user_id and t.user_id=u.id and u.id=?"
	rows := db.query(stmt, userID)

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&comment.ID, &taskID, &comment.Content, &created, &comment.Username)
		comment.Content = string(md.Markdown([]byte(comment.Content)))
		if err != nil {
			return commentMap, err
		}
		created = created.Local()
		comment.Created = created.Format("Jan 2 2006 15:04:05")
		commentMap[taskID] = append(commentMap[taskID], comment)
	}
	return commentMap, nil

}

func AddTask(title, content string) error {
	query := "insert into task(title,content,created_date,last_modified_at) values(?,?,now(),now())"
	addTaskSQL := db.prepare(query)

	tx := db.begin()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = tx.Stmt(addTaskSQL).Exec(title, content)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}
	return nil
}

func taskQuery(stmt string, args ...interface{}) error {
	SQL := db.prepare(stmt)
	tx := db.begin()
	_, err = tx.Stmt(SQL).Exec(args...)
	if err != nil {
		log.Println(err)
		tx.Rollback()
	} else {
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("commit successful")
	}
	return err
}
