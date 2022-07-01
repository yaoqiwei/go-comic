package body

// CommentsListParam 获取评论需传入参数
type CommentsListParam struct {
	Id       int64 `json:"id" form:"id" binding:"required"`
	Genre    byte  `json:"genre" form:"genre"`        //评论的类型 1.漫画评论 2.章节评论
	SortType byte  `json:"sortType" form:"sortType" ` //排序类型 1：按照添加时间倒序 2：按照热度
	PageTrait
}

// CartoonIdListParam 漫画列表传入参数
type CartoonIdListParam struct {
	CartoonId int64 `json:"cartoonId" form:"cartoonId" binding:"required"` //漫画id
	Genre     byte  `json:"genre" form:"genre"`                            //评论的类型 1.漫画评论 2.章节评论
	PageTrait
}
