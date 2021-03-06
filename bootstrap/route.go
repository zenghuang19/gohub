package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"
)

func SetupRoute(route *gin.Engine) {
	//注册全局中间件
	registerGlobMiddleWare(route)

	// 路由
	routes.RegisterAPIRoutes(route)

	//404
	setup404Handler(route)
}

func registerGlobMiddleWare(route *gin.Engine) {
	route.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(route *gin.Engine) {
	// 处理404
	route.NoRoute(func(c *gin.Context) {
		//获取标头信息的Accept信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			//是html
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			//默认返回JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认URL和请求方法是否正确。",
			})
		}
	})
}
