package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controlles/api/v1"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseApiController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	//生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志
	logger.LogIf(err)

	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}