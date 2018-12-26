package main

import "gin-demo/router"
import db "gin-demo/database"

func main() {
	defer db.My.Close()
	router := router.InitRouter()
	router.Run(":10086")
}
