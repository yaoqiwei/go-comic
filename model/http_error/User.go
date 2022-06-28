package http_error

var JwtError = HttpError{
	ErrorCode: 700,
	ErrorMsg:  "用户登录状态已失效，请重新登录！",
}
var CoinNotEnough = HttpError{
	ErrorCode: 701,
	ErrorMsg:  "余额不足，请充值！",
}

var AccountNotExist = HttpError{
	ErrorCode: 1002,
	ErrorMsg:  "账号不存在",
}

var AccountOrPasswordError = HttpError{
	ErrorCode: 1003,
	ErrorMsg:  "账号或密码错误",
}

var AccountBanned = HttpError{
	ErrorCode: 1004,
	ErrorMsg:  "该账号已被禁用",
}

var CodeNotSend = HttpError{
	ErrorCode: 1005,
	ErrorMsg:  "请先获取验证码",
}

var CodeError = HttpError{
	ErrorCode: 1006,
	ErrorMsg:  "验证码错误",
}

var PasswordTypeError = HttpError{
	ErrorCode: 1007,
	ErrorMsg:  "密码不能纯数字或纯字母",
}

var PasswordCountError = HttpError{
	ErrorCode: 1008,
	ErrorMsg:  "密码6-12位数字与字母",
}

var RegisterFail = HttpError{
	ErrorCode: 1009,
	ErrorMsg:  "注册失败，请重试！",
}

var AccountExist = HttpError{
	ErrorCode: 1010,
	ErrorMsg:  "该账号已被注册！",
}

var CodeNotExpdired = HttpError{
	ErrorCode: 1011,
	ErrorMsg:  "验证码5分钟有效，请勿多次发送",
}

var IsNotPhoneOrEmail = HttpError{
	ErrorCode: 1012,
	ErrorMsg:  "请输入正确的账号格式",
}

var TooManyCodeCalled = HttpError{
	ErrorCode: 1013,
	ErrorMsg:  "您已当日发送次数过多",
}

var NicknameIsNil = HttpError{
	ErrorCode: 1014,
	ErrorMsg:  "昵称不能为空",
}

var VipInfoError = HttpError{
	ErrorCode: 1015,
	ErrorMsg:  "VIP信息错误",
}

var IsNotPhone = HttpError{
	ErrorCode: 1016,
	ErrorMsg:  "请输入正确的手机号",
}

var IsAuthed = HttpError{
	ErrorCode: 1017,
	ErrorMsg:  "已经完成认证！",
}

var AuthNoFront = HttpError{
	ErrorCode: 1018,
	ErrorMsg:  "请上传正面照片",
}

var AuthNoBack = HttpError{
	ErrorCode: 1019,
	ErrorMsg:  "请上传正面照片",
}

var AuthNoHandSet = HttpError{
	ErrorCode: 1020,
	ErrorMsg:  "请上传手持照片",
}

var IsAuthing = HttpError{
	ErrorCode: 1017,
	ErrorMsg:  "正在审核认证！",
}

var IsNotAuthed = HttpError{
	ErrorCode: 1018,
	ErrorMsg:  "请先进行身份认证或等待审核",
}

var NoChargeRule = HttpError{
	ErrorCode: 1019,
	ErrorMsg:  "无充值选项",
}

var OldPasswordWrong = HttpError{
	ErrorCode: 1020,
	ErrorMsg:  "旧密码错误",
}

var IsBindEmail = HttpError{
	ErrorCode: 1021,
	ErrorMsg:  "已经绑定邮箱",
}

var IsBindMobile = HttpError{
	ErrorCode: 1022,
	ErrorMsg:  "已经绑定手机",
}

var FollowSelf = HttpError{
	ErrorCode: 1023,
	ErrorMsg:  "不能关注自己",
}

var ErrorSign = HttpError{
	ErrorCode: 1024,
	ErrorMsg:  "签名错误",
}

var NoOrder = HttpError{
	ErrorCode: 1025,
	ErrorMsg:  "订单不存在",
}

var NoChannel = HttpError{
	ErrorCode: 1026,
	ErrorMsg:  "无充值通道",
}

var WithdrawalsLimit = HttpError{
	ErrorCode: 1027,
	ErrorMsg:  "达到200元后方可提现~",
}

var NoBank = HttpError{
	ErrorCode: 1028,
	ErrorMsg:  "银行信息不存在",
}

var NoChatCoin = HttpError{
	ErrorCode: 1029,
	ErrorMsg:  "对方未设置聊天阈值无需支付",
}

var NoMatchChatCoin = HttpError{
	ErrorCode: 1030,
	ErrorMsg:  "对方修改了需支付的聊天阈值",
}

var HasChat = HttpError{
	ErrorCode: 1031,
	ErrorMsg:  "已和对方建立聊天，无需支付",
}

var HasChatPaid = HttpError{
	ErrorCode: 1032,
	ErrorMsg:  "已付款，无需支付",
}

var NoAuthChat = HttpError{
	ErrorCode: 1033,
	ErrorMsg:  "无权限聊天",
}

var ChatIsCustomerService = HttpError{
	ErrorCode: 1034,
	ErrorMsg:  "客服私聊无需支付",
}

var IsNotGuest = HttpError{
	ErrorCode: 1035,
	ErrorMsg:  "已经绑定账号",
}

var IsBindAccount = HttpError{
	ErrorCode: 1036,
	ErrorMsg:  "该账号已经绑定其他账号",
}

var WithdrawalsErr = HttpError{
	ErrorCode: 1037,
	ErrorMsg:  "提现失败",
}

var NotCustomerService = HttpError{
	ErrorCode: 1038,
	ErrorMsg:  "不是客服",
}

var EmailCodeNotSupport = HttpError{
	ErrorCode: 1039,
	ErrorMsg:  "暂不支持邮箱发送验证码",
}
