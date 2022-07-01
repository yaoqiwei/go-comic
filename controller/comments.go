package controller

import (
	"fehu/model/body"
	"fehu/service/comments"
	"fehu/util/jwt"
	"fehu/util/request"
	"github.com/gin-gonic/gin"
)

type CommentsRegister struct {
}

func CommentRegister(router *gin.RouterGroup, needLoginedRouter *gin.RouterGroup) {
	c := CommentsRegister{}
	needLoginedRouter.POST("/comments", c.commentsList)
}

// @Summary 获取评论
// @Description 获取评论
// @Tags 动态相关
// @Accept json
// @Param param body body.NewsCommentParam true "-"
// @Success 0 {object} success.ListData
// @Success 0 {object} newsComment.CommentList
// @Router /comments/list [post]
func (*CommentsRegister) commentsList(c *gin.Context) {
	var param body.CommentsListParam
	request.Bind(c, &param)
	uid := jwt.GetUid(c, true)
	comments.GetCommentsList(param, uid)
}
