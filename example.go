package main

import "github.com/gin-gonic/gin"
import "net/http"

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
	// query string parameters
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "guset")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})
	// listen and serve on 0.0.0.0:8080
	// router.Run(":3000") for a hard coded port
	r.Run(":10086")
}
