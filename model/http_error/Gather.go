package http_error

var GatherNotExistError = HttpError{
	ErrorCode: 9001,
	ErrorMsg:  "该合集不存在！",
}

var GatherPriceChanged = HttpError{
	ErrorCode: 9002,
	ErrorMsg:  "该合集价格有变化，请刷新后重试！",
}

var GatherVipBuy = HttpError{
	ErrorCode: 9003,
	ErrorMsg:  "该付费合集需要VIP才可购买！",
}
