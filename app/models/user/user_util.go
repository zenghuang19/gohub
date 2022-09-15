package user

import "gohub/pkg/database"

// IsEmailExist 判断email是否被注册
func IsEmailExist(email string) bool {
	var count int64

	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

//IsPhoneExist 判断电话号码是否被注册
func IsPhoneExist(phone string) bool {
	var count int64

	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)

	return count > 0
}

// Get 通过Id获取数据
func Get(idstr string) (userModel User) {
	database.DB.Where("id", idstr).First(&userModel)
	return
}
