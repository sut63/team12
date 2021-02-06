package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/clubbranch"
	"github.com/OMENX/app/ent/clubtype"
	"github.com/OMENX/app/ent/user"

	"github.com/gin-gonic/gin"
)

// ClubController defines the struct for the club controller
type ClubController struct {
	client *ent.Client
	router gin.IRouter
}

type Club struct {
	UserID       int
	ClubBranchID int
	ClubTypeID   int
	Name         string
	Purpose      string
	Phone        string
}

// CreateClub handles POST requests for adding club entities
// @Summary Create club
// @Description Create club
// @ID create-club
// @Accept   json
// @Produce  json
// @Param club body ent.Club true "Club entity"
// @Success 200 {object} ent.Club
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /club [post]
func (ctl *ClubController) CreateClub(c *gin.Context) {
	obj := Club{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "club binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	cb, err := ctl.client.ClubBranch.
		Query().
		Where(clubbranch.IDEQ(int(obj.ClubBranchID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "clubbranch not found",
		})
		return
	}

	ct, err := ctl.client.ClubType.
		Query().
		Where(clubtype.IDEQ(int(obj.ClubTypeID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "clubtype not found",
		})
		return
	}

	cl, err := ctl.client.Club.
		Create().
		SetUser(u).
		SetClubbranch(cb).
		SetClubtype(ct).
		SetPurpose(obj.Purpose).
		SetPhone(obj.Phone).
		SetName(obj.Name).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   cl,
	})
}

// GetClub handles GET requests to retrieve a club entity
// @Summary Get a club entity by ID
// @Description get club by ID
// @ID get-club
// @Produce  json
// @Param id path int true "Club ID"
// @Success 200 {array} ent.Club
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /club/{id} [get]
func (ctl *ClubController) GetClub(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Club.
		Query().
		WithClubbranch().
		WithClubtype().
		WithUser().
		// 	Where(checkout.HasStatusopinionWith(statusopinion.IDEQ(int(id)))).
		// Where(club.IDEQ(int(id))).
		Where(club.HasClubtypeWith(clubtype.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListClub handles request to get a list of club entities
// @Summary List club entities
// @Description list club entities
// @ID list-club
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Club
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /club [get]
func (ctl *ClubController) ListClub(c *gin.Context) {
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

	club, err := ctl.client.Club.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, club)
}

// DeleteClub handles DELETE requests to delete a club entity
// @Summary Delete a club entity by ID
// @Description get club by ID
// @ID delete-club
// @Produce  json
// @Param id path int true "Club ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /club/{id} [delete]
func (ctl *ClubController) DeleteClub(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Club.
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

// UpdateClub handles PUT requests to update a club entity
// @Summary Update a club entity by ID
// @Description update club by ID
// @ID update-club
// @Accept   json
// @Produce  json
// @Param id path int true "Club ID"
// @Param club body ent.Club true "Club entity"
// @Success 200 {object} ent.Club
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /club/{id} [put]
func (ctl *ClubController) UpdateClub(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Club{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "club binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Club.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewClubController creates and registers handles for the club controller
func NewClubController(router gin.IRouter, client *ent.Client) *ClubController {
	uc := &ClubController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitClubController registers routes to the main engine
func (ctl *ClubController) register() {
	club := ctl.router.Group("/club")

	club.GET("", ctl.ListClub)

	// CRUD
	club.POST("", ctl.CreateClub)
	club.GET(":id", ctl.GetClub)
	club.PUT(":id", ctl.UpdateClub)
	club.DELETE(":id", ctl.DeleteClub)
}

