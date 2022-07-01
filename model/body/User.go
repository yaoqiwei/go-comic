package body

// UserLoginParam 用户登录需要传入参数
type UserLoginParam struct {
	UserLogin string `json:"userLogin" form:"userLogin" binding:"required"` //用户名
	UserPass  string `json:"userPass" form:"userPass" binding:"required"`   //用户密码
}

// UserLoginReturn 用户登录返回参数
type UserLoginReturn struct {
	Token string `json:"token"`
	Id    int64  `json:"id"`
}

// UserRegisterParam 用户注册信息
type UserRegisterParam struct {
	UserLogin    string `json:"user_login" form:"user_login"`                  // 游客类型不用传此参数
	Password     string `json:"user_pass" form:"user_pass" binding:"required"` // 用户密码
	Source       string `json:"source" form:"source"`                          // 注册来源
	HardwareId   string `json:"hardwareId" form:"hardwareId"`                  // 硬件ID
	RegisterType byte   `json:"registerType" form:"registerType"`              // 1游客，默认0普通注册
}

// UserCodeParam 获取验证码传入参数
type UserCodeParam struct {
	UserLogin string `json:"userLogin" form:"userLogin" binding:"required"` // 用户名
	Type      string `json:"type" form:"type"`                              // 找回密码find,注册reg,绑定bind
}

// UserFindPassParam 用户找回密码传入参数
type UserFindPassParam struct {
	UserLogin string `json:"user_login" form:"user_login" binding:"required"` //用户名
	Password  string `json:"user_pass" form:"user_pass" binding:"required"`   //用户密码
	Code      string `json:"code" form:"code" binding:"required"`             //验证码
}

// EmailParam 绑定邮箱传入参数
type EmailParam struct {
	Email string `json:"email" form:"email"`                  //邮箱号
	Code  string `json:"code" form:"code" binding:"required"` //验证码
}

// MobileParam 绑定手机传入参数
type MobileParam struct {
	Mobile string `json:"mobile" form:"mobile"`                //手机号
	Code   string `json:"code" form:"code" binding:"required"` //验证码
}
