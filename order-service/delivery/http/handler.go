package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kyeeego/urfu-microservices/order-service/service"
)

type Handler struct {
	services *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	a := router.Group("/")
	{
		a.GET("/", h.HandleGet)
		a.GET("/:id", h.HandleGetById)
		a.GET("/user/:id", h.HandleGetByUserId)
	}
}
