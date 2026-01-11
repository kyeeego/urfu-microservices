package service

import (
	"github.com/kyeeego/urfu-microservices/product-service/domain"
	"github.com/kyeeego/urfu-microservices/product-service/repository"
)

type ProductService interface {
	Insert(domain.ProductDto) error
	Get() ([]domain.Product, error)
	GetById(id uint) (domain.Product, error)
}

type Service struct {
	Product ProductService
}

func New(repository *repository.Repository) *Service {
	return &Service{
		Product: newProductService(repository),
	}
}
