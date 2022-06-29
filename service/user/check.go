package user

import (
	"fehu/common/lib/mysql_lib"
)

// CheckUserExist 根据用户名判断用户是否存在
func CheckUserExist(userLogin string) bool {
	return mysql_lib.DB("users").
		Query("WHERE user_login=? OR mobile=? OR user_email=?",
			userLogin, userLogin, userLogin).Exist()
}
