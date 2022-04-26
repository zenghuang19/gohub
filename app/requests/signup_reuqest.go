package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

// SignupPhoneExist 手机号码验证
func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	//定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	//错误提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号码必填",
			"digits:手机号码应为11位字符串数字",
		},
	}
	return validate(data, rules, messages)
}

type SignupEmailRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	//定义规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	//错误规则
	messages := govalidator.MapData{
		"email": []string{
			"required: 请填写邮箱",
			"min:邮箱长度需大于 4",
			"max:邮箱长度应小于 30",
			"email:邮箱格式不正确",
		},
	}

	return validate(data, rules, messages)
}
