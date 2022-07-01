package user

import (
	"fehu/common/lib/gorm"
	"fehu/common/lib/redis_lib"
	"time"
)

// CheckUserExist 根据用户名判断用户是否存在
func CheckUserExist(userLogin string) bool {
	users := Users{}
	return gorm.Db.Where("user_login=? OR mobile=? OR user_email=?",
		userLogin, userLogin, userLogin).Find(&users).RowsAffected > 0
}

// CheckToken 校验token
func CheckToken(uid int64, token string) bool {
	if uid == 0 && token == "" {
		return false
	}
	key := redis_lib.GetRedisKey("TOKEN", uid)
	var userTokenRedis UserTokenRedis
	redis_lib.GetObject(key, &userTokenRedis, "")
	if userTokenRedis.UserId == 0 {
		//mysql_lib.DB("users_info").Suffix(uid).Query("WHERE uid=?", uid).Dest(&userTokenRedis).FetchOne()
		gorm.Db.Model(UsersInfo{}).Where("user_id = ?", uid).First(&userTokenRedis)
		if userTokenRedis.UserId != 0 {
			redis_lib.Set(key, userTokenRedis, 60*60*24*300, "")
		}
	}
	if userTokenRedis.UserId != 0 {
		if userTokenRedis.Token != token || userTokenRedis.ExpireTime < time.Now().Unix() {
			return false
		}
		return true
	}
	return false
}
