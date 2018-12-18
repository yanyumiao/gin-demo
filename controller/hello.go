package controller

import "github.com/gin-gonic/gin"
import "net/http"
import "log"
import "time"

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func Async(c *gin.Context) {
	var c_copy = c.Copy()
	// 异步
	go func() {
		time.Sleep(5 * time.Second)
		// you can see it on console
		log.Println("Done! in path" + c_copy.Request.URL.Path)
	}()
	c.String(http.StatusOK, "ok")
}

func Pong(c *gin.Context) {
	// type H map[string]interface{}
	// H is a shortcut for map[string]interface{}
	// interface{}接口类型 表示value可以是任何数据类型
	c.JSON(200, gin.H{"ret": 0, "msg": "ok"})
	//c.JSON(200, map[string]interface{}{"ret": 0, "msg": "pong"})
}
