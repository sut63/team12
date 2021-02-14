package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/purpose"
	"github.com/OMENX/app/ent/room"
	"github.com/OMENX/app/ent/roomuse"
	"github.com/OMENX/app/ent/user"
	"github.com/gin-gonic/gin"
)

// RoomuseController defines the struct for the roomuse controller
type RoomuseController struct {
	client *ent.Client
	router gin.IRouter
}

// Roomuse defines the struct for the roomuse
type Roomuse struct {
	RoomID    int
	PurposeID int
	UserID    int
	People    int
	Note      string
	Contact   string
	InTime    string
	OutTime   string
}

// CreateRoomuse handles POST requests for adding roomuse entities
// @Summary Create roomuse
// @Description Create roomuse
// @ID create-roomuse
// @Accept   json
// @Produce  json
// @Param roomuse body Roomuse true "Roomuse entity"
// @Success 200 {object} ent.Roomuse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomuses [post]
func (ctl *RoomuseController) CreateRoomuse(c *gin.Context) {
	obj := Roomuse{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "roomuse binding failed",
		})
		return
	}

	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.RoomID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "room not found",
		})
		return
	}

	p, err := ctl.client.Purpose.
		Query().
		Where(purpose.IDEQ(int(obj.PurposeID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "purpose not found",
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

	intimes, err := time.Parse(time.RFC3339, obj.InTime)
	outtimes, err := time.Parse(time.RFC3339, obj.OutTime)

	ru, err := ctl.client.Roomuse.
		Create().
		SetUsers(u).
		SetRooms(r).
		SetPurposes(p).
		SetPeople(obj.People).
		SetNote(obj.Note).
		SetContact(obj.Contact).
		SetInTime(intimes).
		SetOutTime(outtimes).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   ru,
	})
}

// ListRoomuse handles request to get a list of roomuse entities
// @Summary List roomuse entities
// @Description list roomuse entities
// @ID list-roomuse
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Roomuse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomuses [get]
func (ctl *RoomuseController) ListRoomuse(c *gin.Context) {
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

	roomuses, err := ctl.client.Roomuse.
		Query().
		WithUsers().
		WithRooms().
		WithPurposes().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, roomuses)
}

// GetRoomuse handles GET requests to retrieve a roomuse entity
// @Summary Get a roomuse entity by ID
// @Description get roomuse by ID
// @ID get-roomuse
// @Produce  json
// @Param id path int true "Roomuse ID"
// @Success 200 {array} ent.Roomuse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomuses/{id} [get]
func (ctl *RoomuseController) GetRoomuse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ru, err := ctl.client.Roomuse.
		Query().
		WithRooms().
		WithPurposes().
		WithUsers().
		Where(roomuse.HasRoomsWith(room.IDEQ(int(id)))).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ru)
}

// DeleteRoomuse handles DELETE requests to delete a roomuse entity
// @Summary Delete a roomuse entity by ID
// @Description get roomuse by ID
// @ID delete-roomuse
// @Produce  json
// @Param id path int true "Roomuse ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomuses/{id} [delete]
func (ctl *RoomuseController) DeleteRoomuse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Roomuse.
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

// NewRoomuseController creates and registers handles for the roomuse controller
func NewRoomuseController(router gin.IRouter, client *ent.Client) *RoomuseController {
	ruc := &RoomuseController{
		client: client,
		router: router,
	}
	ruc.register()
	return ruc
}

// InitRoomuseController registers routes to the main engine
func (ctl *RoomuseController) register() {
	roomuses := ctl.router.Group("/roomuses")

	roomuses.GET("", ctl.ListRoomuse)

	// CRUD
	roomuses.GET(":id", ctl.GetRoomuse)
	roomuses.POST("", ctl.CreateRoomuse)
	roomuses.DELETE(":id", ctl.DeleteRoomuse)
}
