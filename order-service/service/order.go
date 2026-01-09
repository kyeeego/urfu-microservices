package service

import (
	"github.com/kyeeego/urfu-microservices/order-service/domain"
	"github.com/kyeeego/urfu-microservices/order-service/repository"
)

type OrderServiceImpl struct {
	repository *repository.Repository
}

func newOrderService(r *repository.Repository) OrderService {
	return &OrderServiceImpl{r}
}

func (s OrderServiceImpl) GetById(id uint) (domain.Order, error) {
	return s.repository.Order.GetById(id)
}

func (s OrderServiceImpl) Get() ([]domain.Order, error) {
	return s.repository.Order.Get()
}
