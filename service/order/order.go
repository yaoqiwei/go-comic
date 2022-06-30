package order

import (
	"fehu/common/lib/gorm"
	"fehu/model/body"
)

func Add(p body.Order) error {
	err := gorm.Db.Create(&p).Error
	return err
}

func GetByUid(uid int64) []body.Order {
	order := make([]body.Order, 0)
	gorm.Db.Where("user_id = ?", uid).Find(&order)
	return order
}
func GetByOrderId(orderId int64) []body.Order {
	order := make([]body.Order, 0)
	gorm.Db.Where("id = ?", orderId).Find(&order)
	return order
}
