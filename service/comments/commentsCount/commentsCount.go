package commentsCount

// CartoonCommentsCount 对应数据库users_comment_count视频评论数量表
type CartoonCommentsCount struct {
	Id         int64 `json:"id"`
	CommentId  int64 `json:"commentId"`  //评论id
	ReplyCount int   `json:"replyCount"` //回复数
	LikeCount  int   `json:"likeCount"`  //点赞数
	Hot        int   `json:"hot"`        //热度
}

// GetCommentsReplyCount 获取回复评论数量
func GetCommentsReplyCount() {

}
