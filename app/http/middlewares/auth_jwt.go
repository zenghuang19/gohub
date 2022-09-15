package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/pkg/config"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取到token,解析令牌
		token, err := jwt.NewJWT().ParserToken(c)
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}
		// 解析成功，获取用户信息
		userModel := user.Get(token.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "未找到该用户信息")
			return
		}

		// 将用户信息存入content里面，以供后续使用
		c.Set("user_id", userModel.GetStringID())
		c.Set("user_name", userModel.Name)
		c.Set("user", userModel)
		c.Next()
	}
}
