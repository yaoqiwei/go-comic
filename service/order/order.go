package order

import (
	"fehu/common/lib"
	"fehu/model/body"
)

func Add(p body.Order) error {
	err := lib.Db.Create(&p).Error
	return err
}

func GetByUid(uid int64) []body.Order {
	order := make([]body.Order, 0)
	lib.Db.Where("user_id = ?", uid).Find(&order)
	return order
}
func GetByOrderId(orderId int64) []body.Order {
	order := make([]body.Order, 0)
	lib.Db.Where("order_id = ?", orderId).Find(&order)
	return order
}
