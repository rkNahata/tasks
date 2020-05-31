package db

import "log"

func CreateUser(username, password, email string) error {
	err := taskQuery("insert into users(username,password,email)values(?,?,?)", username, password, email)
	return err
}

func IsValidUser(username, password string) bool {
	var passwordFromdb string
	getPasswordFromdb := "select password from users where username=?"
	rows := db.query(getPasswordFromdb, username)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&passwordFromdb)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	if password == passwordFromdb {
		return true
	}
	return false
}


func GetUserId(username string)(int,error){
	var userId int
	userSQL := "select id from users where username=?"
	rows := db.query(userSQL,username)
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&userId)
		if err !=nil{
			log.Println(err)
			return -1,err
		}
	}
	return userId,nil
}