package service

import (
	"github.com/kyeeego/urfu-microservices/product-service/domain"
	"github.com/kyeeego/urfu-microservices/product-service/repository"
)

type ProductServiceImpl struct {
	repository *repository.Repository
}

func newProductService(r *repository.Repository) ProductService {
	return &ProductServiceImpl{r}
}

func (s ProductServiceImpl) GetById(id uint) (domain.Product, error) {
	return s.repository.Product.GetById(id)
}

func (s ProductServiceImpl) Get() ([]domain.Product, error) {
	return s.repository.Product.Get()
}
