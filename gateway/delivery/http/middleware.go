package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kyeeego/urfu-microservices/gateway/domain/dto"
	"golang.org/x/time/rate"
)

func (h *Handler) JwtAuthorize(c *gin.Context) {
	status, body, err := h.http.Get(h.cfg.UsersUrl+"/authorize", map[string]string{
		"Authorization": c.GetHeader("Authorization"),
	})
	if err != nil || status != 200 {
		slog.Error("request attempt with invalid jwt")
		c.AbortWithStatusJSON(403, map[string]string{
			"error": "invalid JWT",
		})
		return
	}

	var authResponse dto.AuthorizeDto
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		slog.Error("unable to read authorization response")
		c.AbortWithStatusJSON(500, map[string]string{
			"error": err.Error(),
		})
	}
	c.Set("user_id", authResponse.UserID)
}

func SlogLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		level := slog.LevelInfo
		if statusCode >= 500 {
			level = slog.LevelError
		} else if statusCode >= 400 {
			level = slog.LevelWarn
		}

		logger.Log(c.Request.Context(), level, "HTTP Request",
			slog.Int("status", statusCode),
			slog.String("method", method),
			slog.String("path", path),
			slog.String("ip", clientIP),
			slog.Duration("latency", latency),
			slog.String("error", errorMessage),
		)
	}
}

func RateLimit(limit float64, burst int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(limit), burst)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			slog.Warn("too many requests: ", c.ClientIP())
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			return
		}
		c.Next()
	}
}
