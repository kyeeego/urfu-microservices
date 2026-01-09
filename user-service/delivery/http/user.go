package http

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kyeeego/urfu-microservices/user-service/domain/dto"
)

func (h *Handler) HandleRegister(c *gin.Context) {
	var body dto.LoginDto
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, map[string]string{"error": err.Error()})
		return
	}

	err := h.services.User.Register(body.Username, body.Password)
	if err != nil {
		c.AbortWithStatusJSON(401, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]string{"username": body.Username})
}

func (h *Handler) HandleLogin(c *gin.Context) {
	var body dto.LoginDto
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, map[string]string{"error": err.Error()})
		return
	}

	token, err := h.services.Auth.Login(body.Username, body.Password)
	if err != nil {
		c.AbortWithStatusJSON(401, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]string{"token": token})
}

func (h *Handler) HandleAuthorize(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, errors.New("no auth header"))
		return
	}
	a := strings.Split(authHeader, " ")
	if a[0] != "Bearer" || len(a) < 2 {
		c.JSON(401, errors.New("invalid auth header"))
		return
	}

	token := a[1]
	username, err := h.services.Auth.Authorize(token)
	if err != nil {
		c.JSON(403, errors.New("invalid token"))
		return
	}

	c.JSON(200, map[string]string{"username": username})
}
