package param

import (
	"fehu/model/success"
)

type LoginParam struct {
	UserLogin string `json:"user_login" binding:"required"`
	Password  string `json:"user_pass" binding:"required"`
}

type LoginParamReturn struct {
	Token string `json:"token"`
}

type LoginParamReturnFull struct {
	success.DataBase
	Data LoginParamReturn
}
