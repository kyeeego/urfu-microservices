package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"name" gorm:"not null"`
	Price float64 `json:"price" gorm:"not null"`
}

type ProductDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
