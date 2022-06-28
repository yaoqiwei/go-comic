package jwt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fehu/constant"
	"fehu/model/http_error"
	"fehu/util/cryp"
	"io/ioutil"
	"sort"
	"strconv"
	"time"

	"fehu/util/stringify"

	"github.com/gin-gonic/gin"
)

type Token struct {
	Token       string
	EncodeToken string
	Uid         int64
	UserLogin   string
	Expire      time.Time
}

func SetPass(pass string, authcode string) string {
	return cryp.MD5(authcode + pass)
}

func SetToken(uid int64, userLogin string) *Token {

	t := time.Now()
	token := cryp.MD5(cryp.MD5(strconv.FormatInt(uid, 10) + userLogin + strconv.FormatInt(t.Unix(), 10)))

	encodeToken := base64.StdEncoding.EncodeToString([]byte(token + "|" + strconv.FormatInt(uid, 10)))

	return &Token{
		Token:       token,
		EncodeToken: encodeToken,
		Uid:         uid,
		UserLogin:   userLogin,
		Expire:      t,
	}
}

func GetUid(c *gin.Context, e ...bool) int64 {
	var i int64
	if uid, exists := c.Get("uid"); exists {
		i = uid.(int64)
	} else if len(e) > 0 && e[0] {
		i = int64(163)
		//panic(http_error.JwtError)
	}
	return i
}

func GetToken(c *gin.Context, e ...bool) string {
	var i string
	if token, exists := c.Get("token"); exists {
		i = token.(string)
	} else if len(e) > 0 && e[0] {
		panic(http_error.JwtError)
	}
	return i
}

type keys []string

func CheckSign(c *gin.Context) bool {

	raw := map[string]interface{}{}
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	decoder := json.NewDecoder(bytes.NewBuffer(bodyBytes))
	decoder.UseNumber()
	decoder.Decode(&raw)

	keys := []string{}
	for k := range raw {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	signStr := ""
	sign := ""
	for _, v := range keys {

		rawV := raw[v]
		value := stringify.ToString(rawV)
		if v == "sign" {
			sign = value
			continue
		}

		signStr += v + "=" + value + "&"
	}

	signStr += constant.WsSignKey

	return sign == cryp.MD5(signStr)

}
