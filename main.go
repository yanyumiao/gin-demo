package main

import "gindemo/router"
import db "gindemo/database"

func main() {
	defer db.SqlDB.Close()
	router := router.InitRouter()
	router.Run(":10086")
}
