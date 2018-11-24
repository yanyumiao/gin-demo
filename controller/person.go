package controller

import "github.com/gin-gonic/gin"
import "net/http"
import model "gin/model"
import "strconv"

func GetPerson(c *gin.Context) {
	var p model.Person
	var id = c.Param("id") // Notice: string
	// Notice: 强制类型转换
	// TODO 总结Go的类型转换
	// 参考:http://blog.sina.com.cn/s/blog_537517170102xkc2.html
	p.Id, _ = strconv.Atoi(id)
	person := p.Get()
	c.String(http.StatusOK, person.FirstName)
}
