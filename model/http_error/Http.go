package http_error

type ResData map[string]interface{}

type HttpError struct {
	ErrorCode int    `json:"code" example:"400"`
	ErrorMsg  string `json:"msg" example:"status bad request"`
}

var MissingParametersError = HttpError{
	ErrorCode: 5001,
	ErrorMsg:  "缺少参数",
}

var EsError = HttpError{
	ErrorCode: 5005,
	ErrorMsg:  "Error creating the  es client",
}

var WrongPage = HttpError{
	ErrorCode: 5002,
	ErrorMsg:  "wrong page",
}

var UploadFileIsTooLarge = HttpError{
	ErrorCode: 5003,
	ErrorMsg:  "uploaded file is too large",
}

var UploadFileErr = HttpError{
	ErrorCode: 5004,
	ErrorMsg:  "uploaded file error",
}

var Maintain = HttpError{
	ErrorCode: 5006,
	ErrorMsg:  "维护中！",
}

var NoType = HttpError{
	ErrorCode: 600,
	ErrorMsg:  "不支持的类型",
}

var SignErr = HttpError{
	ErrorCode: 601,
	ErrorMsg:  "签名错误",
}

var MsmNetworkErr = HttpError{
	ErrorCode: 602,
	ErrorMsg:  "网络错误",
}

var EmailNetworkErr = HttpError{
	ErrorCode: 603,
	ErrorMsg:  "网络错误",
}

var NoRedisKey = HttpError{
	ErrorCode: 604,
	ErrorMsg:  "没有redis key",
}

var FrequentOperations = HttpError{
	ErrorCode: 605,
	ErrorMsg:  "操作太频繁，请稍后再试",
}

var NoAESKey = HttpError{
	ErrorCode: 606,
	ErrorMsg:  "没有aes key",
}
