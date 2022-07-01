package base

import (
	"encoding/json"
	"fehu/common/lib/mysql_lib"
	"fehu/common/lib/redis_lib"
	"reflect"
	"strings"
	"sync"
	"time"
)

type ConfigPri struct {
	BarrageFee          string `json:"barrage_fee"`          // 弹幕费用
	NewsVipFee          string `json:"news_vip_fee"`         // 线下VIP费用
	WithdrawalThreshold string `json:"withdrawal_threshold"` // 提现阈值
	RegReward           string `json:"reg_reward"`           // 注册奖励
	ChatServer          string `json:"chatserver"`           // 聊天服务器带端口
	CdnSwitch           string `json:"cdn_switch"`           // CDN, 2.腾讯云/4.网宿

	TxAppId   string `json:"tx_appid"`    // 直播appid
	TxBizid   string `json:"tx_bizid"`    // 直播bizid
	TxPushkey string `json:"tx_push_key"` // 直播API鉴权key
	TxPush    string `json:"tx_push"`     // 直播推流域名
	TxPull    string `json:"tx_pull"`     // 直播播流域名
	WsPush    string `json:"ws_push"`     // 网宿推流域名
	WsPull    string `json:"ws_pull"`     // 网宿播流地址
	WsApn     string `json:"ws_apn"`      // 网宿发布点

	PayUrl   string `json:"payUrl"`   // 支付地址
	PayMerId string `json:"payMerId"` // 支付商户号
	PayKey   string `json:"payKey"`   // 支付密钥

	ImServer string `json:"imserver"` // 客服聊天ws地址配置

	IpLimitSwitch  string `json:"iplimit_switch"`  // 限制每个ip每天发送验证码次数
	IpLimitTimes   string `json:"iplimit_times"`   // 短信验证码IP限制次数
	SendcodeWwitch string `json:"sendcode_switch"` // 短信验证码开关,关闭后不再发送真实验证码，采用返回密码

	AliappSwitch string `json:"aliapp_switch"` // 支付宝支付开关
	WxSwitch     string `json:"wx_switch"`     // 支微信支付开关
	IosWwitch    string `json:"ios_switch"`    // IOS支付开关

	MicLimit     string `json:"mic_limit"`     // 连麦等级限制
	LevelIsLimit string `json:"level_islimit"` // 直播等级控制是否开启
	LevelLimit   string `json:"level_limit"`   // 直播限制等级
	MsgLimit     string `json:"msg_limit"`     // 直播发言限制等级
	UserListTime string `json:"userlist_time"` // 用户列表请求间隔

	ImAutoSend string `json:"imautosend"` // 客服自动回复内容
}

type ConfigMysql struct {
	OptionValue string `db:"option_value"`
}

type LiveType struct {
}

type ConfigPubBase struct {
	MaintainSwitch string `json:"maintain_switch"`  // 网站维护
	MaintainTips   string `json:"maintain_tips"`    // 维护提示
	NewsAnno       string `json:"newsAnno"`         // 动态公告
	Sitename       string `json:"sitename"`         // 网站标题
	Site           string `json:"site"`             // 网站域名
	SiteAdminEmail string `json:"site_admin_email"` // 官方邮箱
	JoinUrl        string `json:"join_url"`         // 加盟URL
	Copyright      string `json:"copyright"`        // 版权信息
	NameCoin       string `json:"name_coin"`        // 钻石名称
	NameVotes      string `json:"name_votes"`       // 映票名称
	AgentPic       string `json:"agent_pic"`        // 全民代理图片
	ApkEwm         string `json:"apk_ewm"`          // android版下载二维码
	IpaEwm         string `json:"ipa_ewm"`          // iPhone版下载二维码

	Isup       string `json:"isup"`        // 强制更新
	ApkVer     string `json:"apk_ver"`     // APK版本号
	ApkVerMin  string `json:"apk_ver_min"` // APK最小版本号
	ApkUrl     string `json:"apk_url"`     // APK下载链接
	ApkOut     string `json:"apk_out"`     // APK外部链接
	ApkDes     string `json:"apk_des"`     // APk更新说明
	IpaVer     string `json:"ipa_ver"`     // IPA版本号
	IpaVerMin  string `json:"ipa_ver_min"` // IPA最小版本号
	IosShelves string `json:"ios_shelves"` // IPA上架版本号
	IpaUrl     string `json:"ipa_url"`     // IPA下载链接
	IpaDes     string `json:"ipa_des"`     // IPA更新说明

	ShareSiteurl    string `json:"share_siteurl"`     // 推广域名
	ShareTitle      string `json:"share_title"`       // 直播分享标题
	ShareDes        string `json:"share_des"`         // 直播分享话术
	VideoShareTitle string `json:"video_share_title"` // 短视频分享标题
	VideoShareDes   string `json:"video_share_des"`   // 短视频分享话术

	SproutKey       string `json:"sprout_key"`       // 萌颜授权码
	SproutWhite     string `json:"sprout_white"`     // 美白默认值
	SproutSkin      string `json:"sprout_skin"`      // 磨皮默认值
	SproutSaturated string `json:"sprout_saturated"` // 饱和默认值
	SproutPink      string `json:"sprout_pink"`      // 粉嫩默认值
	SproutEye       string `json:"sprout_eye"`       // 大眼默认值
	SproutFace      string `json:"sprout_face"`      // 瘦脸默认值

	ImgUrl        string `json:"webImgurl"`       // 网站图片地址
	VideoImgUrl   string `json:"webVideoimgurl"`  // 视频图片地址
	VideoUrl      string `json:"webVideourl"`     // 视频播放地址
	VideoDownUrl  string `json:"webVideoDownurl"` // 视频下载地址
	AgentAdminUrl string `json:"agentAdminUrl"`   // 代理后台地址
	AgtShareQQ    string `json:"agtShare_qq"`     // 代理推广地址QQ
	AgtShareWc    string `json:"agtShare_wc"`     // 代理推广地址微信
	AgtShareSite  string `json:"agtShare_site"`   // 代理推广网站地址
	AgtShareApp   string `json:"agtShare_app"`    // 代理推广APP下载地址
	ChatSite      string `json:"chatSite"`        // 聊天地址
	PreviewOn     string `json:"previewOn"`       // 是否试看
	PreviewRange  string `json:"previewRange"`    // 试看范围
	PreviewNum    string `json:"previewNum"`      // 视频试看次数
}

type ConfigPubMysqlRaw struct {
	ConfigPubBase
	ShareType string `json:"share_type"` // 分享方式
}

type ConfigPub struct {
	ConfigPubBase
	ShareType []string `json:"share_type"` // 分享方式
}

type JackpotSet struct {
	Switch      string `json:"switch"`
	LuckAnchor  string `json:"luck_anchor"`
	LuckJackpot string `json:"luck_jackpot"`
}

func GetConfigPri() ConfigPri {

	var configpri ConfigPri

	err := redis_lib.GetObject(redis_lib.GetRedisKey("CONFIG_PRI"), &configpri, "")

	if err == nil {
		return configpri
	}

	var configPriMysql ConfigMysql
	mysql_lib.FetchOne(&configPriMysql, "SELECT option_value FROM comic_options WHERE option_name='configpri'")
	redis_lib.Set(redis_lib.GetRedisKey("CONFIG_PRI"), configPriMysql.OptionValue, 60*60*24*300, "")

	json.Unmarshal([]byte(configPriMysql.OptionValue), &configpri)
	return configpri
}

type Guide struct {
	Thumb string `db:"thumb" json:"thumb"` // 图片、视频链接
	Href  string `db:"href" json:"href"`   // 页面链接
}

type GuideList []*Guide

type GuideConfig struct {
	Switch string   `db:"switch" json:"switch"` // 开关
	Type   string   `db:"type" json:"type"`     // 类型
	Time   string   `db:"time" json:"time"`     // 时间
	List   []*Guide `json:"list"`
}

func GetGuide() GuideConfig {

	var guideConfig GuideConfig
	var configMysql ConfigMysql
	mysql_lib.FetchOne(&configMysql, "SELECT option_value FROM cmf_options WHERE option_name='guide'")

	json.Unmarshal([]byte(configMysql.OptionValue), &guideConfig)

	guideConfig.List = []*Guide{}
	if guideConfig.Switch == "1" {

		list := GuideList{}
		mysql_lib.Select(&list, "SELECT thumb,href FROM cmf_guide WHERE type=? ORDER BY orderno,uptime DESC", guideConfig.Type)
		for _, v := range list {
			v.Thumb = CompleteImgurl(v.Thumb)
		}
		guideConfig.List = list
	}
	return guideConfig

}

var globalConfigPub *ConfigPub
var lockGlobalConfigPub sync.RWMutex

func getGlobalConfigPub() *ConfigPub {
	lockGlobalConfigPub.RLock()
	defer lockGlobalConfigPub.RUnlock()
	return globalConfigPub
}
func setGlobalConfigPub(c *ConfigPub) {
	lockGlobalConfigPub.Lock()
	defer lockGlobalConfigPub.Unlock()
	globalConfigPub = c

	go func() {
		time.Sleep(5 * time.Second)
		globalConfigPub = nil
	}()
}

func GetConfigPub() ConfigPub {

	var configpubRaw ConfigPubMysqlRaw
	var configPub *ConfigPub

	if data := getGlobalConfigPub(); data != nil {
		return *data
	}

	err := redis_lib.GetObject(redis_lib.GetRedisKey("CONFIG_PUB"), &configpubRaw, "")

	if err != nil {
		var configPriMysql ConfigMysql
		mysql_lib.FetchOne(&configPriMysql, "SELECT option_value FROM cmf_options WHERE option_name='configpub'")

		redis_lib.Set(redis_lib.GetRedisKey("CONFIG_PUB"), configPriMysql.OptionValue, 60*60*24*300, "")
		json.Unmarshal([]byte(configPriMysql.OptionValue), &configpubRaw)
	}

	configPub = &ConfigPub{}
	configPub.ConfigPubBase = configpubRaw.ConfigPubBase

	for _, v := range []string{"ShareType"} {

		str := reflect.ValueOf(&configpubRaw).Elem().FieldByName(v).Interface().(string)
		arr := make([]string, 0)

		if str != "" {
			str = strings.Replace(str, "，", ",", -1)
			arr = strings.Split(str, ",")
		}

		r := reflect.ValueOf(arr)
		if v == "LiveType" {
			arr3 := make([]([]string), 0)
			for _, v2 := range arr {
				str2 := strings.Replace(v2, "；", ";", -1)
				arr2 := strings.Split(str2, ";")
				arr3 = append(arr3, arr2)
			}
			r = reflect.ValueOf(arr3)
		}

		reflect.ValueOf(configPub).Elem().FieldByName(v).Set(r)
	}

	setGlobalConfigPub(configPub)

	return *configPub
}

func GetJackpot() JackpotSet {

	var jackpot JackpotSet

	err := redis_lib.GetObject(redis_lib.GetRedisKey("JACKPOT_SET"), &jackpot, "")

	if err == nil {
		return jackpot
	}

	var conf ConfigMysql
	mysql_lib.FetchOne(&conf, "SELECT option_value FROM cmf_options WHERE option_name='jackpot'")
	redis_lib.Set(redis_lib.GetRedisKey("JACKPOT_SET"), conf.OptionValue, 60*60*24*300, "")

	json.Unmarshal([]byte(conf.OptionValue), &jackpot)
	return jackpot
}
