package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/IaC/go-kcloutie/pkg/model"
)

// GetTime godoc
//
//	@title       Get Time
//	@version     1.0
//	@Description Get a Time object
//	@Accept      json
//	@Produce     json
//	@Success     200 {object} model.Time
//	@Failure     400 {object} model.ErrorDetail
//	@Router      /api/v1/time [get]
func (con *Controller) GetTime(ctx context.Context, c *gin.Context) {
	timeO := model.Time{
		CurrentTime: time.Now(),
	}
	c.JSON(http.StatusOK, timeO)
}

// CreateTime godoc
//
//	@title       Post Time
//	@version     1.0
//	@Description Send a Time object and get it back with the current time and a custom message or error
//	@Accept      json
//	@Produce     json
//	@Param			 body		body		model.Time	true	"The new Time to create"
//	@Success     200 {object} model.Time
//	@Failure     400 {object} model.ErrorDetail
//	@Router      /api/v1/time [post]
func (con *Controller) PostTime(ctx context.Context, c *gin.Context) {
	var Time model.Time

	err := c.ShouldBindJSON(&Time)
	if err != nil {
		errD := &model.ErrorDetail{
			Type:     "unmarshal-time",
			Title:    "Unmarshal Time",
			Status:   400,
			Detail:   fmt.Sprintf("failed to unmarshal the json body to a Time - %v", err),
			Instance: "POST - /api/v1/time",
		}
		c.JSON(int(errD.Status), errD)
		return
	}
	if Time.ThrowError {
		errD := &model.ErrorDetail{
			Type:     "throw-error",
			Title:    "Throw Error",
			Status:   400,
			Detail:   "ThrowError was set to true...throwing an  error",
			Instance: "POST - /api/v1/time",
		}
		c.JSON(int(errD.Status), errD)
		return
	}

	Time.CurrentTime = time.Now()
	if Time.Name == "" {
		Time.Name = "Unknown"
	}
	Time.Message = fmt.Sprintf("Hello %v, the current time is %v", Time.Name, Time.CurrentTime)

	c.JSON(http.StatusOK, Time)
}
