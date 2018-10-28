package controller

import "github.com/gin-gonic/gin"
import "net/http"

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
