package cartoonCommentsLike

import (
	"fehu/common/lib/gorm"
	"time"
)

// CartoonCommentsLike 对应数据库cartoon_comments_like评论点赞表
type CartoonCommentsLike struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"`       //用户id
	CommentId int64     `json:"commentId"` //评论id
	ToUid     int64     `json:"toUid"`     //被喜欢的评论者id
	CartoonId int64     `json:"cartoonId"` //漫画id
	Genre     byte      `json:"genre"`     //评论的类型 1.漫画评论 2.章节评论
	CreateAt  time.Time `json:"createAt"`  //创建时间
}

// CheckCommentLikeExist 是否点赞这条评论
func CheckCommentLikeExist(uid, commentId int64) bool {
	return gorm.Db.Where("uid=? and comment_id=?", uid, commentId).
		First(&CartoonCommentsLike{}).RowsAffected > 0
}
