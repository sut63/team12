package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/clubbranch"
	"github.com/gin-gonic/gin"
)

// ClubBranchController defines the struct for the clubbranch controller
type ClubBranchController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateClubBranch handles POST requests for adding clubbranch entities
// @Summary Create ClubBranch
// @Description Create clubbranch
// @ID create-clubbranch
// @Accept   json
// @Produce  json
// @Param clubbranch body ent.ClubBranch true "ClubBranch entity"
// @Success 200 {object} ent.ClubBranch
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubbranch [post]
func (ctl *ClubBranchController) CreateClubBranch(c *gin.Context) {
	obj := ent.ClubBranch{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "clubbranch binding failed",
		})
		return
	}

	u, err := ctl.client.ClubBranch.
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

// GetClubBranch handles GET requests to retrieve a clubbranch entity
// @Summary Get a clubbranch entity by ID
// @Description get clubbranch by ID
// @ID get-clubbranch
// @Produce  json
// @Param id path int true "ClubBranch ID"
// @Success 200 {object} ent.ClubBranch
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubbranch/{id} [get]
func (ctl *ClubBranchController) GetClubBranch(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.ClubBranch.
		Query().
		Where(clubbranch.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListClubBranch handles request to get a list of clubbranch entities
// @Summary List clubbranch entities
// @Description list clubbranch entities
// @ID list-clubbranch
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ClubBranch
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubbranch [get]
func (ctl *ClubBranchController) ListClubBranch(c *gin.Context) {
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

	clubbranch, err := ctl.client.ClubBranch.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, clubbranch)
}

// DeleteClubBranch handles DELETE requests to delete a clubbranch entity
// @Summary Delete a clubbranch entity by ID
// @Description get clubbranch by ID
// @ID delete-clubbranch
// @Produce  json
// @Param id path int true "ClubBranch ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubbranch/{id} [delete]
func (ctl *ClubBranchController) DeleteClubBranch(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ClubBranch.
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

// UpdateClubBranch handles PUT requests to update a clubbranch entity
// @Summary Update a clubbranch entity by ID
// @Description update clubbranch by ID
// @ID update-clubbranch
// @Accept   json
// @Produce  json
// @Param id path int true "ClubBranch ID"
// @Param clubbranch body ent.ClubBranch true "ClubBranch entity"
// @Success 200 {object} ent.ClubBranch
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubbranch/{id} [put]
func (ctl *ClubBranchController) UpdateClubBranch(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.ClubBranch{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "clubbranch binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.ClubBranch.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewClubBranchController creates and registers handles for the clubbranch controller
func NewClubBranchController(router gin.IRouter, client *ent.Client) *ClubBranchController {
	uc := &ClubBranchController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitClubBranchController registers routes to the main engine
func (ctl *ClubBranchController) register() {
	clubbranch := ctl.router.Group("/clubbranch")

	clubbranch.GET("", ctl.ListClubBranch)

	// CRUD
	clubbranch.POST("", ctl.CreateClubBranch)
	clubbranch.GET(":id", ctl.GetClubBranch)
	clubbranch.PUT(":id", ctl.UpdateClubBranch)
	clubbranch.DELETE(":id", ctl.DeleteClubBranch)
}

