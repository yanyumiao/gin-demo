package main

import "gin/router"

func main() {
	router := router.InitRouter()
	router.Run(":10086")
}
