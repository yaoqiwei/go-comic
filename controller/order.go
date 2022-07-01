package controller

import (
	"fehu/middleware"
	"fehu/model/body"
	"fehu/service/order"
	"fehu/service/order/orderInfo"
	"fehu/util/map_builder"
	"fehu/util/request"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

func OrderRegister(router *gin.RouterGroup, needLoginedRouter *gin.RouterGroup) {
	c := OrderController{}
	router.POST("/order/add", c.add)
	router.POST("/order/addOrderInfo", c.addOrderInfo)
	router.POST("/order/list", c.getList)
	router.POST("/order/getOrderInfoList", c.getOrderInfoList)
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
	res := order.GetByOrderId(params.Id)
	middleware.Success(c, map_builder.BuilderMap("list", list, "res", res))
}

func (*OrderController) addOrderInfo(c *gin.Context) {
	var params body.OrderInfo
	request.Bind(c, &params)
	orderInfo.Add(params)
	middleware.Success(c)
}

func (*OrderController) getOrderInfoList(c *gin.Context) {
	var params body.OrderInfo
	request.Bind(c, &params)
	list := orderInfo.GetByUid(params.UserId)
	res := orderInfo.GetByOrderId(params.OrderId)
	middleware.Success(c, map_builder.BuilderMap("list", list, "res", res))
}
