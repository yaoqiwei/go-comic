package http_error

var VideoNotExist = HttpError{
	ErrorCode: 2001,
	ErrorMsg:  "视频不存在！",
}

var VipVideo = HttpError{
	ErrorCode: 2002,
	ErrorMsg:  "无权限获取",
}

var VideoNotBought = HttpError{
	ErrorCode: 2003,
	ErrorMsg:  "该视频需要购买",
}

var VideoPreviewClosed = HttpError{
	ErrorCode: 2004,
	ErrorMsg:  "视频试看功能已关闭！",
}

var VideoIsLiked = HttpError{
	ErrorCode: 2005,
	ErrorMsg:  "该视频已点赞",
}

var VideoIsCollected = HttpError{
	ErrorCode: 2006,
	ErrorMsg:  "该视频已收藏",
}

var VideoLimitPreview = HttpError{
	ErrorCode: 2009,
	ErrorMsg:  "今日试看次数达限 请充值会员无限观看",
}

var NotVipVideo = HttpError{
	ErrorCode: 2010,
	ErrorMsg:  "非VIP视频",
}

var IsBought = HttpError{
	ErrorCode: 2011,
	ErrorMsg:  "该视频已经购买，无需再次购买！",
}

var UserNotVideoVip = HttpError{
	ErrorCode: 2012,
	ErrorMsg:  "用户不是点播VIP",
}

var NoPermissionToCommentVideo = HttpError{
	ErrorCode: 2013,
	ErrorMsg:  "无权限评论",
}

var NotPointVideo = HttpError{
	ErrorCode: 2014,
	ErrorMsg:  "非钻石视频",
}
var VideoLinkNotExist = HttpError{
	ErrorCode: 2015,
	ErrorMsg:  "短视频链接不存在！",
}
var CommentIsLiked = HttpError{
	ErrorCode: 2016,
	ErrorMsg:  "该评论已点赞",
}
var CommentNotLike = HttpError{
	ErrorCode: 2016,
	ErrorMsg:  "该评论未点赞",
}
