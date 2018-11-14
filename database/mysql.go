package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 大写字母开头
var SqlDB *sql.DB

// init
func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
