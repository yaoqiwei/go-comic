package comments

import (
	"fehu/model/body"
	"fehu/service/comments/cartoonComments"
	"fehu/service/comments/cartoonCommentsLike"
	"fmt"
)

// CommentsList 评论列表
type CommentsList struct {
	Base      cartoonComments.CommentsBase
	ReplyList []*CommentsList `json:"replyList,omitempty"` // 该评论的回复列表
	IsLiked   bool            `json:"isLiked"`             // 是否点赞 true：已点赞，false 未点赞
}

// GetCommentsList 获取评论列表
func GetCommentsList(v body.CommentsListParam, uid int64) {
	res := cartoonComments.GetCommentsBase(v)
	list := make([]*CommentsList, 0)
	for _, v := range res {
		isLiked := cartoonCommentsLike.CheckCommentLikeExist(uid, v.Id)
		param := body.CartoonIdListParam{}
		param.Length = 3
		GetCommentReply(uid, v.Id, param, 2)
		list = append(list, &CommentsList{
			Base:    *v,
			IsLiked: isLiked,
		})
		fmt.Println(v)
	}
}

// GetCommentReply 获取评论回复列表
func GetCommentReply(uid, rootId int64, p body.PageParam, sortType byte) []*CommentsList {
	res := cartoonComments.GetCommentsReply(rootId, p, sortType)
	list := make([]*CommentsList, 0)
	for _, v := range res {
		isLiked := cartoonCommentsLike.CheckCommentLikeExist(uid, v.Id)
		list = append(list, &CommentsList{
			Base:    *v,
			IsLiked: isLiked,
		})
	}
	return list
}
