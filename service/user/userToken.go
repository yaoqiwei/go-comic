package user

import (
	"fehu/common/lib/mysql_lib"
	"fehu/common/lib/redis_lib"
	"fehu/util/jwt"
	"time"
)

// UserTokenRedis 用户token结构体
type UserTokenRedis struct {
	Uid        int64  `db:"uid" json:"uid"`               //用户id
	Token      string `db:"token" json:"token"`           //用户token
	Expiretime int64  `db:"expiretime" json:"expiretime"` //失效时间
}

// UpdateToken 更新token
func UpdateToken(token *jwt.Token, ip string) {

	expiretime := token.Expire.Unix() + 60*60*24*300
	tl := token.Expire.Format("2006-01-02 15:04:05")

	mysql_lib.DB("users_info").Suffix(token.Uid).Query("SET token=?,expiretime=?,last_login_time=?,last_login_ip=? WHERE uid=?", token.Token, expiretime, tl, ip, token.Uid).Update()

	tokenRedisInfo := UserTokenRedis{
		Uid:        token.Uid,
		Token:      token.Token,
		Expiretime: expiretime,
	}

	tokenKey := redis_lib.GetRedisKey("TOKEN", token.Uid)
	redis_lib.Set(tokenKey, tokenRedisInfo, 60*60*24*300, "")
}

// GetTokenMysql 从数据库中获取token
func GetTokenMysql(uid int64) *UserTokenRedis {
	userTokenRedis := &UserTokenRedis{}
	err := mysql_lib.DB("users_info").Suffix(uid).Query("WHERE uid=?", uid).Dest(userTokenRedis).FetchOne()
	if err != nil {
		return nil
	}
	if userTokenRedis.Uid == 0 {
		return nil
	}
	if userTokenRedis.Expiretime < time.Now().Unix() {
		return nil
	}
	return userTokenRedis
}
