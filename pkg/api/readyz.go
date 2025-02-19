package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Readyz godoc
//
//	@Summary     API readiness
//	@Description API readiness response
//	@Tags        readiness
//	@Accept      json
//	@Produce     json
//	@Success     200  {object} ReadyzResponse
//	@Router      /readyz [get]
func Readyz(c *gin.Context) {
	c.JSON(http.StatusOK, ReadyzResponse{
		Ready: true,
	})
}

type ReadyzResponse struct {
	Ready bool `json:"ready" yaml:"ready"`
}
