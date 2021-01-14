package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/clubtype"
	"github.com/gin-gonic/gin"
)

// ClubTypeController defines the struct for the clubtype controller
type ClubTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateClubType handles POST requests for adding clubtype entities
// @Summary Create ClubType
// @Description Create clubtype
// @ID create-clubtype
// @Accept   json
// @Produce  json
// @Param clubtype body ent.ClubType true "ClubType entity"
// @Success 200 {object} ent.ClubType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubtype [post]
func (ctl *ClubTypeController) CreateClubType(c *gin.Context) {
	obj := ent.ClubType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "clubtype binding failed",
		})
		return
	}

	u, err := ctl.client.ClubType.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetClubType handles GET requests to retrieve a clubtype entity
// @Summary Get a clubtype entity by ID
// @Description get clubtype by ID
// @ID get-clubtype
// @Produce  json
// @Param id path int true "ClubType ID"
// @Success 200 {object} ent.ClubType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubtype/{id} [get]
func (ctl *ClubTypeController) GetClubType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.ClubType.
		Query().
		Where(clubtype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListClubType handles request to get a list of clubtype entities
// @Summary List clubtype entities
// @Description list clubtype entities
// @ID list-clubtype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ClubType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubtype [get]
func (ctl *ClubTypeController) ListClubType(c *gin.Context) {
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

	clubtype, err := ctl.client.ClubType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, clubtype)
}

// DeleteClubType handles DELETE requests to delete a clubtype entity
// @Summary Delete a clubtype entity by ID
// @Description get clubtype by ID
// @ID delete-clubtype
// @Produce  json
// @Param id path int true "ClubType ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubtype/{id} [delete]
func (ctl *ClubTypeController) DeleteClubType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ClubType.
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

// UpdateClubType handles PUT requests to update a clubtype entity
// @Summary Update a clubtype entity by ID
// @Description update clubtype by ID
// @ID update-clubtype
// @Accept   json
// @Produce  json
// @Param id path int true "ClubType ID"
// @Param clubtype body ent.ClubType true "ClubType entity"
// @Success 200 {object} ent.ClubType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubtype/{id} [put]
func (ctl *ClubTypeController) UpdateClubType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.ClubType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "clubtype binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.ClubType.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewClubTypeController creates and registers handles for the clubtype controller
func NewClubTypeController(router gin.IRouter, client *ent.Client) *ClubTypeController {
	uc := &ClubTypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitClubTypeController registers routes to the main engine
func (ctl *ClubTypeController) register() {
	clubtype := ctl.router.Group("/clubtype")

	clubtype.GET("", ctl.ListClubType)

	// CRUD
	clubtype.POST("", ctl.CreateClubType)
	clubtype.GET(":id", ctl.GetClubType)
	clubtype.PUT(":id", ctl.UpdateClubType)
	clubtype.DELETE(":id", ctl.DeleteClubType)
}

