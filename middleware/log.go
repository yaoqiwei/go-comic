package middleware

import (
	"bytes"
	"fehu/util/jwt"
	"fehu/util/stringify"
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path
		start := time.Now()
		method := c.Request.Method
		var body string

		contentType := c.Request.Header.Get("Content-Type")
		for _, v := range stringify.ToStringSlice(contentType, ";") {
			if v == "application/json" {
				bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

				body = string(bodyBytes)
				if body != "{}" && body != "" {
					body = ", body:" + body
				} else {
					body = ""
				}
			}
		}

		// logrus.Infof("[api]: %15s | %-7s %#v", clientIP, method, path)

		defer func() {
			var user string
			uid := jwt.GetUid(c)
			if uid > 0 {
				user = c.ClientIP() + ", " + stringify.ToString(uid)
			} else {
				user = c.ClientIP()
			}

			logrus.Infof("%s %s:%s%s", user, method, path, body)

			latency := time.Since(start)
			if latency > time.Second {
				logrus.Warnf("%s %s:%s, latency:%v", user, method, path, latency)
			}
		}()

		c.Next()

	}
}
