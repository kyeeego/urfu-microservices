package repository

import (
	"github.com/kyeeego/urfu-microservices/user-service/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id uint) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	Insert(model *domain.User) error
}

type Repository struct {
	User UserRepository
}

func New(db *gorm.DB) *Repository {
	db.AutoMigrate(&domain.User{})
	return &Repository{
		User: newUserRepository(db),
	}
}
