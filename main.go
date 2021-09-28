package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"hello-gin/models"
	"hello-gin/routers"
	"html/template"
	"net/http"
	"time"
)

type User struct {
	Name  string `json:"name" form:"name"`
	Title string `json:"title" form:"title"`
	Desc  string `json:"desc" form:"desc"`
}

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("中间件开始1")
	c.Next()
	fmt.Println("中间件开始2")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func main() {

	//创建一个默认路由引擎
	//Default是默认有中间件的，不希望使用就gin.New()
	r := gin.Default()

	//自定义函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})

	//配置模板文件
	r.LoadHTMLGlob("templates/**/*")

	//静态web目录 第一个路由 第二个映射目录
	r.Static("/static", "./static")

	//配置session
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("miyao"))
	r.Use(sessions.Sessions("mysession", store))

	//全局中间件
	r.Use(initMiddleware)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong1111",
		})
	})

	//配置路由
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "值：%v", "你好vin")
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"meg":     "成功",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"su": true,
		})
	})

	r.GET("/json3", func(c *gin.Context) {
		a := &User{
			Title: "标题json3",
			Desc:  "描述",
		}
		c.JSON(http.StatusOK, a)
	})

	//解决跨域问题
	r.GET("/jsonp", func(c *gin.Context) {
		a := &User{
			Title: "标题jsonp",
			Desc:  "描述",
		}
		c.JSONP(http.StatusOK, a)
	})

	//html
	r.GET("/html", func(c *gin.Context) {
		a := &User{
			Title: "标题test",
			Desc:  "描述",
		}
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"users": a,
			"num":   80,
			"hobby": []string{
				"吃饭",
				"睡觉",
				"打豆豆",
			},
			"userList": []interface{}{
				a,
				a,
			},
			"null": []interface{}{

			},
			"news": &User{
				Title: "new title",
				Desc:  "new desc",
			},
			"date": 1632646391,
		})
	})

	//post
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/user.html", gin.H{})
	})

	r.POST("/doAdd", func(c *gin.Context) {
		user := &User{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err,
			})
		}
	})

	r.GET("/getUser", func(c *gin.Context) {
		user := &User{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err,
			})
		}
	})

	r.GET("/", initMiddleware, func(c *gin.Context) {
		fmt.Println("首页")
		c.String(http.StatusOK, "你好首页")
	})

	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)

	r.Run(":9999") // 监听并在 0.0.0.0:8080 上启动服务

}
