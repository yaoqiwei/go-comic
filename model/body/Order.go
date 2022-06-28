package body

type Order struct {
	Id          int64  ` json:"id"`
	OrderId     int64  ` json:"orderId"`
	UserId      int64  `json:"userId"`
	ProductName string ` json:"productName"`
}
