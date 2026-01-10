package http

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func (h *Handler) JwtAuthorize(c *gin.Context) {
	status, _, err := h.http.Get(h.cfg.UsersUrl+"/authorize", map[string]string{
		"Authorization": c.GetHeader("Authorization"),
	})

	if err != nil || status != 200 {
		slog.Error("request attempt with invalid jwt")
		c.AbortWithStatusJSON(403, map[string]string{
			"error": "invalid JWT",
		})
	}
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
