package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello-gin/models"
	"net/http"
)

type UserController struct {
	baseController
}

func (con UserController) Index(c *gin.Context) {
	//c.String(200, "用户列表")
	//查询数据库

	userList := []models.Member{}

	models.DB.Where(" id < 10 ").Find(&userList)

	c.JSON(http.StatusOK, gin.H{
		"res": userList,
	})
	//con.success(c)
}

func (con UserController) Add(c *gin.Context) {
	user := models.Member{
		Username: "test",
		Nickname: "test",
		License:  "",
	}

	models.DB.Create(&user)

	fmt.Println(user)

	c.String(200, "添加用户成功")
}

func (con UserController) Edit(c *gin.Context) {
	c.String(200, "修改用户")
}

func (con UserController) Delete(c *gin.Context) {
	c.String(200, "删除用户")
}
