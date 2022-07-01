package orderInfo

import (
	"fehu/common/lib/gorm"
	"fehu/model/body"
)

func Add(p body.OrderInfo) error {
	p.Id = gorm.Snowflake.GetId()
	err := gorm.ShardingTable("comic_order_info", p.UserId).Create(&p).Error
	return err
}

func GetByUid(uid int64) []body.OrderInfo {
	order := make([]body.OrderInfo, 0)
	gorm.ShardingTable("comic_order_info", uid).Where("user_id = ?", uid).Find(&order)
	return order
}
func GetByOrderId(orderId int64) []body.OrderInfo {
	order := make([]body.OrderInfo, 0)
	gorm.ShardingTable("comic_order_info", orderId).Where("order_id = ?", orderId).Find(&order)
	return order
}
