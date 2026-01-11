package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID   uint            `json:"user_id"`
	Products []OrderProducts `json:"products" gorm:"foreignKey:OrderID"`
}

type OrderProducts struct {
	gorm.Model
	OrderID   uint `json:"order_id" gorm:"not null"`
	ProductID uint `json:"product_id" gorm:"not null"`
	Quantity  int  `json:"quantity" gorm:"not null"`
}

type OrderDto struct {
	UserID   uint               `json:"user_id"`
	Products []OrderProductsDto `json:"products"`
}

type OrderProductsDto struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
