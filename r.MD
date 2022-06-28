### 文档说明2.0

BaseUrl: https://api.xxxx.com

###### 请求接口方式
BaseUrl + Path
例：POST https://api.xxx.com/home/getBanner

###### 请求类型均为POST
请求的Content-Type支持application/json,application/x-www-form-urlencoded,multipart/form-data,建议使用application/json

###### 返回状态永远是
Status Code: 200; Content-Type: application/json;

###### 返回格式
```
{
  "code": number,   /** code==0表示成功,其他都是失败 */
  "data": object,   /** 内容在这里显示，永远是一个对象 */
  "msg" : string    /** code!=0时候，失败信息在这里显示 */
}
```

###### 常见的Error
```
HttpError{
	ErrorCode: 5001,
	ErrorMsg:  "缺少参数",
}

HttpError{
	ErrorCode: 5001,
	ErrorMsg:  "Error creating the  es client",
}

HttpError{
	ErrorCode: 5002,
	ErrorMsg:  "wrong page",
}

HttpError{
	ErrorCode: 5003,
	ErrorMsg:  "uploaded file is too large",
}

HttpError{
	ErrorCode: 700,
	ErrorMsg:  "用户登录状态已失效，请重新登录！",
}

HttpError{
	ErrorCode: 701,
	ErrorMsg:  "余额不足，请充值！",
}

```
