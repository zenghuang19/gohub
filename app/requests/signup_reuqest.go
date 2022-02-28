package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	//定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	//错误提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号码必填",
			"digits:手机号码应为11位的数字",
		},
	}

	//配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	//开始验证
	return govalidator.New(opts).ValidateStruct()
}
