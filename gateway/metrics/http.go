package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpRequestsAmount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_amount",
			Help: "Total number of HTTP requests by service, method, endpoint, and status",
		},
		[]string{"service", "method", "endpoint", "status"},
	)
)
