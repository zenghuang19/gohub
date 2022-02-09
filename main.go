package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	// new 一个 Gin Engine 实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	// 处理404
	r.NoRoute(func(c *gin.Context) {
		//获取标头信息的Accept信息
		accpetString := c.Request.Header.Get("Accept")
		if strings.Contains(accpetString, "text/html") {
			//是html
			c.String(http.StatusNotFound, "页面返回404")
		}else {
			//默认返回JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_message": "路由未定义，请确认URL和请求方法是否正确。",
			})
		}
	})

	// 运行服务，默认为 8080，我们指定端口为 8000
	r.Run(":8000")
}