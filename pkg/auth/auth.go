package auth

import (
	"github.com/pkg/errors"
	"gohub/app/models/user"
)

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	userModel := user.GetByMulti(email)

	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码不正确")
	}

	return userModel, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)

	if userModel.ID == 0 {
		return user.User{}, errors.New("用户不存在")
	}

	return userModel, nil
}
