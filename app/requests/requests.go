package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

//ValidateFun 验证函数
type ValidateFun func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidateFun) bool {
	//1.解析请求，支持JSON数据、表单请求和URL Query
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors":  err.Error(),
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
		})

		fmt.Println(err.Error())

		return false
	}

	//2.表单验证
	erro := handler(obj, c)

	//3.判断验证是否通过
	if len(erro) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error":   erro,
			"message": "请求验证不通过",
		})

		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, message govalidator.MapData) map[string][]string {
	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      message,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}
