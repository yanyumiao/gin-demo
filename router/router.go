package router

import "github.com/gin-gonic/gin"
import . "gin-demo/controller"

func InitRouter() *gin.Engine {
	router := gin.Default()
	// router
	// hello.go
	router.GET("/", Hello)
	router.GET("/hello", Hello)
	router.GET("/ping", Pong)
	router.GET("/async", Async)
	router.GET("/redis/set", Set)
	// person.go
	router.GET("/person/add", AddPerson) // http://127.0.0.1:10086/person/add?firstname=a&lastname=b
	router.GET("/person/get", GetPerson) // http://127.0.0.1:10086/person/get?id=1
	router.GET("/person/all", GetAllPerson)
	router.GET("/person/del", DelPerson)       // http://127.0.0.1:10086/person/del?id=4
	router.GET("/person/update", UpdatePerson) // http://127.0.0.1:10086/person/update?id=4&firstname=a&lastname=b
	// router.StaticFile("/favicon.ico", "./resource/favicon.ico")
	return router
}
