package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/requests/validators"
)

type ResetByEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func ResetByEmail(data interface{}, c *gin.Context) map[string][]string {
	// 规则
	rules := govalidator.MapData{
		"password":    []string{"required", "min:6"},
		"verify_code": []string{"required"},
		"email":       []string{"required", "min:4", "max:30", "email"},
	}

	// 规则字段对应的错误消息
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			//"digits:验证码长度必须为 6 位的数字",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	errs := validate(data, rules, messages)

	//检查验证
	_data := data.(*ResetByEmailRequest)
	errs = validators.ValidateCaptcha(_data.Email, _data.VerifyCode, errs)

	return errs
}
