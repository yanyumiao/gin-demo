package main

import "gin-demo/router"

func main() {
	router := router.InitRouter()
	router.Run(":10086")
}
