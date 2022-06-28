package user

import "fehu/common/lib"

type Users struct {
	Id         int64
	UserLogin  string
	UserPass   string
	PrivateKey string
	UserStatus byte
}

// GetUserLoginInfo 根据用户查询用户信息
func GetUserLoginInfo(userLogin string) *Users {
	users := new(Users)
	lib.Db.Where("user_login=? or mobile = ? or user_email=?",
		userLogin, userLogin, userLogin).First(users)
	return users
}
