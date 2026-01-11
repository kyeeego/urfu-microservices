package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kyeeego/urfu-microservices/product-service/service"
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
		a.GET("/all", h.HandleGet)
		a.GET("/id/:id", h.HandleGetById)
		a.POST("/", h.HandleInsert)
	}
}
