package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestPrometheusMiddleware(t *testing.T) {

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create a new Gin engine
	r := gin.Default()

	// Use the Prometheus middleware
	r.Use(PrometheusMiddleware())

	// Create a test endpoint
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	// Make a request to the test endpoint
	r.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the Prometheus metrics
	metrics, _ := prometheus.DefaultGatherer.Gather()
	for _, m := range metrics {
		if *m.Name == "go-kcloutie_http_requests_total" {
			for _, metric := range m.Metric {
				if *metric.Label[0].Value == "/test" && *metric.Label[1].Value == "200" {
					assert.Equal(t, float64(1), *metric.Counter.Value)
				}
			}
		}
	}
}
