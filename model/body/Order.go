package body

type Order struct {
	Id          int64  ` json:"id"`
	UserId      int64  `json:"userId"`
	ProductName string ` json:"productName"`
}

type OrderInfo struct {
	Id      int64  ` json:"id"`
	UserId  int64  `json:"userId"`
	OrderId int64  `json:"orderId"`
	Info    string ` json:"info"`
}
