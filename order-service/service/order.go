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

func (s OrderServiceImpl) GetByUserId(id uint) ([]domain.Order, error) {
	return s.repository.Order.GetByUserId(id)
}

func (s OrderServiceImpl) Get() ([]domain.Order, error) {
	return s.repository.Order.Get()
}

func (s OrderServiceImpl) Insert(body domain.OrderDto) error {
	model := &domain.Order{
		UserID: body.UserID,
	}
	err := s.repository.Order.Insert(model)
	if err != nil {
		return err
	}

	for _, pr := range body.Products {
		op := &domain.OrderProducts{
			OrderID:   model.ID,
			ProductID: pr.ProductID,
			Quantity:  pr.Quantity,
		}
		err = s.repository.Order.InsertOrderProducts(op)
		if err != nil {
			return err
		}
	}

	return nil
}
