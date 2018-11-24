package router

import "github.com/gin-gonic/gin"
import . "gindemo/controller" // Notice: ".", 点操作 导入包后 调用这个包的函数时 可以省略包名

func InitRouter() *gin.Engine {
	router := gin.Default()
	// router
	router.GET("/", Hello)
	router.GET("/test", Test)
	router.GET("/ping", Pong)           // JSON
	router.GET("/addperson", AddPerson) // ?firstname=a&lastname=b
	router.GET("/getperson/:id", GetPerson)
	router.GET("/getallperson", GetAllPerson)
	router.GET("/delperson", DelPerson) // ?id=4
	//router.StaticFile("/favicon.ico", "./resource/favicon.ico") // 相对路径 相对main.go
	//
	return router
}
