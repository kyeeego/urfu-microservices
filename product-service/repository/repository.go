package repository

import (
	"github.com/kyeeego/urfu-microservices/product-service/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Insert(*domain.Product) error
	Get() ([]domain.Product, error)
	GetById(id uint) (domain.Product, error)
}

type Repository struct {
	Product ProductRepository
}

func New(db *gorm.DB) *Repository {
	db.AutoMigrate(&domain.Product{})
	return &Repository{
		Product: newProductRepository(db),
	}
}
