package middleware

import (
	"fehu/util/convert"
	"fehu/util/math"

	"github.com/gin-gonic/gin"
)

type ResponseCode int

//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode ResponseCode = iota
	ErrorCode
	EncryptionErrorCode ResponseCode = -99
)

type Response struct {
	ErrorCode ResponseCode `json:"code"`
	ErrorMsg  string       `json:"msg"`
	Data      interface{}  `json:"data"`
	TraceId   interface{}  `json:"traceId"`
	Stack     interface{}  `json:"stack,omitempty"`
}

func Error(c *gin.Context, code ResponseCode, msg string) {
	resp := &Response{ErrorCode: code, ErrorMsg: msg}
	SerializeJSON(c, resp, 200)
	c.Abort()
}

func Success(c *gin.Context, datas ...interface{}) {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}

	resp := &Response{ErrorCode: SuccessCode, ErrorMsg: "", Data: data}
	SerializeJSON(c, resp, 200)
}

func SerializeJSON(c *gin.Context, res *Response, httpCode int) {
	res.TraceId = math.GetRandomStri(6)
	if res.Data == nil {
		res.Data = struct{}{}
	}
	c.Set("response", convert.ToJson(res))
	c.JSON(httpCode, res)
}
