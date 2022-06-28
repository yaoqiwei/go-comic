package jwt_test

import (
	"bytes"
	"encoding/json"
	"fehu/constant"
	"fehu/util/cryp"
	"sort"
	"testing"

	"fehu/util/stringify"
)

func TestSign(t *testing.T) {

	data := `{
		"uid": 45175,
		"liveuid": 45666,
		"timeStamp": 1615282681,
		"sign": "46e7ed4ec0048b157bfbdd9d0dd707aa"
	  }`

	raw := map[string]interface{}{}
	decoder := json.NewDecoder(bytes.NewBufferString(data))
	decoder.UseNumber()
	decoder.Decode(&raw)

	if !CheckSign(raw) {
		t.Error("签名错误")
	}
}

func CheckSign(raw map[string]interface{}) bool {

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
