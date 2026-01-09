package service

import (
	"errors"
	"time"

	"github.com/kyeeego/urfu-microservices/user-service/pkg/jwt"
	"github.com/kyeeego/urfu-microservices/user-service/pkg/password"
	"github.com/kyeeego/urfu-microservices/user-service/repository"
)

type AuthServiceImpl struct {
	tokenManager jwt.TokenManager
	repository   *repository.Repository
}

func newAuthService(manager jwt.TokenManager, user *repository.Repository) *AuthServiceImpl {
	return &AuthServiceImpl{
		tokenManager: manager,
		repository:   user,
	}
}

func (s *AuthServiceImpl) Login(username, pass string) (string, error) {
	user, err := s.repository.User.GetByUsername(username)
	if err != nil {
		return "", err
	}

	if !password.Verify(pass, user.Password) {
		return "", errors.New("invalid password")
	}

	return s.tokenManager.Sign(username, time.Hour*24)
}

func (s *AuthServiceImpl) Authorize(jwt string) (string, error) {
	return s.tokenManager.Verify(jwt)
}
