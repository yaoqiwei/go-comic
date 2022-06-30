package user

import (
	"fehu/common/lib/gorm"
)

// GetUserLoginInfo 根据用户查询用户信息
func GetUserLoginInfo(userLogin string) *Users {
	users := new(Users)
	gorm.Db.Where("user_login=? or mobile = ? or user_email=?",
		userLogin, userLogin, userLogin).First(users)
	return users
}
