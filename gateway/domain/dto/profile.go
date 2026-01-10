package dto

type ProfileResponseDto struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Orders   []Order `json:"orders"`
}

type Order struct {
	ID         uint           `json:"id"`
	TotalPrice float64        `json:"total_price"`
	Products   []OrderProduct `json:"products"`
}

type OrderProduct struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ProductClientResponse struct {
	ID    uint
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UserClientResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type OrderClientResponse struct {
	ID       uint                         `json:"id"`
	Products []OrderProductClientResponse `json:"products"`
}

type OrderProductClientResponse struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
