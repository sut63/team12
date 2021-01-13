package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/complainttype"
	"github.com/gin-gonic/gin"
)

// ComplainttypeController defines the struct for the complainttype controller
type ComplainttypeController struct {
	client *ent.Client
	router gin.IRouter
}

// Complainttype defines the struct for the complainttype
type Complainttype struct {
	Description string
}

// CreateComplainttype handles POST requests for adding complainttype entities
// @Summary Create complainttype
// @Description Create complainttype
// @ID create-complainttype
// @Accept   json
// @Produce  json
// @Param complainttype body Complainttype true "Complainttype entity"
// @Success 200 {object} ent.Complainttype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /complainttypes [post]
func (ctl *ComplainttypeController) CreateComplainttype(c *gin.Context) {
	obj := ent.ComplaintType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "complainttype binding failed",
		})
		return
	}

	cpt, err := ctl.client.ComplaintType.
		Create().
		SetDescription(obj.Description).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cpt)
}

// GetComplainttype handles GET requests to retrieve a complainttype entity
// @Summary Get a complainttype entity by ID
// @Description get complainttype by ID
// @ID get-complainttype
// @Produce  json
// @Param id path int true "Complainttype ID"
// @Success 200 {object} ent.Complainttype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /complainttypes/{id} [get]
func (ctl *ComplainttypeController) GetComplainttype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cpt, err := ctl.client.ComplaintType.
		Query().
		Where(complainttype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cpt)
}

// ListComplainttype handles request to get a list of complainttype entities
// @Summary List complainttype entities
// @Description list complainttype entities
// @ID list-complainttype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Complainttype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /complainttypes [get]
func (ctl *ComplainttypeController) ListComplainttype(c *gin.Context) {
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

	complainttypes, err := ctl.client.ComplaintType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, complainttypes)
}

// DeleteComplainttype handles DELETE requests to delete a complainttype entity
// @Summary Delete a complainttype entity by ID
// @Description get complainttype by ID
// @ID delete-complainttype
// @Produce  json
// @Param id path int true "Complainttype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /complainttypes/{id} [delete]
func (ctl *ComplainttypeController) DeleteComplainttype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ComplaintType.
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

// UpdateComplainttype handles PUT requests to update a complainttype entity
// @Summary Update a complainttype entity by ID
// @Description update complainttype by ID
// @ID update-complainttype
// @Accept   json
// @Produce  json
// @Param id path int true "Complainttype ID"
// @Param complainttype body ent.Complainttype true "Complainttype entity"
// @Success 200 {object} ent.Complainttype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /complainttypes/{id} [put]
func (ctl *ComplainttypeController) UpdateComplainttype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.ComplaintType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "complainttype binding failed",
		})
		return
	}
	obj.ID = int(id)
	cpt, err := ctl.client.ComplaintType.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, cpt)
}

// NewComplainttypeController creates and registers handles for the complainttype controller
func NewComplainttypeController(router gin.IRouter, client *ent.Client) *ComplainttypeController {
	cptc := &ComplainttypeController{
		client: client,
		router: router,
	}
	cptc.register()
	return cptc
}

// InitComplainttypeController registers routes to the main engine
func (ctl *ComplainttypeController) register() {
	complainttypes := ctl.router.Group("/complainttypes")

	complainttypes.GET("", ctl.ListComplainttype)

	// CRUD
	complainttypes.POST("", ctl.CreateComplainttype)
	complainttypes.GET(":id", ctl.GetComplainttype)
	complainttypes.PUT(":id", ctl.UpdateComplainttype)
	complainttypes.DELETE(":id", ctl.DeleteComplainttype)
}
