package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Livez godoc
//
//	@Summary     API liveness
//	@Description API liveness response
//	@Tags        liveness
//	@Accept      json
//	@Produce     json
//	@Success     200  {object} LivezResponse
//	@Router      /livez [get]
func Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, LivezResponse{
		Status: "ok",
	})
}

type LivezResponse struct {
	Status string `json:"status" yaml:"status"`
}
