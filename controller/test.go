package controller

import "github.com/gin-gonic/gin"
import "net/http"

func Test(c *gin.Context) {
	c.String(http.StatusOK, "test")
}
func Pong(c *gin.Context) {
	// type H map[string]interface{}
	// H is a shortcut for map[string]interface{}
	// interface{}接口类型 表示value可以是任何数据类型
	c.JSON(200, map[string]interface{}{"code": 1, "msg": "pong"})
	//c.JSON(200, gin.H{"code": 1, "msg": "pong"})
}
