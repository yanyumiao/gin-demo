package router

import "github.com/gin-gonic/gin"
import . "gin-demo/controller" // Notice: ".", 点操作 导入包后 调用这个包的函数时 可以省略包名

func InitRouter() *gin.Engine {
	router := gin.Default()
	// router
	// hello.go
	router.GET("/", Hello)
	router.GET("/ping", Pong)
	router.GET("/async", Async)
	// person.go
	router.GET("/addperson", AddPerson)     // ?firstname=a&lastname=b
	router.GET("/getperson/:id", GetPerson) // /getperson/1
	router.GET("/getallperson", GetAllPerson)
	router.GET("/delperson", DelPerson)       // ?id=4
	router.GET("/updateperson", UpdatePerson) // ?id=4&firstname=a&lastname=b
	//router.StaticFile("/favicon.ico", "./resource/favicon.ico") // 相对路径 相对main.go
	return router
}
