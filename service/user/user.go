package user

import (
	"fehu/common/lib/gorm"
	"fehu/common/lib/redis_lib"
	"fehu/model/http_error"
	"fehu/service/base"
	"fehu/util/jwt"
	"fehu/util/math"
	"fehu/util/tool"
	"strconv"
	"time"
)

// Users 对应数据表users字段
type Users struct {
	Id           int64
	UserLogin    string //用户名
	UserPass     string //用户密码
	Mobile       string //用户手机号
	UserEmail    string //用户邮箱
	UserNiceName string //用户美明
	Signature    string //用户个性签名
	Avatar       string //用户头像
	AvatarThumb  string //用户头像缩略图
	UserType     byte   //用户类型，1:admin ;2:会员
	Source       string //注册来源
	PrivateKey   string //密码密钥
	HardwareId   string //硬件id
	CheckEmail   byte   //是否绑定和验证邮箱
	CheckMobile  byte   //是否绑定手机
	UserStatus   byte   //用户状态 0：禁用； 1：正常 ；2：未验证
}

// UsersInfo 对应数据表user_info字段
type UsersInfo struct {
	UserId        int64     //用户id
	LastLoginIp   string    //最后登录ip
	LastLoginTime time.Time //最后登录时间
	Score         int       //用户积分
	Coin          int64     //金币
	Token         string    //授权token
	ExpireTime    int64     //token到期时间
}

// GetLoginNameByHardwareId 根据硬件id查询对应用户名
func GetLoginNameByHardwareId(hardwareId string) string {
	user := new(Users)
	gorm.Db.Select("user_login").Where("hardware_id=?", hardwareId).First(user)
	return user.UserLogin
}

// AddUser 添加用户
func AddUser(userLogin, password, source, email, mobile, ip, hardwareId string) int64 {
	privateKey := math.GetRandomStr(6)
	password = jwt.SetPass(password, privateKey)
	niceName := tool.RandName()
	if CheckUserExist(userLogin) {
		panic(http_error.AccountExist)
	}
	randAvatar := "/userIcon" + strconv.FormatInt(int64(math.Rand(1, 20)), 10) + ".png"
	users := Users{
		UserLogin:    userLogin,
		Mobile:       mobile,
		UserEmail:    email,
		UserNiceName: niceName,
		UserPass:     password,
		Signature:    "这家伙很懒，什么都没留下",
		Avatar:       randAvatar,
		AvatarThumb:  randAvatar,
		UserType:     0,
		UserStatus:   1,
		Source:       source,
		PrivateKey:   privateKey,
		HardwareId:   hardwareId,
	}
	err := gorm.Db.Create(&users).Error
	uid := users.Id
	if err != nil {
		panic(http_error.RegisterFail)
	}
	usersInfo := UsersInfo{
		UserId:      uid,
		LastLoginIp: ip,
	}
	gorm.Db.Create(&usersInfo)
	return users.Id
}

// UsersCodeIpLimit 限制每个ip每天发送验证码次数
func UsersCodeIpLimit(ip string) {
	configPri := base.GetConfigPri()
	ipLimitSwitch, _ := strconv.ParseInt(configPri.IpLimitSwitch, 10, 64)
	ipLimitTimes, _ := strconv.ParseInt(configPri.IpLimitTimes, 10, 64)
	if ipLimitSwitch == 0 {
		return
	}
	date := time.Now().Format("20060102")
	key := redis_lib.GetRedisKey("CODE_IP_LIMIT", ip, date)
	times, _ := redis_lib.GetInt64(key, "")
	if times >= ipLimitTimes {
		panic(http_error.TooManyCodeCalled)
	}
	redis_lib.Set(key, times+1, 3600*24, "")
}

// UpdateUserPass 更新用户密码
func UpdateUserPass(uid int64, password, privateKey string) {
	password = jwt.SetPass(password, privateKey)
	users := Users{
		UserPass:   password,
		PrivateKey: privateKey,
	}
	gorm.Db.Where("id = ?", uid).Updates(users)
}

// GetUserInfo 获取用户信息
func GetUserInfo(userId int64) Users {
	user := Users{Id: userId}
	gorm.Db.First(&user)
	return user
}

// UpdateUser 根据id修改用户信息
func UpdateUser(user Users) {
	gorm.Db.Where("id = ?", user.Id).Updates(user)
}
