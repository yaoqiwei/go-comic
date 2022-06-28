package middleware

import (
	"bytes"
	"fehu/conf"
	"fehu/constant"
	"fehu/util/cryp"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func HeaderAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !conf.Http.HeaderCheck {
			c.Next()
			return
		}

		timestamp := c.GetHeader("timestamp")

		if timestamp == "" {
			panic("block_4")
		}

		ti, _ := strconv.ParseInt(timestamp, 10, 64)
		now := time.Now().Unix()

		if ti-10 > now || ti+10 < now {
			panic("block_3")
		}

		if !conf.Http.AesOpen {
			c.Next()
			return
		}

		type stru struct {
			Endata string `json:"endata" form:"endata"`
		}

		var o stru
		c.ShouldBind(&o)

		if o.Endata == "" {
			panic("block_1")
		}

		str, _ := cryp.AesDecrypt(o.Endata, constant.ApiAesKey)
		if str == "" {
			panic("block_2")
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(str)))
		c.Next()
	}
}
