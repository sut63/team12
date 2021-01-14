package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/discipline"
	"github.com/gin-gonic/gin"
)

// DisciplineController is ...
type DisciplineController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDiscipline handles POST requests for adding discipline entities
// @Summary Create discipline
// @Description Create discipline
// @ID create-discipline
// @Accept   json
// @Produce  json
// @Param discipline body ent.Discipline true "Discipline entity"
// @Success 200 {object} ent.Discipline
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /disciplines [post]
func (ctl *DisciplineController) CreateDiscipline(c *gin.Context) {
	obj := ent.Discipline{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "discipline binding failed",
		})
		return
	}

	d, err := ctl.client.Discipline.
		Create().
		SetDiscipline(obj.Discipline).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetDiscipline handles GET requests to retrieve a discipline entity
// @Summary Get a discipline entity by ID
// @Description get discipline by ID
// @ID get-discipline
// @Produce  json
// @Param id path int true "Discipline ID"
// @Success 200 {object} ent.Discipline
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /disciplines/{id} [get]
func (ctl *DisciplineController) GetDiscipline(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	d, err := ctl.client.Discipline.
		Query().
		Where(discipline.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, d)
}

// ListDiscipline handles request to get a list of discipline entities
// @Summary List discipline entities
// @Description list discipline entities
// @ID list-discipline
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Discipline
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /disciplines [get]
func (ctl *DisciplineController) ListDiscipline(c *gin.Context) {
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

	disciplines, err := ctl.client.Discipline.
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

	c.JSON(200, disciplines)
}

// DeleteDiscipline handles DELETE requests to delete a discipline entity
// @Summary Delete a discipline entity by ID
// @Description get discipline by ID
// @ID delete-discipline
// @Produce  json
// @Param id path int true "Discipline ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /disciplines/{id} [delete]
func (ctl *DisciplineController) DeleteDiscipline(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Discipline.
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

// UpdateDiscipline handles PUT requests to update a discipline entity
// @Summary Update a discipline entity by ID
// @Description update discipline by ID
// @ID update-discipline
// @Accept   json
// @Produce  json
// @Param id path int true "Discipline ID"
// @Param discipline body ent.Discipline true "Discipline entity"
// @Success 200 {object} ent.Discipline
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /disciplines/{id} [put]
func (ctl *DisciplineController) UpdateDiscipline(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Discipline{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "discipline binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	dr, err := ctl.client.Discipline.
		UpdateOneID(int(id)).
		SetDiscipline(obj.Discipline).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, dr)
}

// NewDisciplineController creates and registers handles for the discipline controller
func NewDisciplineController(router gin.IRouter, client *ent.Client) *DisciplineController {
	drc := &DisciplineController{
		client: client,
		router: router,
	}

	drc.register()
	return drc

}

func (ctl *DisciplineController) register() {
	disciplines := ctl.router.Group("/disciplines")

	disciplines.GET("", ctl.ListDiscipline)

	//CRUD
	disciplines.POST("", ctl.CreateDiscipline)
	disciplines.GET(":id", ctl.GetDiscipline)
	disciplines.PUT(":id", ctl.UpdateDiscipline)
	disciplines.DELETE(":id", ctl.DeleteDiscipline)
}
