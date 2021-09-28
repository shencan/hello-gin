package admin

import "github.com/gin-gonic/gin"

type UserController struct {
	baseController
}

func (con UserController) Index(c *gin.Context) {
	//c.String(200, "用户列表")
	con.success(c)
}

func (con UserController) Add(c *gin.Context) {
	c.String(200, "添加用户")
}

func (con UserController) Edit(c *gin.Context) {
	c.String(200, "修改用户")
}

func (con UserController) Delete(c *gin.Context) {
	c.String(200, "删除用户")
}
