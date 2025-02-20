package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/IaC/go-kcloutie/pkg/model"
)

// GetAllWidgets godoc
//
//	@title       Get all widgets
//	@version     1.0
//	@Description Get all widgets within the database
//	@Accept      json
//	@Produce     json
//	@Success     200 {object} []model.Widget
//	@Failure     400 {object} model.ErrorDetail
//	@Router      /api/v1/widgets [get]
func (con *Controller) GetAllWidgets(ctx context.Context, c *gin.Context) {
	var widgets []model.Widget
	DB := con.DBInterface.GetDB()

	if result := DB.Find(&widgets); result.Error != nil {
		errD := &model.ErrorDetail{
			Type:     "get-all-widgets",
			Title:    "Get all widgets",
			Status:   400,
			Detail:   fmt.Sprintf("failed to get all widgets - %v", result.Error),
			Instance: "GET - /api/v1/widgets",
		}
		c.JSON(int(errD.Status), errD)
		return
	}
	c.JSON(http.StatusOK, widgets)
}

// GetWidget godoc
//
//	@title       Get widget
//	@version     1.0
//	@Description Get a widget by its ID
//	@Accept      json
//	@Produce     json
//	@Param       id		path		int		true	"The id of the widget to get"
//	@Success     200 {object} model.Widget
//	@Failure     400 {object} model.ErrorDetail
//	@Router      /api/v1/widgets/{id} [get]
func (con *Controller) GetWidget(ctx context.Context, c *gin.Context) {
	var widget model.Widget
	DB := con.DBInterface.GetDB()

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if result := DB.First(&widget, id); result.Error != nil {
		errD := &model.ErrorDetail{
			Type:     "get-widget",
			Title:    "Get widget",
			Status:   400,
			Detail:   fmt.Sprintf("failed to find a widget with ID %v - %v", id, result.Error),
			Instance: "GET - /api/v1/widgets/{id}",
		}
		c.JSON(int(errD.Status), errD)
		return
	}
	c.JSON(http.StatusOK, widget)
}

// CreateWidget godoc
//
//	@title       Create widget
//	@version     1.0
//	@Description Creates a new widget
//	@Accept      json
//	@Produce     json
//	@Param			 body		body		model.Widget	true	"The new widget to create"
//	@Success     201
//	@Failure     400 {object} model.ErrorDetail
//	@Router      /api/v1/widgets [post]
func (con *Controller) CreateWidget(ctx context.Context, c *gin.Context) {
	var widget model.Widget
	DB := con.DBInterface.GetDB()

	err := c.ShouldBindJSON(&widget)
	if err != nil {
		errD := &model.ErrorDetail{
			Type:     "unmarshal-widget",
			Title:    "Unmarshal widget",
			Status:   400,
			Detail:   fmt.Sprintf("failed to unmarshal the json body to a widget - %v", err),
			Instance: "POST - /api/v1/widgets",
		}
		c.JSON(int(errD.Status), errD)
		return
	}

	if result := DB.Create(&widget); result.Error != nil {
		errD := &model.ErrorDetail{
			Type:     "create-widget",
			Title:    "Create widget",
			Status:   400,
			Detail:   fmt.Sprintf("failed to create a new widget - %v", result.Error),
			Instance: "POST - /api/v1/widgets",
		}
		c.JSON(int(errD.Status), errD)
		return
	}

	c.Redirect(http.StatusCreated, fmt.Sprintf("/api/v1/widgets/%v", widget.ID))
}

// UpdateWidget godoc
//
//	@title       Update widget
//	@version     1.0
//	@Description Updates an existing widget
//	@Accept      json
//	@Produce     json
//	@Param			 body		body		model.Widget	true	"The new widget to update"
//	@Param       id		path		int		true	"The id of the widget to get"
//	@Success     200 {object} model.Widget
//	@Failure     400 {object} model.ErrorDetail
//	@Router      /api/v1/widgets/{id} [put]
func (con *Controller) UpdateWidget(ctx context.Context, c *gin.Context) {
	var updatedWidget model.Widget
	var existingWidget model.Widget
	DB := con.DBInterface.GetDB()

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := c.ShouldBindJSON(&updatedWidget)
	if err != nil {
		errD := &model.ErrorDetail{
			Type:     "unmarshal-widget",
			Title:    "Unmarshal widget",
			Status:   400,
			Detail:   fmt.Sprintf("failed to unmarshal the json body to a widget - %v", err),
			Instance: "PUT - /api/v1/widgets",
		}
		c.JSON(int(errD.Status), errD)
		return
	}

	if result := DB.First(&existingWidget, id); result.Error != nil {
		errD := &model.ErrorDetail{
			Type:     "get-widget",
			Title:    "Get widget",
			Status:   400,
			Detail:   fmt.Sprintf("failed to find a widget with ID %v - %v", id, result.Error),
			Instance: "PUT - /api/v1/widgets/{id}",
		}
		c.JSON(int(errD.Status), errD)
		return
	}

	existingWidget.Count = updatedWidget.Count
	existingWidget.Name = updatedWidget.Name
	existingWidget.Description = updatedWidget.Description
	existingWidget.Updater = "unknown"

	result := DB.Save(&updatedWidget)

	if result.Error != nil {
		errD := &model.ErrorDetail{
			Type:     "update-widget",
			Title:    "Update widget",
			Status:   400,
			Detail:   fmt.Sprintf("failed to update widget '%v' - %v", id, result.Error),
			Instance: "PUT - /api/v1/widgets",
		}
		c.JSON(int(errD.Status), errD)
		return
	}
	c.Redirect(http.StatusCreated, fmt.Sprintf("/api/v1/widgets/%v", updatedWidget.ID))
}

// DeleteWidget godoc
//
//	@Description Delete an existing widget
//	@Accept      json
//	@Produce     json
//	@Param       id		path		int		true	"The id of the widget to get"
//	@Success     200 {object} model.Widget
//	@Failure     400 {object} model.ErrorDetail
//	@Router      /api/v1/widgets/{id} [delete]
func (con *Controller) DeleteWidget(ctx context.Context, c *gin.Context) {

	var existingWidget model.Widget
	DB := con.DBInterface.GetDB()

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if result := DB.First(&existingWidget, id); result.Error != nil {
		errD := &model.ErrorDetail{
			Type:     "get-widget",
			Title:    "Get widget",
			Status:   400,
			Detail:   fmt.Sprintf("failed to find a widget with ID %v - %v", id, result.Error),
			Instance: "GET - /api/v1/widgets/{id}",
		}
		c.JSON(int(errD.Status), errD)
		return
	}

	if result := DB.Delete(&existingWidget); result.Error != nil {
		errD := &model.ErrorDetail{
			Type:     "delete-widget",
			Title:    "Delete widget",
			Status:   400,
			Detail:   fmt.Sprintf("failed to delete widget '%v' - %v", id, result.Error),
			Instance: "DELETE - /api/v1/widgets",
		}
		c.JSON(int(errD.Status), errD)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
