package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kyeeego/urfu-microservices/gateway/config"
	"github.com/kyeeego/urfu-microservices/gateway/delivery/http/clients"
)

type Handler struct {
	cfg  *config.Config
	http clients.HttpClientWithRetry
}

func New(cfg *config.Config, http clients.HttpClientWithRetry) *Handler {
	return &Handler{cfg, http}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	a := router.Group("/api")
	{
		a.POST("/signup", h.HandleRegister)
		a.POST("/login", h.HandleLogin)
		a.GET("/profile/:id", h.JwtAuthorize, h.HandleAggregateProfile)
	}
}
