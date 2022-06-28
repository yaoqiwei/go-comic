package http_error

var NewsNotExist = HttpError{
	ErrorCode: 3001,
	ErrorMsg:  "动态不存在！",
}

var EmptyNews = HttpError{
	ErrorCode: 3002,
	ErrorMsg:  "无法发送空动态！",
}

var EmptyReply = HttpError{
	ErrorCode: 3003,
	ErrorMsg:  "无法发送空回复！",
}

var NewsNoHide = HttpError{
	ErrorCode: 3004,
	ErrorMsg:  "无需购买隐藏内容！",
}

var IsNewsVip = HttpError{
	ErrorCode: 3005,
	ErrorMsg:  "已经是VIP，无需开通",
}

var NewsTypeNotExist = HttpError{
	ErrorCode: 3006,
	ErrorMsg:  "动态类型不存在！",
}

var UserNotNewsVip = HttpError{
	ErrorCode: 3007,
	ErrorMsg:  "用户不是线下VIP",
}

var DenounceNoContent = HttpError{
	ErrorCode: 3008,
	ErrorMsg:  "请填写举报内容",
}

var DenounceNoNews = HttpError{
	ErrorCode: 3009,
	ErrorMsg:  "请填写举报动态",
}
var NewsIsLiked = HttpError{
	ErrorCode: 3010,
	ErrorMsg:  "该评论已点赞",
}

var NewsIsQueryErr = HttpError{
	ErrorCode: 3011,
	ErrorMsg:  "请求参数不足",
}
