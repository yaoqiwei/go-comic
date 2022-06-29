package controller

import (
	"encoding/base64"
	"fehu/common/lib/redis_lib"
	"fehu/middleware"
	"fehu/model/body"
	"fehu/model/http_error"
	"fehu/service/base"
	"fehu/service/user"
	"fehu/util/jwt"
	"fehu/util/map_builder"
	"fehu/util/math"
	"fehu/util/request"
	"fehu/util/tool"
	"fehu/util/validator"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type UserAuthController struct {
}

func UserAuthRegister(router *gin.RouterGroup, needLoginedRouter *gin.RouterGroup) {
	c := UserAuthController{}
	router.POST("/login/userLogin", c.login)
}

// @Summary 登录
// @Description 登录
// @Tags 登录相关
// @Accept json
// @Param param body body.UserLoginParam true "账号和密码"
// @Success 0 {object} body.UserLoginReturnFull
// @Failure 1002 {object} http_error.HttpError
// @Failure 1003 {object} http_error.HttpError
// @Failure 1004 {object} http_error.HttpError
// @Router /login/userLogin [post]
func (*UserAuthController) login(c *gin.Context) {
	var params body.UserLoginParam
	request.Bind(c, &params)
	// 给密码加密，获取用户信息
	userInfo := user.GetUserLoginInfo(params.UserLogin)
	// 用户不存在
	if userInfo == nil {
		panic(http_error.AccountNotExist)
	}
	pass := jwt.SetPass(params.UserPass, userInfo.UserPass)
	// 密码不匹配
	if userInfo.UserPass != pass {
		panic(http_error.AccountOrPasswordError)
	}
	// 用户封禁
	if userInfo.UserStatus == 0 {
		panic(http_error.AccountBanned)
	}
	var encodeToken string
	token := user.GetTokenMysql(userInfo.Id)
	if token != nil {
		encodeToken = base64.StdEncoding.EncodeToString([]byte(token.Token + "|" + strconv.FormatInt(token.Uid, 10)))
	}
	if encodeToken == "" {
		// 生成token并更新
		token := jwt.SetToken(userInfo.Id, params.UserLogin)
		user.UpdateToken(token, c.ClientIP())
		encodeToken = token.EncodeToken
	}
	middleware.Success(c, body.UserLoginReturn{Token: encodeToken, Id: userInfo.Id})
}

// @Summary 注册
// @Description 注册
// @Tags 登录相关
// @Accept json
// @Param registerParam body body.UserRegisterParam true "账号和密码"
// @Success 0 {object} body.UserRegisterReturnFull
// @Failure 1005 {object} http_error.HttpError
// @Failure 1006 {object} http_error.HttpError
// @Failure 1007 {object} http_error.HttpError
// @Failure 1008 {object} http_error.HttpError
// @Failure 1010 {object} http_error.HttpError
// @Failure 1012 {object} http_error.HttpError
// @Router /login/userReg [post]
func (*UserController) register(c *gin.Context) {
	var param body.UserRegisterParam
	request.Bind(c, &param)
	// 检查密码格式
	validator.Passcheck(param.Password)
	userLogin := param.UserLogin
	var email string
	var mobile string
	if param.RegisterType == 1 && param.HardwareId != "" {
		userLogin = user.GetLoginNameByHardwareId(param.HardwareId)
		if userLogin != "" {
			middleware.Success(c, map_builder.BuilderMap("userName", userLogin))
			return
		}
		userLogin = "user_" + time.Now().Format("20060102150405") + math.GetRandomInt(6)
	} else if validator.IsMail(param.UserLogin) {
		email = userLogin
	} else if validator.IsPhone(param.UserLogin) {
		mobile = userLogin
	} else {
		panic(http_error.IsNotPhoneOrEmail)
	}
	uid := user.AddUser(userLogin, param.Password, param.Source, email, mobile, c.ClientIP(), param.HardwareId)
	middleware.Success(c, map_builder.BuilderMap("userName", userLogin, "uid", uid))
}

// @Summary 获取验证码
// @Description 获取验证码
// @Tags 登录相关
// @Accept json
// @Param codeParam body body.UserCodeParam true "账号和类型，type:find找回密码，reg注册"
// @Success 0 {object} body.UserCodeReturnFull
// @Failure 1010 {object} http_error.HttpError
// @Failure 1011 {object} http_error.HttpError
// @Failure 1012 {object} http_error.HttpError
// @Failure 1013 {object} http_error.HttpError
// @Router /login/getCode [post]
func (*UserAuthController) getCode(c *gin.Context) {
	var param body.UserCodeParam
	request.Bind(c, &param)
	// 获取注册验证码缓存
	codeRedisKey := redis_lib.GetRedisKey("REG_CODE", param.UserLogin)
	if param.Type == "find" {
		codeRedisKey = redis_lib.GetRedisKey("FIND_CODE", param.UserLogin)
	}
	if param.Type == "bind" {
		codeRedisKey = redis_lib.GetRedisKey("BIND_CODE", param.UserLogin)
	}
	regCode, _ := redis_lib.GetString(codeRedisKey, "")
	if regCode != "" {
		panic(http_error.CodeNotExpdired)
	}
	isMail := validator.IsMail(param.UserLogin)
	isPhone := validator.IsPhone(param.UserLogin)
	if !isMail && !isPhone {
		panic(http_error.IsNotPhoneOrEmail)
	}
	user.UsersCodeIpLimit(c.ClientIP())
	regCode = math.GetRandomInt(6)
	redis_lib.Set(codeRedisKey, regCode, 300, "")
	configPri := base.GetConfigPri()
	sendCodeWitch, _ := strconv.ParseInt(configPri.SendcodeWwitch, 10, 64)
	if sendCodeWitch == 0 {
		middleware.Success(c, map_builder.BuilderMap("code", regCode))
		return
	}
	if isMail {
		panic(http_error.EmailNetworkErr)
		// tool.EmailSend(userCodeParam.UserLogin, regCode)
	}
	if isPhone {
		tool.TxSmsSend(param.UserLogin, regCode)
	}
	middleware.Success(c)
}

// @Summary 找回密码
// @Description 找回密码
// @Tags 登录相关
// @Accept json
// @Param param body body.UserFindPassParam true "账号和密码和验证码"
// @Success 0 {object} body.UserLoginReturnFull
// @Failure 1002 {object} http_error.HttpError
// @Failure 1003 {object} http_error.HttpError
// @Failure 1004 {object} http_error.HttpError
// @Router /login/userFindPass [post]
func (*UserAuthController) userFindPass(c *gin.Context) {
	var param body.UserFindPassParam
	request.Bind(c, &param)
	// 获取注册验证码缓存
	codeRedisKey := redis_lib.GetRedisKey("FIND_CODE", param.UserLogin)
	regCode, _ := redis_lib.GetString(codeRedisKey, "")
	// 没有发送验证码
	if regCode == "" {
		panic(http_error.CodeNotSend)
	}
	// 验证码不匹配
	if regCode != param.Code {
		panic(http_error.CodeError)
	}
	// 检查密码格式
	validator.Passcheck(param.Password)
	// 给密码加密，获取用户信息
	userInfo := user.GetUserLoginInfo(param.UserLogin)
	// 用户不存在
	if userInfo == nil {
		panic(http_error.AccountNotExist)
	}
	// 用户封禁
	if userInfo.UserStatus == 0 {
		panic(http_error.AccountBanned)
	}
	privateKey := math.GetRandomStr(6)
	user.UpdateUserPass(userInfo.Id, param.Password, privateKey)
	// 生成token并更新
	token := jwt.SetToken(userInfo.Id, userInfo.UserLogin)
	user.UpdateToken(token, c.ClientIP())
	middleware.Success(c, body.UserLoginReturn{Token: token.EncodeToken, Id: userInfo.Id})
}

// @Summary 绑定邮箱
// @Description 绑定邮箱
// @Tags 用户相关
// @Accept json
// @Param Auth header string true "token"
// @Param param body body.EmailParam true "-"
// @Success 0 {object} success.InfoData
// @Router /user/bindEmail [post]
func (*UserAuthController) bindEmail(c *gin.Context) {
	var params body.EmailParam
	request.Bind(c, &params)
	email := params.Email
	uid := jwt.GetUid(c, true)
	userInfo := user.GetUserInfo(uid)
	// 检查是否已经绑定邮箱
	if userInfo.CheckEmail != 0 {
		panic(http_error.IsBindEmail)
	}
	if userInfo.UserEmail != "" {
		email = userInfo.UserEmail
	}
	// 获取注册验证码缓存
	codeRedisKey := redis_lib.GetRedisKey("BIND_CODE", email)
	regCode, _ := redis_lib.GetString(codeRedisKey, "")
	// 没有发送验证码
	if regCode == "" {
		panic(http_error.CodeNotSend)
	}
	// 验证码不匹配
	if regCode != params.Code {
		panic(http_error.CodeError)
	}
	// 更新邮箱
	user.UpdateUser(user.Users{
		Id:         uid,
		UserEmail:  email,
		CheckEmail: 1,
	})
	middleware.Success(c)
}

// @Summary 绑定手机
// @Description 绑定手机
// @Tags 用户相关
// @Accept json
// @Param Auth header string true "token"
// @Param param body body.MobileParam true "-"
// @Success 0 {object} success.InfoData
// @Router /user/bindMobile [post]
func (*UserAuthController) bindMobile(c *gin.Context) {
	var params body.MobileParam
	request.Bind(c, &params)
	mobile := params.Mobile
	uid := jwt.GetUid(c, true)
	userInfo := user.GetUserInfo(uid)
	// 检查是否已经绑定手机
	if userInfo.CheckMobile != 0 {
		panic(http_error.IsBindMobile)
	}
	if userInfo.Mobile != "" {
		mobile = userInfo.Mobile
	}
	// 获取注册验证码缓存
	codeRedisKey := redis_lib.GetRedisKey("BIND_CODE", mobile)
	regCode, _ := redis_lib.GetString(codeRedisKey, "")
	// 没有发送验证码
	if regCode == "" {
		panic(http_error.CodeNotSend)
	}
	// 验证码不匹配
	if regCode != params.Code {
		panic(http_error.CodeError)
	}
	// 更新手机
	user.UpdateUser(user.Users{
		Id:          uid,
		Mobile:      mobile,
		CheckMobile: 1,
	})
	middleware.Success(c)
}
