package repository

import (
	"github.com/kyeeego/urfu-microservices/order-service/domain"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Get() ([]domain.Order, error)
	GetById(id uint) (domain.Order, error)
	GetByUserId(id uint) ([]domain.Order, error)
	Insert(*domain.Order) error
	InsertOrderProducts(*domain.OrderProducts) error
}

type Repository struct {
	Order OrderRepository
}

func New(db *gorm.DB) *Repository {
	db.AutoMigrate(&domain.Order{}, &domain.OrderProducts{})
	return &Repository{
		Order: newOrderRepository(db),
	}
}
