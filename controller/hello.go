package controller

import "github.com/gin-gonic/gin"
import "net/http"
import "log"
import "time"
import "github.com/gomodule/redigo/redis"

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func Async(c *gin.Context) {
	var c_copy = c.Copy() // Notice: copy context
	// async
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c_copy.Request.URL.Path)
	}()
	c.String(http.StatusOK, "ok")
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{"ret": 0, "msg": "ok"})
	// type H map[string]interface{}
	//c.JSON(200, map[string]interface{}{"ret": 0, "msg": "pong"})
}

func ConRedis(c *gin.Context) {
	con, err := redis.Dial("tcp", ":6379")
	if err != nil {
	    // handle error
		log.Println("error")
	}else{
		log.Println("ok")
	}
	defer con.Close()
	c.JSON(200, gin.H{"ret": 0, "msg": "ok"})
}