package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"

var My *sql.DB // MySQL

func init() {
	var err error
	My, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = My.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
