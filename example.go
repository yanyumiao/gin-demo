package main

import "net/http"
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// example
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// parameters in path
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})

	r.Run(":10086") // listen and serve on 0.0.0.0:8080
	// router.Run(":3000") for a hard coded port
}
