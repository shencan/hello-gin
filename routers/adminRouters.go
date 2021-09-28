package routers

import (
	"github.com/gin-gonic/gin"
	"hello-gin/controller/admin"
	"net/http"
)

type User struct {
	Name  string `json:"name" form:"name"`
	Title string `json:"title" form:"title"`
	Desc  string `json:"desc" form:"desc"`
}

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/test", func(c *gin.Context) {

			id := c.Query("id")
			title := c.Query("title")
			page := c.DefaultQuery("page", "1")

			a := &User{
				Title: "后台首页",
				Desc:  "描述",
			}
			c.HTML(http.StatusOK, "admin/index.html", gin.H{
				"title": title,
				"page":  page,
				"id":    id,
				"users": a,
			})
		})

		//不加括号，加括号为引用，不加括号为注册
		adminRouters.GET("/", admin.IndexController{}.Index)
		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/user/delete", admin.UserController{}.Delete)
	}
}
