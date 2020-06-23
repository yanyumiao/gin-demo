package controller

import "github.com/gin-gonic/gin"
import "net/http"
import "log"
import "time"
import db "gin-demo/database"

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

// redis set
func Set(c *gin.Context) {
	con := db.Pool.Get()
	defer con.Close()
	con.Do("set", "test", "abc")
	//re, err :=con.Do("set", "test", "abc")
	//fmt.Print(con, re, err)
	c.JSON(200, gin.H{"ret": 0, "msg": "ok"})
}
