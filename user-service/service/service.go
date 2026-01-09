package service

import (
	"fmt"
	"os"

	"github.com/kyeeego/urfu-microservices/user-service/pkg/jwt"
	"github.com/kyeeego/urfu-microservices/user-service/repository"
)

type UserService interface {
	Register(username string, password string) error
}

type AuthService interface {
	Login(username, password string) (string, error)
	Authorize(jwt string) (string, error)
}

type Service struct {
	User UserService
	Auth AuthService
}

func New(repository *repository.Repository) *Service {
	m, err := jwt.NewManager(os.Getenv("JWT_KEY"))
	if err != nil {
		panic(fmt.Sprintf("couldnt initialize services: %e\n", err))
	}

	return &Service{
		User: newUserService(repository),
		Auth: newAuthService(m, repository),
	}
}
