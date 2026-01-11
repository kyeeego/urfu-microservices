package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kyeeego/urfu-microservices/product-service/domain"
)

func (h *Handler) HandleGetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]error{"error": err})
		return
	}

	model, err := h.services.Product.GetById(uint(id))
	if err != nil {
		c.JSON(404, map[string]error{"error": err})
		return
	}

	c.JSON(200, model)
}

func (h *Handler) HandleGet(c *gin.Context) {
	models, err := h.services.Product.Get()
	if err != nil {
		c.JSON(404, map[string]error{"error": err})
		return
	}

	c.JSON(200, models)
}

func (h *Handler) HandleInsert(c *gin.Context) {
	var body domain.ProductDto
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, map[string]string{"error": err.Error()})
		return
	}

	err := h.services.Product.Insert(body)
	if err != nil {
		c.AbortWithStatusJSON(401, map[string]string{"error": err.Error()})
		return
	}
}
