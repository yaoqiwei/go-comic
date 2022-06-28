package middleware

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		type stru struct {
			Uid  string `json:"uid" form:"uid"`
			Auth string `json:"token" form:"token"`
		}

		var o stru
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.ShouldBind(&o)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		auth := c.GetHeader("Auth")
		if auth == "" {
			auth = o.Auth
		}

		if auth == "" {
			c.Next()
			return
		}

		var uid int64

		authTrue, err := base64.StdEncoding.DecodeString(auth)
		if err == nil {
			arr := strings.Split(string(authTrue), "|")
			if len(arr) == 2 {
				auth = arr[0]
				uid, _ = strconv.ParseInt(arr[1], 10, 64)
			}
		}

		if uid == 0 {
			uidStr := c.GetHeader("Auth-Uid")
			if uidStr == "" {
				uidStr = o.Uid
			}
			uid, _ = strconv.ParseInt(uidStr, 10, 64)
		}

		// if user.CheckToken(uid, auth) {
		// 	redis_lib.ZAdd(redis_lib.GetRedisKey("LAST_USER_OPTION_TIME"), time.Now().Unix(), uid, "")
		// 	c.Set("uid", uid)
		// 	c.Set("token", auth)
		// }

		c.Next()
	}
}
