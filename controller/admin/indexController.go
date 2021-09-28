package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {

	//username类型
	username, _ := c.Get("username")
	fmt.Println(username)

	//类型断言

	v, ok := username.(string)

	if ok == true{
		c.String(200, "用户列表："+v)
	}else {
		c.String(200, "用户列表用户获取失败")

	}
}
