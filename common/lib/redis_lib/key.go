package redis_lib

import (
	"fehu/model/http_error"

	"fehu/util/stringify"
)

type RedisKeyCode string

var RedisKey = map[RedisKeyCode]string{
	"UPLOAD_CACHE":            "uploadCache",
	"AUTH_UPLOAD_CACHE":       "uploadCache:auth",
	"IM_UPLOAD_CACHE":         "uploadCache:im",
	"LIVE_UPLOAD_CACHE":       "uploadCache:live",
	"NEWS_UPLOAD_CACHE":       "uploadCache:news",
	"FEEDBACK_UPLOAD_CACHE":   "uploadCache:feedback",
	"GIFT_CACHE":              "giftCache",
	"MESSAGE_CACHE":           "messageCache",
	"IM_SEND_DATA":            "imSendData",
	"BARRAGE_CACHE":           "barrageCache",
	"LIVE_PK":                 "LivePK",
	"LIVE_PK_INFO":            "livePkInfo",
	"LIVE_CONNECT":            "LiveConnect",
	"LIVE_CONNECT_PULL":       "LiveConnectPull",
	"VIP":                     "vipV2",
	"LIVE_VIP":                "liveVip",
	"VIDEO_CACHE":             "videoCache",
	"RED_USER_WINNING":        "red_user_winning",
	"RED_LIST":                "red_list",
	"LEVEL":                   "level",
	"CODE_IP_LIMIT":           "userCodeIpLimit",
	"BIND_CODE":               "bindCode",
	"FIND_CODE":               "findCode",
	"REG_CODE":                "regCode",
	"CONFIG_PRI":              "getConfigPri",
	"CONFIG_PUB":              "getConfigPub",
	"JACKPOT_SET":             "jackpotset",
	"ROOM_USER":               "roomUser",
	"LIVE_USER_INFO":          "liveUserInfo",
	"TIME_CHARGE":             "timeCharge",
	"ZOMBIE_SET":              "zombieSet",
	"GIFT":                    "getGiftList",
	"GUARD":                   "guard_list",
	"LIVE_CLASS":              "liveAllClass",
	"LIVE_LEVEL":              "levelanchor",
	"TOKEN":                   "token",
	"AREA":                    "area",
	"LOCK":                    "lock",
	"LOCK_CHANGE_BAG":         "lock:changeBag",
	"LOCK_CROWD_LAUNCH":       "lock:crowdLaunch",
	"LOCK_CROWD":              "lock:crowd",
	"LOCK_CHAT_ROOM":          "lock:chatRoom",
	"LOCK_BANK_UPD":           "lock:bankUpd",
	"LOCK_UPD_THRESHOLD":      "lock:updThreshold",
	"CROWD_DRAW_TOTAL":        "crowdDrawTotal",
	"CROWD_FUNDING_REWARD":    "crowdFundingReward",
	"ONLINE_CUSTOMER_SERVICE": "onlineCustomerService",
	"SITE_OPTIONS":            "siteOptions",
	"AGT_LEVEL_CONFIG":        "agtLevelConfig",
	"VIDEO_INFO":              "videoInfo",
	"VIDEO_SEARCH_DATA":       "videoSearchData",
	"INVITE_IP_TO_AGENT_ID":   "inviteIpToAgentId",
	"AGENT_TOTAL_RESULT":      "agentTotalResult",
	"LAST_USER_OPTION_TIME":   "lastUserOptionTime",
	"LOCK_CARD_BUY":           "lock:cardBuy",
	"LOCK_CARD_EXCHANGE":      "lock:cardExchage",
	"AGT_STATISTIC":           "agtStatistic",
	"ADD_COMMENT_COUNT":       "addCommentCount",
	"USER_VIDEO_COUNT_DATA":   "userVideoCountData",
	"ADD_NEWS_COMMENT_COUNT":  "addNews+CommentCount",
}

func GetRedisKey(keyCode RedisKeyCode, others ...interface{}) string {
	key, ok := RedisKey[keyCode]
	if !ok {
		panic(http_error.NoRedisKey)
	}

	for _, v := range others {
		key += ":" + stringify.ToString(v)
	}
	return key
}

func Lock(keyCode RedisKeyCode, others ...interface{}) {
	set, _ := SetNx(GetRedisKey(keyCode, others...), 1, 5, "")
	if !set {
		panic(http_error.FrequentOperations)
	}
}

func UnLock(keyCode RedisKeyCode, others ...interface{}) {
	Del(GetRedisKey(keyCode, others...), "")
}

func BLock(keyCode RedisKeyCode, others ...interface{}) {
	set, _ := SetNx(GetRedisKey(keyCode, others...), 1, 5, "")
	if !set {
		_, err := BRPopInt(GetRedisKey(keyCode, others...)+".B", 5, "")
		if err != nil {
			panic(http_error.FrequentOperations)
		}
	}
}

func BUnLock(keyCode RedisKeyCode, others ...interface{}) {
	key := GetRedisKey(keyCode, others...)
	keyB := GetRedisKey(keyCode, others...) + ".B"
	RPush(keyB, 1, "")
	Expire(key, 5, "")
	Expire(keyB, 5, "")
}
