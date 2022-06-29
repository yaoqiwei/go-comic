package base

import (
	"fehu/common/lib/mysql_lib"
	"fehu/model/http_error"
	"fehu/util/math"
	"html"
	"strings"
)

func Error(code int) {
	panic(http_error.HttpError{ErrorCode: code})
}

func CompleteImgurl(file string) string {

	if file == "" {
		return file
	}

	if strings.Index(file, "http") == 0 {
		return html.UnescapeString(file)
	}

	config := GetConfigPub()
	return html.UnescapeString(config.ImgUrl + file)
}

func CompleteVideoImgUrl(file string) string {

	if file == "" {
		return file
	}

	if strings.Index(file, "http") == 0 {
		return html.UnescapeString(file)
	}

	config := GetConfigPub()
	return html.UnescapeString(config.VideoImgUrl + file)
}

func CompleteVideoUrl(file string) string {

	if file == "" {
		return file
	}

	if strings.Index(file, "http") == 0 {
		return html.UnescapeString(file)
	}

	config := GetConfigPub()
	return html.UnescapeString(config.VideoUrl + file)
}

func CompleteVideoDownUrl(file string) string {

	if file == "" {
		return file
	}

	if strings.Index(file, "http") == 0 {
		return html.UnescapeString(file)
	}

	config := GetConfigPub()
	return html.UnescapeString(config.VideoDownUrl + file)
}

type ChargeRule struct {
	Id                int     `db:"id" json:"id"`
	Name              string  `db:"name" json:"name"`
	Coin              int64   `db:"coin" json:"coin"`
	Vip               int64   `db:"vip" json:"vip"`                              // vip天数
	Money             float64 `db:"money" json:"money"`                          // 当前价格
	RawPrice          float64 `db:"raw_price" json:"rawPrice"`                   // 原价
	Give              int64   `db:"give" json:"give"`                            // 赠送数
	Type              byte    `db:"type" json:"type"`                            // 类型，1钻石2VIP3LIVEVIP
	BuyDiscount       byte    `db:"buy_discount" json:"buyDiscount"`             // 购片折扣，100不打折，0免费
	Level             byte    `db:"level" json:"level"`                          // 标注等级
	UnlimitedDownload bool    `db:"unlimited_download" json:"unlimitedDownload"` // 无限下载，1是0否
	Cdn               bool    `db:"cdn" json:"cdn"`                              // CDN权限，1是0否
	VideoReply        bool    `db:"video_reply" json:"videoReply"`               // 视频评论权限
	NewsVip           bool    `db:"news_vip" json:"newsVip"`                     // 赠送动态会员
	Per               *byte   `json:"per,omitempty"`                             // 折扣
}

func (r *ChargeRule) WithPer() *ChargeRule {
	r.Per = new(byte)
	if r.RawPrice > 0 {
		*r.Per = byte(math.RoundedFixed(r.Money/r.RawPrice*100, 0))
	}
	return r
}

type ChargeRuleList []*ChargeRule

func GetRule(id int) *ChargeRule {
	info := &ChargeRule{}
	mysql_lib.DB("charge_rules").Dest(info).Query("WHERE id=?", id).FetchOne()
	return info
}

func GetChargeRules() ChargeRuleList {
	list := ChargeRuleList{}
	mysql_lib.DB("charge_rules").Dest(&list).Query("WHERE type=1 ORDER BY orderno").Select()
	return list
}

func GetVipRules() ChargeRuleList {
	list := ChargeRuleList{}
	mysql_lib.DB("charge_rules").Dest(&list).Query("WHERE type=2 ORDER BY orderno").Select()
	return list
}

func GetLiveVipRules() ChargeRuleList {
	list := ChargeRuleList{}
	mysql_lib.DB("charge_rules").Dest(&list).Query("WHERE type=3 ORDER BY orderno").Select()
	return list
}

func GetChargeRule(id int) *ChargeRule {
	rules := GetChargeRules()
	for _, v := range rules {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func GetVipRule(id int) *ChargeRule {
	rules := GetVipRules()
	for _, v := range rules {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func GetVipRuleByLevel(level byte) *ChargeRule {
	rules := GetVipRules()
	for _, v := range rules {
		if v.Level == level {
			return v
		}
	}
	return nil
}

func CheckMaintain() {
	config := GetConfigPub()
	if config.MaintainSwitch == "1" {
		panic(http_error.HttpError{
			ErrorCode: http_error.Maintain.ErrorCode,
			ErrorMsg:  config.MaintainTips,
		})
	}
}
