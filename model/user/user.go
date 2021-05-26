package user

import (
	"log"
	"mydemo/model"
)

type User struct {
	Id       int
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}

func Signup(userName string, password string) bool {
	db := model.Conn()

	stmtIns, err := db.Prepare("INSERT INTO user VALUES(NULL, ?, ?)") // ? = placeholder
	if err != nil {
		return false
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err2 := stmtIns.Exec(userName, password)
	if err2 != nil {
		return false
	}

	return true
}

func Login(userName string, password string) User {
	db := model.Conn()

	stmt, err := db.Prepare("SELECT id, username FROM user WHERE username = ? AND password = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var name string
	var id int
	err = stmt.QueryRow(userName, password).Scan(&id, &name)
	if err != nil {
		return User{}
	} else {
		return User{Id: id, UserName: name}
	}
}
