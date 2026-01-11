package metrics

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GinMetricsMiddleware(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/metrics" {
			c.Next()
			return
		}

		c.Next()

		HttpRequestsAmount.WithLabelValues(
			serviceName,
			c.Request.Method,
			c.FullPath(),
			strconv.Itoa(c.Writer.Status()),
		).Inc()
	}
}
