package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ApiController struct {

}


func (con ApiController) Index(c *gin.Context) {

	//设置session
	session := sessions.Default(c)
	session.Set("username","张三11")
	session.Save()


	//设置cookie 二级域名
	c.SetCookie("username", "张三cookie", 6000, "/", ".test.com", false, false)
	c.String(200, "api index")
}

func (con ApiController) UserList(c *gin.Context) {

	session := sessions.Default(c)

	fmt.Print(session.Get("username"))

	username, _ := c.Cookie("username")
	c.String(200, "api UserList"+username)
}
