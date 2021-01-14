package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/year"
)

// YearController is ...
type YearController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateYear handles POST requests for adding year entities
// @Summary Create year
// @Description Create year
// @ID create-year
// @Accept   json
// @Produce  json
// @Param year body ent.Year true "Year entity"
// @Success 200 {object} ent.Year
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /years [post]
func (ctl *YearController) CreateYear(c *gin.Context) {
	obj := ent.Year{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "year binding failed",
		})
		return
	}

	y, err := ctl.client.Year.
		Create().
		SetYear(obj.Year).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, y)
}

// GetYear handles GET requests to retrieve a year entity
// @Summary Get a year entity by ID
// @Description get year by ID
// @ID get-year
// @Produce  json
// @Param id path int true "Year ID"
// @Success 200 {object} ent.Year
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /years/{id} [get]
func (ctl *YearController) GetYear(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	y, err := ctl.client.Year.
		Query().
		Where(year.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, y)
}

// ListYear handles request to get a list of year entities
// @Summary List year entities
// @Description list year entities
// @ID list-year
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Year
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /years [get]
func (ctl *YearController) ListYear(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	years, err := ctl.client.Year.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, years)
}

// DeleteYear handles DELETE requests to delete a year entity
// @Summary Delete a year entity by ID
// @Description get year by ID
// @ID delete-year
// @Produce  json
// @Param id path int true "Year ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /years/{id} [delete]
func (ctl *YearController) DeleteYear(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Year.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateYear handles PUT requests to update a year entity
// @Summary Update a year entity by ID
// @Description update year by ID
// @ID update-year
// @Accept   json
// @Produce  json
// @Param id path int true "Year ID"
// @Param year body ent.Year true "Year entity"
// @Success 200 {object} ent.Year
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /years/{id} [put]
func (ctl *YearController) UpdateYear(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Year{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "year binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	yr, err := ctl.client.Year.
		UpdateOneID(int(id)).
		SetYear(obj.Year).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, yr)
}

// NewYearController creates and registers handles for the year controller
func NewYearController(router gin.IRouter, client *ent.Client) *YearController {
	yrc := &YearController{
		client: client,
		router: router,
	}

	yrc.register()
	return yrc

}

func (ctl *YearController) register() {
	years := ctl.router.Group("/years")

	years.GET("", ctl.ListYear)

	//CRUD
	years.POST("", ctl.CreateYear)
	years.GET(":id", ctl.GetYear)
	years.PUT(":id", ctl.UpdateYear)
	years.DELETE(":id", ctl.DeleteYear)
}
