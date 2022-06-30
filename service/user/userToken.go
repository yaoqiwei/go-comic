package user

import (
	"fehu/common/lib/gorm"
	"fehu/common/lib/redis_lib"
	"fehu/util/jwt"
	"time"
)

// UserTokenRedis 用户token结构体
type UserTokenRedis struct {
	UserId     int64  `json:"userId"`      //用户id
	Token      string `json:"token"`       //用户token
	ExpireTime int64  ` json:"expireTime"` //失效时间
}

// UpdateToken 更新token
func UpdateToken(token *jwt.Token, ip string) {
	expireTime := token.Expire.Unix() + 60*60*24*300
	tl := token.Expire
	usersInfo := UsersInfo{
		Token:         token.Token,
		ExpireTime:    expireTime,
		LastLoginTime: tl,
		LastLoginIp:   ip,
	}
	gorm.Db.Where("user_id = ?", token.Uid).Updates(&usersInfo)
	tokenRedisInfo := UserTokenRedis{
		UserId:     token.Uid,
		Token:      token.Token,
		ExpireTime: expireTime,
	}
	tokenKey := redis_lib.GetRedisKey("TOKEN", token.Uid)
	redis_lib.Set(tokenKey, tokenRedisInfo, 60*60*24*300, "")
}

// GetTokenMysql 从数据库中获取token
func GetTokenMysql(uid int64) *UsersInfo {
	userInfo := &UsersInfo{}
	err := gorm.Db.Where("user_id=?", uid).First(userInfo).Error
	if err != nil {
		return nil
	}
	if userInfo.UserId == 0 {
		return nil
	}
	if userInfo.ExpireTime < time.Now().Unix() {
		return nil
	}
	return userInfo
}
