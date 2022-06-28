package http_error

var LiveNotExist = HttpError{
	ErrorCode: 12001,
	ErrorMsg:  "直播间不存在！",
}
var BarrageEmpty = HttpError{
	ErrorCode: 12002,
	ErrorMsg:  "弹幕内容不能为空",
}

var NoGiftInfo = HttpError{
	ErrorCode: 12003,
	ErrorMsg:  "礼物信息不存在！",
}

var GuardGift = HttpError{
	ErrorCode: 12004,
	ErrorMsg:  "该礼物是守护专属礼物奥~",
}

var BagNotEnoughGift = HttpError{
	ErrorCode: 12005,
	ErrorMsg:  "背包中数量不足",
}

var LiveEnd = HttpError{
	ErrorCode: 12006,
	ErrorMsg:  "直播已结束！",
}

var LiveNotChangeType = HttpError{
	ErrorCode: 12007,
	ErrorMsg:  "该房间非扣费房间！",
}

var RedNotExist = HttpError{
	ErrorCode: 12008,
	ErrorMsg:  "红包不存在！",
}

var IsMyLiveRoom = HttpError{
	ErrorCode: 12009,
	ErrorMsg:  "不能进入自己的直播间",
}

var IsKicked = HttpError{
	ErrorCode: 12010,
	ErrorMsg:  "您已被踢出房间",
}

var SuperLimitRoom = HttpError{
	ErrorCode: 12011,
	ErrorMsg:  "超管不能进入1v1房间",
}

var LiveChangeErr = HttpError{
	ErrorCode: 120012,
	ErrorMsg:  "房间费用有误",
}

var RedCoinErr = HttpError{
	ErrorCode: 12013,
	ErrorMsg:  "请输入正确的金额",
}

var RedNumErr = HttpError{
	ErrorCode: 12014,
	ErrorMsg:  "请输入正确的个数",
}

var RedCoinNumErr = HttpError{
	ErrorCode: 12015,
	ErrorMsg:  "红包数量不能超过红包金额",
}

var RedDesLong = HttpError{
	ErrorCode: 12016,
	ErrorMsg:  "红包名称最多20个字",
}

var RedSendFail = HttpError{
	ErrorCode: 12017,
	ErrorMsg:  "发送失败，请重试",
}

var MicLevelLimit = HttpError{
	ErrorCode: 12018,
	ErrorMsg:  "用户等级达到1级才可与主播连麦哦~",
}
var LiveRoomNoMic = HttpError{
	ErrorCode: 12019,
	ErrorMsg:  "主播未开启连麦功能哦~",
}
var LiveBanned = HttpError{
	ErrorCode: 12020,
	ErrorMsg:  "已被禁播~",
}

var LiveLimit = HttpError{
	ErrorCode: 12021,
	ErrorMsg:  "等级小于1级，不能直播~",
}

var PasswordEmpty = HttpError{
	ErrorCode: 12022,
	ErrorMsg:  "密码不能为空",
}

var PriceEmpty = HttpError{
	ErrorCode: 12023,
	ErrorMsg:  "价格不能小于等于0",
}

var ThumbGetErr = HttpError{
	ErrorCode: 12024,
	ErrorMsg:  "封面图片获取失败",
}

var CreateRoomFail = HttpError{
	ErrorCode: 12025,
	ErrorMsg:  "开播失败，请重试",
}

var MergeVideoStreamFail = HttpError{
	ErrorCode: 12026,
	ErrorMsg:  "连麦混流错误",
}

var NoTurnConn = HttpError{
	ErrorCode: 12027,
	ErrorMsg:  "信息错误",
}

var CanNotSetAdminSelf = HttpError{
	ErrorCode: 12028,
	ErrorMsg:  "不能设置自己为管理",
}

var OverAdminCount = HttpError{
	ErrorCode: 12030,
	ErrorMsg:  "最多设置5个管理员!",
}

var NotAdmin = HttpError{
	ErrorCode: 12031,
	ErrorMsg:  "您不是该直播间的管理员，无权操作",
}

var IsAdmin = HttpError{
	ErrorCode: 12032,
	ErrorMsg:  "无权操作",
}

var IsGrandGuard = HttpError{
	ErrorCode: 12033,
	ErrorMsg:  "对方是尊贵守护，不能禁言",
}

var PkDiffTime = HttpError{
	ErrorCode: 12034,
	ErrorMsg:  "时间不匹配",
}

var IsNotZombie = HttpError{
	ErrorCode: 12035,
	ErrorMsg:  "未开启僵尸粉",
}

var IsVisited = HttpError{
	ErrorCode: 12036,
	ErrorMsg:  "用户已访问",
}

var ZombieTimesOut = HttpError{
	ErrorCode: 12037,
	ErrorMsg:  "次数已满",
}

var NotSuper = HttpError{
	ErrorCode: 12038,
	ErrorMsg:  "你不是超管，无权操作",
}

var ConnectNotLive = HttpError{
	ErrorCode: 12039,
	ErrorMsg:  "连麦对象没有开启直播~",
}

var ConnectNotExist = HttpError{
	ErrorCode: 12040,
	ErrorMsg:  "连麦对象没有找到~",
}

var LiveNowLimit = HttpError{
	ErrorCode: 12041,
	ErrorMsg:  "主播刚开播，请稍后再试~",
}

var NoGiftCount = HttpError{
	ErrorCode: 12042,
	ErrorMsg:  "请选择礼物数量",
}

var StartCrowdFundingFail = HttpError{
	ErrorCode: 12043,
	ErrorMsg:  "发起抽奖失败",
}

var ExistCrowdFunding = HttpError{
	ErrorCode: 12044,
	ErrorMsg:  "已发起抽奖",
}

var NotExistCrowdFunding = HttpError{
	ErrorCode: 12045,
	ErrorMsg:  "未发起抽奖",
}

var CrowdFundingExpired = HttpError{
	ErrorCode: 12046,
	ErrorMsg:  "抽奖已结束",
}

var DrawCrowdFundingFail = HttpError{
	ErrorCode: 12047,
	ErrorMsg:  "抽奖失败",
}

var LiveVipNeed = HttpError{
	ErrorCode: 12048,
	ErrorMsg:  "该直播间需要开通直播VIP方可进入",
}

var BarrageLevelLimit = HttpError{
	ErrorCode: 12049,
	ErrorMsg:  "直播等级达到4级即可发言，赠送礼物可提高用户等级哦！",
}

var CanNotSendGift = HttpError{
	ErrorCode: 12050,
	ErrorMsg:  "该房间禁止发送礼物",
}
