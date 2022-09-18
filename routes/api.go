package routes

import (
	"github.com/gin-gonic/gin"
	controller "gohub/app/http/controlles/api/v1"
	"gohub/app/http/controlles/api/v1/auth"
	"gohub/app/http/middlewares"
)

func RegisterAPIRoutes(r *gin.Engine) {
	//路由组
	//版本v1
	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")

		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 注册
			suc := new(auth.SignupController)
			//判断手机是否已被注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-h"), vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-h"), vcc.SendUsingPhone)
			//email 验证码
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-h"), vcc.SendUsingEmail)

			// 登录
			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			//刷新token
			authGroup.POST("/login/refresh-token", middlewares.AuthJwt(), lgc.RefreshToken)

			// 重置密码
			pws := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.AuthJwt(), pws.ResetPassword)
			authGroup.POST("/password-reset/using-email", middlewares.AuthJwt(), pws.ResetByEmail)
		}
	}

	user := new(controller.UsersController)
	// 获取当前用户
	v1.GET("/user", middlewares.AuthJwt(), user.CurrentUser)

	userGroup := v1.Group("/users")
	{
		userGroup.GET("", user.Index)
	}
}
