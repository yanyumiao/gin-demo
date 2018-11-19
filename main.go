package main

import "gin/router"
import db "gin/database"

func main() {
	defer db.SqlDB.Close()
	router := router.InitRouter()
	router.Run(":10086")
}
