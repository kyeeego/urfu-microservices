package http

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/kyeeego/urfu-microservices/gateway/config"
	"github.com/kyeeego/urfu-microservices/gateway/delivery/http/clients"
)

type Handler struct {
	cfg   *config.Config
	http  clients.HttpClientWithRetry
	redis *redis.Client
}

func New(cfg *config.Config, http clients.HttpClientWithRetry, redis *redis.Client) *Handler {
	return &Handler{cfg, http, redis}
}

func (h *Handler) Init(logger *slog.Logger, rateLimit float64, rateLimitBurst int) *gin.Engine {
	router := gin.Default()

	router.Use(SlogLogger(logger))
	router.Use(RateLimit(rateLimit, rateLimitBurst))

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
