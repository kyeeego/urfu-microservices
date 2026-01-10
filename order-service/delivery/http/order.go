package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleGetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]error{"error": err})
		return
	}

	model, err := h.services.Order.GetById(uint(id))
	if err != nil {
		c.JSON(404, map[string]error{"error": err})
		return
	}

	c.JSON(200, model)
}

func (h *Handler) HandleGet(c *gin.Context) {
	models, err := h.services.Order.Get()
	if err != nil {
		c.JSON(404, map[string]error{"error": err})
		return
	}

	c.JSON(200, models)
}

func (h *Handler) HandleGetByUserId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, map[string]error{"error": err})
		return
	}

	model, err := h.services.Order.GetByUserId(uint(id))
	if err != nil {
		c.JSON(404, map[string]error{"error": err})
		return
	}

	c.JSON(200, model)
}
