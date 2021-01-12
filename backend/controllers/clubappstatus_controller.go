package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/clubappstatus"
	"github.com/gin-gonic/gin"
)

// ClubappStatusController defines the struct for the clubappstatus controller
type ClubappStatusController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateClubappStatus handles POST requests for adding clubappstatus entities
// @Summary Create clubappstatus
// @Description Create clubappstatus
// @ID create-clubappstatus
// @Accept   json
// @Produce  json
// @Param clubappstatus body ent.ClubappStatus true "ClubappStatus entity"
// @Success 200 {object} ent.ClubappStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubappstatuss [post]
func (ctl *ClubappStatusController) CreateClubappStatus(c *gin.Context) {
	obj := ent.ClubappStatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "clubappstatus binding failed",
		})
		return
	}

	cbs, err := ctl.client.ClubappStatus.
		Create().
		SetClubstatus(obj.Clubstatus).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cbs)
}

// GetClubappStatus handles GET requests to retrieve a clubappstatus entity
// @Summary Get a clubappstatus entity by ID
// @Description get clubappstatus by ID
// @ID get-clubappstatus
// @Produce  json
// @Param id path int true "ClubappStatus ID"
// @Success 200 {object} ent.ClubappStatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubappstatuss/{id} [get]
func (ctl *ClubappStatusController) GetClubappStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cbs, err := ctl.client.ClubappStatus.
		Query().
		Where(clubappstatus.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cbs)
}

// ListClubappStatus handles request to get a list of clubappstatus entities
// @Summary List clubappstatus entities
// @Description list clubappstatus entities
// @ID list-clubappstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ClubappStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubappstatuss [get]
func (ctl *ClubappStatusController) ListClubappStatus(c *gin.Context) {
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

	clubappstatuss, err := ctl.client.ClubappStatus.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, clubappstatuss)
}

// DeleteClubappStatus handles DELETE requests to delete a clubappstatus entity
// @Summary Delete a clubappstatus entity by ID
// @Description get clubappstatus by ID
// @ID delete-clubappstatus
// @Produce  json
// @Param id path int true "ClubappStatus ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubappstatuss/{id} [delete]
func (ctl *ClubappStatusController) DeleteClubappStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ClubappStatus.
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

// UpdateClubappStatus handles PUT requests to update a clubappstatus entity
// @Summary Update a clubappstatus entity by ID
// @Description update clubappstatus by ID
// @ID update-clubappstatus
// @Accept   json
// @Produce  json
// @Param id path int true "ClubappStatus ID"
// @Param clubappstatus body ent.ClubappStatus true "ClubappStatus entity"
// @Success 200 {object} ent.ClubappStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubappstatuss/{id} [put]
func (ctl *ClubappStatusController) UpdateClubappStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.ClubappStatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "clubappstatus binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.ClubappStatus.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewClubappStatusController creates and registers handles for the clubappstatus controller
func NewClubappStatusController(router gin.IRouter, client *ent.Client) *ClubappStatusController {
	cbsc := &ClubappStatusController{
		client: client,
		router: router,
	}
	cbsc.register()
	return cbsc
}

// InitClubappStatusController registers routes to the main engine
func (ctl *ClubappStatusController) register() {
	clubappstatuss := ctl.router.Group("/clubappstatuss")

	clubappstatuss.GET("", ctl.ListClubappStatus)

	// CRUD
	clubappstatuss.POST("", ctl.CreateClubappStatus)
	clubappstatuss.GET(":id", ctl.GetClubappStatus)
	clubappstatuss.PUT(":id", ctl.UpdateClubappStatus)
	clubappstatuss.DELETE(":id", ctl.DeleteClubappStatus)
}
