package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mydemo/config"
)

func Conn() *sql.DB {
	conf := config.Config().Db

	db, err := sql.Open(
		"mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			conf["user"], conf["password"], conf["ip"], conf["port"], conf["db_name"]))
	if err != nil {
		panic(err.Error())
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return db
}
