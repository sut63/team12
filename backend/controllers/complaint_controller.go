package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/complaint"
	"github.com/OMENX/app/ent/complainttype"
	"github.com/OMENX/app/ent/user"

	"github.com/gin-gonic/gin"
)

// ComplaintController defines the struct for the complaint controller
type ComplaintController struct {
	client *ent.Client
	router gin.IRouter
}

// Complaint defines the struct for the complaint
type Complaint struct {
	UserID int
	ClubID int
	TypeID int
	info   string
	date   string
}

// CreateComplaint handles POST requests for adding complaint entities
// @Summary Create complaint
// @Description Create complaint
// @ID create-complaint
// @Accept   json
// @Produce  json
// @Param complaint body Complaint true "Complaint entity"
// @Success 200 {object} ent.Complaint
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /complaints [post]
func (ctl *ComplaintController) CreateComplaint(c *gin.Context) {
	obj := Complaint{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "complaint binding failed",
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

	cb, err := ctl.client.Club.
		Query().
		Where(club.IDEQ(int(obj.ClubID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "club not found",
		})
		return
	}

	ct, err := ctl.client.ComplaintType.
		Query().
		Where(complainttype.IDEQ(int(obj.TypeID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "type not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.date)

	cp, err := ctl.client.Complaint.
		Create().
		SetComplaintToUser(u).
		SetComplaintToClub(cb).
		SetComplaintToComplaintType(ct).
		SetInfo(obj.info).
		SetDate(time).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cp)
}

// GetComplaint handles GET requests to retrieve a complaint entity
// @Summary Get a complaint entity by ID
// @Description get complaint by ID
// @ID get-complaint
// @Produce  json
// @Param id path int true "Complaint ID"
// @Success 200 {object} ent.Complaint
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /complaints/{id} [get]
func (ctl *ComplaintController) GetComplaint(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cp, err := ctl.client.Complaint.
		Query().
		Where(complaint.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cp)
}

// ListComplaint handles request to get a list of complaint entities
// @Summary List complaint entities
// @Description list complaint entities
// @ID list-complaint
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Complaint
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /complaints [get]
func (ctl *ComplaintController) ListComplaint(c *gin.Context) {
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

	complaints, err := ctl.client.Complaint.
		Query().
		WithComplaintToUser().
		WithComplaintToClub().
		WithComplaintToComplaintType().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, complaints)
}

// DeleteComplaint handles DELETE requests to delete a complaint entity
// @Summary Delete a complaint entity by ID
// @Description get complaint by ID
// @ID delete-complaint
// @Produce  json
// @Param id path int true "Complaint ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /complaints/{id} [delete]
func (ctl *ComplaintController) DeleteComplaint(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Complaint.
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

// UpdateComplaint handles PUT requests to update a complaint entity
// @Summary Update a complaint entity by ID
// @Description update complaint by ID
// @ID update-complaint
// @Accept   json
// @Produce  json
// @Param id path int true "Complaint ID"
// @Param complaint body ent.Complaint true "Complaint entity"
// @Success 200 {object} ent.Complaint
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /complaints/{id} [put]
func (ctl *ComplaintController) UpdateComplaint(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Complaint{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "complaint binding failed",
		})
		return
	}
	obj.ID = int(id)
	cp, err := ctl.client.Complaint.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, cp)
}

// NewComplaintController creates and registers handles for the complaint controller
func NewComplaintController(router gin.IRouter, client *ent.Client) *ComplaintController {
	cpc := &ComplaintController{
		client: client,
		router: router,
	}
	cpc.register()
	return cpc
}

// InitComplaintController registers routes to the main engine
func (ctl *ComplaintController) register() {
	complaints := ctl.router.Group("/complaints")

	complaints.GET("", ctl.ListComplaint)

	// CRUD
	complaints.POST("", ctl.CreateComplaint)
	complaints.GET(":id", ctl.GetComplaint)
	complaints.PUT(":id", ctl.UpdateComplaint)
	complaints.DELETE(":id", ctl.DeleteComplaint)
}
