package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
//
//	@Summary     API Health
//	@Description API Health response
//	@Tags        health
//	@Accept      json
//	@Produce     json
//	@Success     200  {object} HealthResponse
//	@Router      /healthz [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Healthy: true,
	})
}

type HealthResponse struct {
	Healthy bool `json:"healthy" yaml:"healthy"`
}
