package http_error

var NoGuardInfo = HttpError{
	ErrorCode: 15001,
	ErrorMsg:  "守护信息不存在！",
}
var TopGuard = HttpError{
	ErrorCode: 15002,
	ErrorMsg:  "已经是尊贵守护了，不能购买普通守护",
}
