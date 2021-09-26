package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRoutersInit(r *gin.Engine)  {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "值：%v", "你好api")
		})

		apiRouters.GET("/userList", func(c *gin.Context) {
			c.String(http.StatusOK, "值：%v", "xx")
		})
	}
}