package controller

import (
	"fehu/middleware"
	"fehu/model/body"
	"fehu/service/order"
	"fehu/util/map_builder"
	"fehu/util/request"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

func OrderRegister(router *gin.RouterGroup, needLoginedRouter *gin.RouterGroup) {
	c := OrderController{}
	router.POST("/order/add", c.add)
	router.POST("/order/list", c.getList)
}

func (*OrderController) add(c *gin.Context) {
	var params body.Order
	request.Bind(c, &params)
	order.Add(params)
	middleware.Success(c)
}

func (*OrderController) getList(c *gin.Context) {
	var params body.Order
	request.Bind(c, &params)
	list := order.GetByUid(params.UserId)
	res := order.GetByOrderId(params.OrderId)
	middleware.Success(c, map_builder.BuilderMap("list", list, "res", res))
}
