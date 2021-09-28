package routers

import (
	"github.com/gin-gonic/gin"
	"hello-gin/controller/api"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userList", api.ApiController{}.UserList)
	}
}
