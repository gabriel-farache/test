package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gokcloutie_http_requests_total",
			Help: "Number of HTTP requests",
		},
		[]string{"path", "status_code"},
	)
)

func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	http.Handle("/metrics", MetricsHandler())
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		statusCode := strconv.Itoa(c.Writer.Status())
		httpRequestsTotal.With(prometheus.Labels{"path": c.Request.URL.Path, "status_code": statusCode}).Inc()
		c.Next()
	}
}
