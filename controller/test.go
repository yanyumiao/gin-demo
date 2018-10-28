package controller

import "github.com/gin-gonic/gin"
import "net/http"

func Test(c *gin.Context) {
	c.String(http.StatusOK, "test")
}
