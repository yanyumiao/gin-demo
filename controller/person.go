package controller

//import "net/http"
//import "log"
import "github.com/gin-gonic/gin"
import "gin-demo/model"
import "strconv"

func AddPerson(c *gin.Context) {
	var p model.Person
	var firstname = c.Query("firstname")
	var lastname = c.Query("lastname")
	p.FirstName = firstname
	p.LastName = lastname
	var id, err = p.Add()
	if err != nil {
		c.JSON(200, res(500, "error", nil))
	} else {
		p.Id = int(id) // Notice
		c.JSON(200, res(0, "ok", p))
	}
}

func GetPerson(c *gin.Context) {
	var p model.Person
	var id = c.Param("id")     // Notice: string
	p.Id, _ = strconv.Atoi(id) // Notice: int(id) not work!
	person, err := p.Get()
	if err != nil {
		//log.Println("Error:", res.Desc)
		c.JSON(200, res(500, "error", nil))
	} else {
		c.JSON(200, res(0, "ok", person))
	}
}

func GetAllPerson(c *gin.Context) {
	var p model.Person
	persons, _ := p.GetAll()
	c.JSON(200, res(0, "ok", persons))
}

func DelPerson(c *gin.Context) {
	var p model.Person
	var id = c.Query("id")
	p.Id, _ = strconv.Atoi(id)
	p.Del()
	c.JSON(200, res(0, "ok", nil))
}

func UpdatePerson(c *gin.Context) {
	var p model.Person
	var id = c.Query("id")
	var firstname = c.Query("firstname")
	var lastname = c.Query("lastname")
	p.Id, _ = strconv.Atoi(id)
	p.FirstName = firstname
	p.LastName = lastname
	p.Update()
	c.JSON(200, res(0, "ok", nil))
}
