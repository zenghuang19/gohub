package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {

	// new 一个 Gin Engine 实例
	router := gin.New()

	bootstrap.SetupRoute(router)

	//运行服务
	err := router.Run(":3000")
	if err != nil {
		//错误处理，端口被占了或其他错误
		fmt.Println(err.Error())
	}
}
