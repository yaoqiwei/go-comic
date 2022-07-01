package body

import (
	"fehu/util/stringify"
	"time"
)

type PageTrait struct {
	Offset *int `json:"offset" form:"offset" uri:"offset" example:"1"` // 起始位置
	Page   int  `json:"page" form:"page" uri:"page" example:"1"`       // 页数
	Length int  `json:"length" form:"length" uri:"length" example:"1"` // 每页数量
}

type PageParam interface {
	GetOffset() int
	GetPage() int
	GetLength() int
}

type IdParam struct {
	Id int64 `json:"id" form:"id" binding:"required"`
}

func (p PageTrait) GetOffset() int {
	if p.Offset != nil {
		return *p.Offset
	}
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Length * (p.Page - 1)
}

func (p PageTrait) GetPage() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return p.Page
}

func (p PageTrait) GetLength() int {
	return p.Length
}

type DateTimeTrait struct {
	StartTime *string `json:"startTime"` // 开始日期, 格式: YYYY-MM-DD HH:ii:ss
	EndTime   *string `json:"endTime"`   // 结束日期, 格式: YYYY-MM-DD HH:ii:ss
}

func (p DateTimeTrait) GetStart() *time.Time {
	return parseTime(p.StartTime)
}

func (p DateTimeTrait) GetEnd() *time.Time {
	return parseTime(p.EndTime)
}

func parseTime(i *string) *time.Time {
	if i == nil || *i == "" {
		return nil
	}
	sli := stringify.ToStringSlice(*i, " ")
	day := stringify.ToIntSlice(sli[0], "-")
	if len(day) != 3 {
		return nil
	}
	var hour int
	var min int
	var sec int
	if len(sli) > 1 {
		ti := stringify.ToIntSlice(sli[1], ":")
		if len(ti) > 0 {
			hour = int(ti[0])
		}
		if len(ti) > 1 {
			min = int(ti[1])
		}
		if len(ti) > 2 {
			sec = int(ti[2])
		}
	}
	date := time.Date(int(day[0]), time.Month(day[1]), int(day[2]), hour, min, sec, 0, time.Local)
	return &date
}
