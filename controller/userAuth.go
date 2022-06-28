package controller

import (
	"encoding/base64"
	"fehu/middleware"
	"fehu/model/body"
	"fehu/model/http_error"
	"fehu/service/user"
	"fehu/util/jwt"
	"fehu/util/request"
	"github.com/gin-gonic/gin"
	"strconv"
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
	userInfo := user.GetUserLoginInfo(params.UserLogin)
	if userInfo == nil {
		panic(http_error.AccountNotExist)
	}
	pass := jwt.SetPass(params.UserPass, userInfo.UserPass)
	if userInfo.UserPass != pass {
		panic(http_error.AccountOrPasswordError)
	}
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
func (*UserController)register(c *gin.Context){

}