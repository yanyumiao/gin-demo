package router

import "github.com/gin-gonic/gin"
import . "gin/controller" // Notice: ".", 点操作 导入包后 调用这个包的函数时 可以省略包名

func InitRouter() *gin.Engine {
	router := gin.Default()
	// router
	router.GET("/", Hello)
	router.GET("/test", Test)
	//
	return router
}
