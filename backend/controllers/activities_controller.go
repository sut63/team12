package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/academicyear"
	"github.com/OMENX/app/ent/activities"
	"github.com/OMENX/app/ent/activitytype"
	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/user"

	"github.com/gin-gonic/gin"
)

// ActivitiesController defines the struct for the Activities controller
type ActivitiesController struct {
   client *ent.Client
   router gin.IRouter
}

type Activities struct {
	AcademicYearID   int
	ActivityTypeID	 int
	ClubID	int
	Name 		string
	Detail 		string
	Endtime     string
	Starttime   string
}

// CreateActivities handles POST requests for adding activities entities
// @Summary Create activities
// @Description Create activities
// @ID create-activities
// @Accept   json
// @Produce  json
// @Param activities body ent.Activities true "Activities entity"
// @Success 200 {object} ent.Activities
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /activities [post]
func (ctl *ActivitiesController) CreateActivities(c *gin.Context) {
	obj := Activities{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Activities binding failed",
		})
		return
	}

	St, err := time.Parse(time.RFC3339, obj.Starttime)
	E, err := time.Parse(time.RFC3339, obj.Endtime)

	acy, err := ctl.client.AcademicYear.
		Query().
		Where(academicyear.IDEQ(int(obj.AcademicYearID))).
		Only(context.Background())
 
	if err != nil {
		c.JSON(400, gin.H{
			"error": "academicyear not found",
		})
		return
	}

	act, err := ctl.client.ActivityType.
		Query().
		Where(activitytype.IDEQ(int(obj.ActivityTypeID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "activitytype not found",
		})
		return
	}
		

	findclub, err := ctl.client.User.
	Query().
	Where(user.IDEQ(int(obj.ClubID))).
	Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "this user not found",
		})
		return
	} 

	/* cl, err:= ctl.client.User.
	Query().
	Select(user.).
	String(context.Background()) */
	

	club, err := ctl.client.Club.
	Query().
	Where(club.IDEQ(int(findclub.ID))).
	Only(context.Background())
	
	at, err := ctl.client.Activities.
		Create().
		SetAcademicyear(acy).
		SetActivitytype(act).
		SetClub(club).
		SetDetail(obj.Detail).
		SetEndtime(E).
		SetName(obj.Name).
		SetStarttime(St).
		Save(context.Background())
		 
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed", 
		})
		return
	}
  
	c.JSON(200, at)
 }

// GetActivities handles GET requests to retrieve a Activities entity
// @Summary Get a Activities entity by ID
// @Description get Activities by ID
// @ID get-Activities
// @Produce  json
// @Param id path int true "Activities ID"
// @Success 200 {object} ent.Activities
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Activitiess/{id} [get]
func (ctl *ActivitiesController) GetActivities(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Activities.
		Query().
		Where(activities.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }
 
 // ListActivities handles request to get a list of Activities entities
// @Summary List Activities entities
// @Description list Activities entities
// @ID list-Activities
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Activities
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Activitiess [get]
func (ctl *ActivitiesController) ListActivities(c *gin.Context) {
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
  
	Activitiess, err := ctl.client.Activities.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, Activitiess)
 }

// DeleteActivities handles DELETE requests to delete a Activities entity
// @Summary Delete a Activities entity by ID
// @Description get Activities by ID
// @ID delete-Activities
// @Produce  json
// @Param id path int true "Activities ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Activitiess/{id} [delete]
func (ctl *ActivitiesController) DeleteActivities(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Activities.
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

 // UpdateActivities handles PUT requests to update a Activities entity
// @Summary Update a Activities entity by ID
// @Description update Activities by ID
// @ID update-Activities
// @Accept   json
// @Produce  json
// @Param id path int true "Activities ID"
// @Param Activities body ent.Activities true "Activities entity"
// @Success 200 {object} ent.Activities
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Activitiess/{id} [put]
func (ctl *ActivitiesController) UpdateActivities(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }
 
   obj := ent.Activities{}
   if err := c.ShouldBind(&obj); err != nil {
       c.JSON(400, gin.H{
           "error": "Activities binding failed",
       })
       return
   }
   obj.ID = int(id)
   u, err := ctl.client.Activities.
       UpdateOne(&obj).
       Save(context.Background())
   if err != nil {
       c.JSON(400, gin.H{"error": "update failed",})
       return
   }
 
   c.JSON(200, u)
}

// NewActivitiesController creates and registers handles for the Activities controller
func NewActivitiesController(router gin.IRouter, client *ent.Client) *ActivitiesController {
	uc := &ActivitiesController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitActivitiesController registers routes to the main engine
 func (ctl *ActivitiesController) register() {
	activities := ctl.router.Group("/activities")
  
	activities.GET("", ctl.ListActivities)
  
	// CRUD
	activities.POST("", ctl.CreateActivities)
	activities.GET(":id", ctl.GetActivities)
	activities.PUT(":id", ctl.UpdateActivities)
	activities.DELETE(":id", ctl.DeleteActivities)
 }
 