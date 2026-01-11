package repository

import (
	"github.com/kyeeego/urfu-microservices/product-service/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func newProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r productRepository) Get() ([]domain.Product, error) {
	var products []domain.Product
	res := r.db.Find(&products)

	return products, res.Error
}

func (r productRepository) GetById(id uint) (domain.Product, error) {
	product := domain.Product{}
	res := r.db.First(&product, id)

	return product, res.Error
}

func (r productRepository) Insert(model *domain.Product) error {
	return r.db.Create(model).Error
}
