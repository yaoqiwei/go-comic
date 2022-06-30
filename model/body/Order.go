package body

type Order struct {
	Id          int64  ` json:"id"`
	UserId      int64  `json:"userId"`
	ProductName string ` json:"productName"`
}
