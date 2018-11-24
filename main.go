package main

import "gin-demo/router"
import db "gin-demo/database"

func main() {
	defer db.SqlDB.Close()
	router := router.InitRouter()
	router.Run(":10086")
}
