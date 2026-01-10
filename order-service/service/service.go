package service

import (
	"github.com/kyeeego/urfu-microservices/order-service/domain"
	"github.com/kyeeego/urfu-microservices/order-service/repository"
)

type OrderService interface {
	Get() ([]domain.Order, error)
	GetById(id uint) (domain.Order, error)
	GetByUserId(id uint) ([]domain.Order, error)
}

type Service struct {
	Order OrderService
}

func New(repository *repository.Repository) *Service {
	return &Service{
		Order: newOrderService(repository),
	}
}
