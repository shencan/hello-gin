package api

import "github.com/gin-gonic/gin"

type ApiController struct {
}

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "api index")
}

func (con ApiController) UserList(c *gin.Context) {
	c.String(200, "api UserList")
}
