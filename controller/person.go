package controller

import "github.com/gin-gonic/gin"
import "net/http"
import model "gin/model"
import "strconv"

func AddPerson(c *gin.Context) {
	var p model.Person
	var firstname = c.Query("firstname")
	var lastname = c.Query("lastname")
	p.FirstName = firstname
	p.LastName = lastname
	var id, _ = p.Add() // TODO check err
	p.Id = int(id)      // TODO cannot use id (type int64) as type int ?
	c.JSON(200, p)
}

func GetPerson(c *gin.Context) {
	var p model.Person
	var id = c.Param("id") // Notice: string
	// Notice: 强制类型转换
	// TODO 总结Go的类型转换
	// 参考:http://blog.sina.com.cn/s/blog_537517170102xkc2.html
	p.Id, _ = strconv.Atoi(id) // Notice: int(id) not work!
	person := p.Get()
	c.String(http.StatusOK, person.FirstName)
}

func GetAllPerson(c *gin.Context) {
	//var persons []model.Person // slice
	var p model.Person
	persons, _ := p.GetAll()
	c.JSON(200, persons)
}

func DelPerson(c *gin.Context) {
	var p model.Person
	var id = c.Query("id")
	p.Id, _ = strconv.Atoi(id) // Noitce: return double
	p.Del()
	c.String(200, "ok")
}
