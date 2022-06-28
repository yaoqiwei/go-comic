package body

// UserLoginParam 用户登录需要传入参数
type UserLoginParam struct {
	UserLogin string `json:"userLogin" form:"userLogin" binding:"required"` //用户名
	UserPass  string `json:"userPass" form:"userPass" binding:"required"`   //用户密码
	IsAgent   byte   `json:"isAgent" form:"isAgent"`
}

type UserLoginReturn struct {
	Token string `json:"token"`
	Id    int64  `json:"id"`
}