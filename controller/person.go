package controller

import "github.com/gin-gonic/gin"

//import "net/http"
import model "gin-demo/model"
import "strconv"
import "log"

func AddPerson(c *gin.Context) {
	var p model.Person
	var firstname = c.Query("firstname")
	var lastname = c.Query("lastname")
	p.FirstName = firstname
	p.LastName = lastname
	var id, _ = p.AddPerson() // TODO check err
	p.Id = int(id)            // TODO cannot use id (type int64) as type int ?
	c.JSON(200, p)
}

func GetPerson(c *gin.Context) {
	var p model.Person
	var id = c.Param("id") // Notice: string
	// Notice: 强制类型转换
	// TODO 总结Go的类型转换
	// 参考:http://blog.sina.com.cn/s/blog_537517170102xkc2.html
	p.Id, _ = strconv.Atoi(id) // Notice: int(id) not work!
	person := p.GetPerson()
	c.JSON(200, person)
}

func GetAllPerson(c *gin.Context) {
	//var persons []model.Person // slice
	var p model.Person
	persons, _ := p.GetAllPerson()
	c.JSON(200, persons)
}

func DelPerson(c *gin.Context) {
	var p model.Person
	var id = c.Query("id")
	p.Id, _ = strconv.Atoi(id) // Noitce: return double
	p.DelPerson()
	c.String(200, "ok")
}

func UpdatePerson(c *gin.Context) {
	var p model.Person
	var id = c.Query("id")
	var firstname = c.Query("firstname")
	var lastname = c.Query("lastname")
	p.Id, _ = strconv.Atoi(id)
	p.FirstName = firstname
	p.LastName = lastname
	_, err := p.UpdatePerson()
	if err != nil {
		// 如果这里不做判断 db 里面的sql写错了 是捕捉不到错误的
		// 在命令行 控制台 可以看到错误信息
		log.Fatalln(err)
	}
	c.String(200, "ok")
}
