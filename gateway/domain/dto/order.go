package dto

type OrderDto struct {
	Products []OrderProductDto `json:"products"`
}

type OrderRequestDro struct {
	UserID   uint              `json:"user_id"`
	Products []OrderProductDto `json:"products"`
}

type OrderProductDto struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
