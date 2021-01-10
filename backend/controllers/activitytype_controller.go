package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/activitytype"
	"github.com/gin-gonic/gin"
)

// ActivityTypeController defines the struct for the ActivityType controller
type ActivityTypeController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateActivityType handles POST requests for adding ActivityType entities
// @Summary Create ActivityType
// @Description Create ActivityType
// @ID create-ActivityType
// @Accept   json
// @Produce  json
// @Param ActivityType body ent.ActivityType true "ActivityType entity"
// @Success 200 {object} ent.ActivityType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ActivityTypes [post]
func (ctl *ActivityTypeController) CreateActivityType(c *gin.Context) {
	obj := ent.ActivityType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "ActivityType binding failed",
		})
		return
	}
  
	ue, err := ctl.client.ActivityType.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, ue)
 }

// GetActivityType handles GET requests to retrieve a ActivityType entity
// @Summary Get a ActivityType entity by ID
// @Description get ActivityType by ID
// @ID get-ActivityType
// @Produce  json
// @Param id path int true "ActivityType ID"
// @Success 200 {object} ent.ActivityType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ActivityTypes/{id} [get]
func (ctl *ActivityTypeController) GetActivityType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.ActivityType.
		Query().
		Where(activitytype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }
 
 // ListActivityType handles request to get a list of ActivityType entities
// @Summary List ActivityType entities
// @Description list ActivityType entities
// @ID list-ActivityType
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ActivityType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ActivityTypes [get]
func (ctl *ActivityTypeController) ListActivityType(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	ActivityTypes, err := ctl.client.ActivityType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, ActivityTypes)
 }

// DeleteActivityType handles DELETE requests to delete a ActivityType entity
// @Summary Delete a ActivityType entity by ID
// @Description get ActivityType by ID
// @ID delete-ActivityType
// @Produce  json
// @Param id path int true "ActivityType ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ActivityTypes/{id} [delete]
func (ctl *ActivityTypeController) DeleteActivityType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.ActivityType.
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

 // UpdateActivityType handles PUT requests to update a ActivityType entity
// @Summary Update a ActivityType entity by ID
// @Description update ActivityType by ID
// @ID update-ActivityType
// @Accept   json
// @Produce  json
// @Param id path int true "ActivityType ID"
// @Param ActivityType body ent.ActivityType true "ActivityType entity"
// @Success 200 {object} ent.ActivityType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ActivityTypes/{id} [put]
func (ctl *ActivityTypeController) UpdateActivityType(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }
 
   obj := ent.ActivityType{}
   if err := c.ShouldBind(&obj); err != nil {
       c.JSON(400, gin.H{
           "error": "ActivityType binding failed",
       })
       return
   }
   obj.ID = int(id)
   u, err := ctl.client.ActivityType.
       UpdateOne(&obj).
       Save(context.Background())
   if err != nil {
       c.JSON(400, gin.H{"error": "update failed",})
       return
   }
 
   c.JSON(200, u)
}

// NewActivityTypeController creates and registers handles for the ActivityType controller
func NewActivityTypeController(router gin.IRouter, client *ent.Client) *ActivityTypeController {
	uc := &ActivityTypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitActivityTypeController registers routes to the main engine
 func (ctl *ActivityTypeController) register() {
	ActivityTypes := ctl.router.Group("/ActivityTypes")
  
	ActivityTypes.GET("", ctl.ListActivityType)
  
	// CRUD
	ActivityTypes.POST("", ctl.CreateActivityType)
	ActivityTypes.GET(":id", ctl.GetActivityType)
	ActivityTypes.PUT(":id", ctl.UpdateActivityType)
	ActivityTypes.DELETE(":id", ctl.DeleteActivityType)
 }
 
