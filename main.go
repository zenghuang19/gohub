package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
)

func init()  {
	//加载配置
	btsConfig.Initialize()
}

func main() {
	//配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

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
