package admin

import "github.com/gin-gonic/gin"

type baseController struct {
}

func (con baseController) success(c *gin.Context) {
	c.String(200, "成功")
}

func (con baseController) error(c *gin.Context) {
	c.String(200, "失败")
}
