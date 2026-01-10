package http

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/kyeeego/urfu-microservices/gateway/domain/dto"
)

func (h *Handler) HandleRegister(c *gin.Context) {
	_, res, err := h.http.Post(h.cfg.UsersUrl+"/signup", map[string]string{}, c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(401, map[string]string{"error": err.Error()})
		return
	}

	var body dto.RegisterLoginDto
	err = json.Unmarshal(res, &body)
	if err != nil {
		c.AbortWithStatusJSON(500, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(200, body)
}

func (h *Handler) HandleLogin(c *gin.Context) {
	_, res, err := h.http.Post(h.cfg.UsersUrl+"/login", map[string]string{}, c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(401, map[string]string{"error": err.Error()})
		return
	}

	var body dto.TokenDto
	err = json.Unmarshal(res, &body)
	if err != nil {
		c.AbortWithStatusJSON(500, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(200, body)
}

func (h *Handler) HandleAggregateProfile(c *gin.Context) {
	id := c.Param("id")

	cached, err := h.redis.Get(id).Result()
	if err == nil {
		fmt.Println("Using cached result")
		c.JSON(200, cached)
		return
	} else if err != redis.Nil {
		fmt.Printf("redis error. proceeding without it: %e\n", err)
	}

	status, res, err := h.http.Get(fmt.Sprintf("%s/%s", h.cfg.UsersUrl, id), map[string]string{})
	if err != nil {
		c.AbortWithStatusJSON(status, map[string]string{"error": err.Error()})
		return
	}

	var user dto.UserClientResponse
	err = json.Unmarshal(res, &user)
	if err != nil {
		c.AbortWithStatusJSON(500, map[string]string{"error": err.Error()})
		return
	}

	status, res, err = h.http.Get(fmt.Sprintf("%s/user/%s", h.cfg.OrdersUrl, id), map[string]string{})
	if err != nil {
		c.AbortWithStatusJSON(status, map[string]string{"error": err.Error()})
		return
	}

	var orders []dto.OrderClientResponse
	err = json.Unmarshal(res, &orders)
	if err != nil {
		c.AbortWithStatusJSON(500, map[string]string{"error": err.Error()})
		return
	}

	result := dto.ProfileResponseDto{
		ID:       user.ID,
		Username: user.Username,
		Orders:   []dto.Order{},
	}

	for _, order := range orders {
		res := dto.Order{
			ID:         order.ID,
			TotalPrice: 0,
			Products:   []dto.OrderProduct{},
		}

		for _, product := range order.Products {
			_, resp, err := h.http.Get(fmt.Sprintf("%s/%d", h.cfg.ProductsUrl, product.ProductID), map[string]string{})
			if err != nil {
				continue
			}

			var prod dto.ProductClientResponse
			err = json.Unmarshal(resp, &prod)
			if err != nil {
				continue
			}

			res.Products = append(res.Products, dto.OrderProduct{
				Name:     prod.Name,
				Price:    prod.Price,
				Quantity: product.Quantity,
			})

			res.TotalPrice += prod.Price * float64(product.Quantity)
		}

		result.Orders = append(result.Orders, res)
	}

	strRes, _ := json.Marshal(result)
	err = h.redis.Set(id, strRes, time.Second*time.Duration(h.cfg.RedisTtl)).Err()
	if err != nil {
		fmt.Printf("redis error: %e\n", err)
	}

	c.JSON(200, result)
}
