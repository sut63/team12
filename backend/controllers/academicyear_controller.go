package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/academicyear"
	"github.com/gin-gonic/gin"
)

// AcademicYearController defines the struct for the AcademicYear controller
type AcademicYearController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateAcademicYear handles POST requests for adding AcademicYear entities
// @Summary Create AcademicYear
// @Description Create AcademicYear
// @ID create-AcademicYear
// @Accept   json
// @Produce  json
// @Param AcademicYear body ent.AcademicYear true "AcademicYear entity"
// @Success 200 {object} ent.AcademicYear
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /AcademicYears [post]
func (ctl *AcademicYearController) CreateAcademicYear(c *gin.Context) {
	obj := ent.AcademicYear{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "AcademicYear binding failed",
		})
		return
	}
  
	u, err := ctl.client.AcademicYear.
		Create().
		SetSemester(obj.Semester).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }

// GetAcademicYear handles GET requests to retrieve a AcademicYear entity
// @Summary Get a AcademicYear entity by ID
// @Description get AcademicYear by ID
// @ID get-AcademicYear
// @Produce  json
// @Param id path int true "AcademicYear ID"
// @Success 200 {object} ent.AcademicYear
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /AcademicYears/{id} [get]
func (ctl *AcademicYearController) GetAcademicYear(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.AcademicYear.
		Query().
		Where(academicyear.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }
 
 // ListAcademicYear handles request to get a list of AcademicYear entities
// @Summary List AcademicYear entities
// @Description list AcademicYear entities
// @ID list-AcademicYear
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.AcademicYear
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /AcademicYears [get]
func (ctl *AcademicYearController) ListAcademicYear(c *gin.Context) {
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
  
	AcademicYears, err := ctl.client.AcademicYear.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, AcademicYears)
 }

// DeleteAcademicYear handles DELETE requests to delete a AcademicYear entity
// @Summary Delete a AcademicYear entity by ID
// @Description get AcademicYear by ID
// @ID delete-AcademicYear
// @Produce  json
// @Param id path int true "AcademicYear ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /AcademicYears/{id} [delete]
func (ctl *AcademicYearController) DeleteAcademicYear(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.AcademicYear.
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

 // UpdateAcademicYear handles PUT requests to update a AcademicYear entity
// @Summary Update a AcademicYear entity by ID
// @Description update AcademicYear by ID
// @ID update-AcademicYear
// @Accept   json
// @Produce  json
// @Param id path int true "AcademicYear ID"
// @Param AcademicYear body ent.AcademicYear true "AcademicYear entity"
// @Success 200 {object} ent.AcademicYear
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /AcademicYears/{id} [put]
func (ctl *AcademicYearController) UpdateAcademicYear(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }
 
   obj := ent.AcademicYear{}
   if err := c.ShouldBind(&obj); err != nil {
       c.JSON(400, gin.H{
           "error": "AcademicYear binding failed",
       })
       return
   }
   obj.ID = int(id)
   u, err := ctl.client.AcademicYear.
       UpdateOne(&obj).
       Save(context.Background())
   if err != nil {
       c.JSON(400, gin.H{"error": "update failed",})
       return
   }
 
   c.JSON(200, u)
}

// NewAcademicYearController creates and registers handles for the AcademicYear controller
func NewAcademicYearController(router gin.IRouter, client *ent.Client) *AcademicYearController {
	uc := &AcademicYearController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitAcademicYearController registers routes to the main engine
 func (ctl *AcademicYearController) register() {
	AcademicYears := ctl.router.Group("/AcademicYears")
  
	AcademicYears.GET("", ctl.ListAcademicYear)
  
	// CRUD
	AcademicYears.POST("", ctl.CreateAcademicYear)
	AcademicYears.GET(":id", ctl.GetAcademicYear)
	AcademicYears.PUT(":id", ctl.UpdateAcademicYear)
	AcademicYears.DELETE(":id", ctl.DeleteAcademicYear)
 }
 