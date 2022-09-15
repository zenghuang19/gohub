package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controlles/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"
)

type PasswordController struct {
	v1.BaseApiController
}

func (p PasswordController) ResetPassword(c *gin.Context) {
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}
	// 更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		user.Save(&userModel)

		response.Success(c)
	}
}

func (p PasswordController) ResetByEmail(c *gin.Context) {
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}

	// 更新密码
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		user.Save(&userModel)

		response.Success(c)
	}
}
