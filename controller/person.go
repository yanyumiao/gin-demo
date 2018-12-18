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
	p.Id, _ = strconv.Atoi(id) // Notice: int(id) not work!
	person, err := p.GetPerson()
	// TODO
	if err != nil {
		var res MyErr
		res.Code = -1
		res.Desc = "error"
		log.Println("Error:", res.Desc)
		c.JSON(200, res)
	} else {
		c.JSON(200, person)
	}
}

func GetAllPerson(c *gin.Context) {
	var p model.Person
	persons, _ := p.GetAllPerson()
	c.JSON(200, persons)
}

func DelPerson(c *gin.Context) {
	var p model.Person
	var id = c.Query("id")
	p.Id, _ = strconv.Atoi(id)
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
		// catch err, so you can see err msg in console
		log.Fatalln(err)
	}
	c.String(200, "ok")
}
