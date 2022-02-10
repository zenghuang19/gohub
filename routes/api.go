package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine)  {
	//路由组
	//版本v1
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			//JSON格式响应
			context.JSON(http.StatusOK,gin.H{
				"hello":"world",
			})
		})
	}
}
