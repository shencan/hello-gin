package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println(time.Now())

	fmt.Println(c.Request.URL)

	c.Set("username", "张三")

	//定义一个goroutine统计日志
	cCp := c.Copy()

	go func() {
		time.Sleep(5*time.Second)
		fmt.Println("Done in path " + cCp.Request.URL.Path)
	}()

}