package cartoonComments

import (
	"fehu/common/lib/gorm"
	"fehu/model/body"
	"time"
)

// CartoonComments 对应数据表cartoon_comments用户评论表
type CartoonComments struct {
	Id        int64     `json:"id"`
	CartoonId int64     `json:"cartoonId"` //漫画id
	Genre     byte      `json:"genre"`     //评论的类型 1.漫画评论 2.章节评论
	Uid       int64     `json:"uid"`       //用户id
	ToUid     int64     `json:"toUid"`     //被回复的用户id
	ReplyId   int64     `json:"replyId"`   //回复的评论id
	Comment   string    `json:"comment"`   //评论内容
	RootId    int64     `json:"rootId"`    //根评论id
	Disabled  byte      `json:"disabled"`  //是否显示 0.显示 1.不显示
	Status    byte      `json:"status"`    //评论状态 0.待审核 1.通过审核 2.审核失败
	CreateAt  time.Time `json:"createAt"`  //评论时间
}

// CommentsBase 评论基础信息
type CommentsBase struct {
	CartoonComments
	ReplyCount int `json:"replyCount"` //回复数
	LikeCount  int `json:"likeCount"`  //点赞数
	Hot        int `json:"hot"`        //热度
}

// GetCommentsBase 获取基础评论数据
func GetCommentsBase(v body.CommentsListParam) []*CommentsBase {
	list := make([]*CommentsBase, 0)
	db := gorm.Db.Table("? AS t1", &CartoonComments{}).
		Select("t1.id,t1.cartoon_id,t1.genre,t1.uid,t1.to_uid,t1.reply_id,t1.comment,"+
			"t1.root_id,t1.create_at,t2.reply_count,t2.like_count,t2.hot").
		Joins("join user_comment_count t2 on t2.comment_id = t1.id").
		Where("cartoon_id=? and genre=? and status=1 and disabled=0 and reply_id=0", v.Id, v.Genre)
	switch v.SortType {
	case 1:
		db.Order("create_at des")
	case 2:
		db.Order("t2.hot desc")
	}
	db.Limit(v.GetLength()).Offset(v.GetOffset()).Scan(&list)
	return list
}

// GetCommentsReply 获取评论回复
func GetCommentsReply(rootId int64, p body.PageParam, sortType byte) []*CommentsBase {
	list := make([]*CommentsBase, 0)
	db := gorm.Db.Table("? AS t1", &CartoonComments{}).
		Select("t1.id,t1.cartoon_id,t1.genre,t1.uid,t1.to_uid,t1.reply_id,t1.comment,"+
			"t1.root_id,t1.create_at,t2.reply_count,t2.like_count,t2.hot").
		Joins("join user_comment_count t2 on t2.comment_id = t1.id").
		Where("root_id=? and status=1 and disabled=0", rootId)
	switch sortType {
	case 1:
		db.Order("create_at des")
	case 2:
		db.Order("t2.hot desc")
	}
	db.Limit(p.GetLength()).Offset(p.GetOffset()).Scan(&list)
	return list
}
