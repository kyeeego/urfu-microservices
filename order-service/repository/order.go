package repository

import (
	"github.com/kyeeego/urfu-microservices/order-service/domain"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func newOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r orderRepository) Get() ([]domain.Order, error) {
	var orders []domain.Order
	res := r.db.Preload("Products").Find(&orders)

	return orders, res.Error
}

func (r orderRepository) GetByUserId(id uint) ([]domain.Order, error) {
	var orders []domain.Order
	res := r.db.Preload("Products").Find(&orders, domain.Order{UserID: id})

	return orders, res.Error
}

func (r orderRepository) GetById(id uint) (domain.Order, error) {
	order := domain.Order{}
	res := r.db.Preload("Products").First(&order, id)

	return order, res.Error
}

func (r orderRepository) Insert(model *domain.Order) error {
	return r.db.Create(model).Error
}

func (r orderRepository) InsertOrderProducts(products *domain.OrderProducts) error {
	return r.db.Create(products).Error
}
