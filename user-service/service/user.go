package service

import (
	"errors"

	"github.com/kyeeego/urfu-microservices/user-service/domain"
	"github.com/kyeeego/urfu-microservices/user-service/pkg/password"
	"github.com/kyeeego/urfu-microservices/user-service/repository"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	repository *repository.Repository
}

func newUserService(r *repository.Repository) UserService {
	return &UserServiceImpl{r}
}

func (s UserServiceImpl) Register(username string, pass string) error {
	_, err := s.repository.User.GetByUsername(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else if err == nil {
		return errors.New("user already exists")
	}

	model := domain.User{Username: username}

	model.Password, err = password.Hash(pass)
	if err != nil {
		return err
	}

	return s.repository.User.Insert(&model)
}
