package router

import "github.com/gin-gonic/gin"
import "net/http"

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	return router
}
