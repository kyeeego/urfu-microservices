package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	Sign(userId int, ttl time.Duration) (string, error)
	Verify(jwt string) (int, error)
}

type Manager struct {
	key string
}

func NewManager(key string) (*Manager, error) {
	if key == "" {
		return nil, errors.New("empty key")
	}

	return &Manager{key: key}, nil
}

func (m *Manager) Sign(userId int, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(ttl).Unix(),
		"sub": userId,
	})

	return token.SignedString([]byte(m.key))
}

func (m *Manager) Verify(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(m.key), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Couldn't get user from token")
	}

	return int(claims["sub"].(float64)), err
}
