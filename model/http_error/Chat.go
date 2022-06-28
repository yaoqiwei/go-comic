package http_error

var IsSelfError = HttpError{
	ErrorCode: 17000,
	ErrorMsg:  "不能给自己发消息",
}

var ChatThumbGetErr = HttpError{
	ErrorCode: 17001,
	ErrorMsg:  "未找到图片信息",
}
