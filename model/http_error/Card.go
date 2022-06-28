package http_error

var CardGoodsNotExistError = HttpError{
	ErrorCode: 8001,
	ErrorMsg:  "商品不存在！",
}

var CardNotEnough = HttpError{
	ErrorCode: 8002,
	ErrorMsg:  "卡密余额不足，请充值！",
}

var CardNotExistError = HttpError{
	ErrorCode: 8003,
	ErrorMsg:  "卡密不存在！",
}

var CardUsedError = HttpError{
	ErrorCode: 8004,
	ErrorMsg:  "卡密已使用！",
}
