package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controlles/api/v1/auth"
)

func RegisterAPIRoutes(r *gin.Engine) {
	//路由组
	//版本v1
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			//判断手机是否已被注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)

			//email 验证码
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
		}

	}
}
