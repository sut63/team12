package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/club"
	"github.com/OMENX/app/ent/clubapplication"
	"github.com/OMENX/app/ent/clubappstatus"
	"github.com/OMENX/app/ent/user"
	"github.com/gin-gonic/gin"
)

// ClubapplicationController defines the struct for the clubapplication controller
type ClubapplicationController struct {
	client *ent.Client
	router gin.IRouter
}

// Clubapplication defines the struct for the clubapplication ...
type Clubapplication struct {
	Contact     string
	Reason      string
	AdderName   string
	Discipline  string
	AdderGender string
	AdderAge    int
	Year        int
	ClubID      int
	UserID      int
	Addedtime   string
	AppstatusID int
}

// CreateClubapplication handles POST requests for adding clubapplication entities
// @Summary Create clubapplication
// @Description Create clubapplication
// @ID create-clubapplication
// @Accept   json
// @Produce  json
// @Param clubapplication body Clubapplication true "Clubapplication entity"
// @Success 200 {object} ent.Clubapplication
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubapplications [post]
func (ctl *ClubapplicationController) CreateClubapplication(c *gin.Context) {
	obj := Clubapplication{}
	if err := c.ShouldBind(&obj); err != nil {

		c.JSON(400, gin.H{
			"error": "clubapplication binding failed",
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

	l, err := ctl.client.Club.
		Query().
		Where(club.IDEQ(int(obj.ClubID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "club not found",
		})
		return
	}

	a, err := ctl.client.ClubappStatus.
		Query().
		Where(clubappstatus.IDEQ(int(obj.AppstatusID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "applicationStatus not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Addedtime)

	ca, err := ctl.client.Clubapplication.
		Create().
		SetOwner(u).
		SetAddername(obj.AdderName).
		SetAge(obj.AdderAge).
		SetGender(obj.AdderGender).
		SetYaer(obj.Year).
		SetDiscipline(obj.Discipline).
		SetContact(obj.Contact).
		SetReason(obj.Reason).
		SetClub(l).
		SetClubappstatus(a).
		SetAddeddatetime(time).
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
		"data":   ca,
	})
}

// GetClubapplication handles GET requests to retrieve a clubapplication entity
// @Summary Get a clubapplication entity by ID
// @Description get clubapplication by ID
// @ID get-clubapplication
// @Produce  json
// @Param id path int true "Clubapplication ID"
// @Param cid query int false "Club ID"
// @Param sid query int false "ClabappStatus ID"
// @Success 200 {object} ent.Clubapplication
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubapplications/{id} [get]
func (ctl *ClubapplicationController) GetClubapplication(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cid, err := strconv.ParseInt(c.Query("cid"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	sid, err := strconv.ParseInt(c.Query("sid"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	switch {
		case int(cid) == 0 && int(sid) == 0 :
			ca, err := ctl.client.Clubapplication.
				Query().
				WithOwner().
				WithClub().
				WithClubappstatus().
				Where(clubapplication.HasOwnerWith(user.IDEQ(int(id)))).
				All(context.Background())

			if err != nil {
				c.JSON(404, gin.H{
					"error": err.Error(),
				})
				return
			}
		
			c.JSON(200, ca)

		case int(cid) == 0 :
			ca, err := ctl.client.Clubapplication.
				Query().
				WithOwner().
				WithClub().
				WithClubappstatus().
				Where(clubapplication.HasOwnerWith(user.IDEQ(int(id))), clubapplication.HasClubappstatusWith(clubappstatus.IDEQ(int(sid)))).
				All(context.Background())

			if err != nil {
				c.JSON(404, gin.H{
					"error": err.Error(),
				})
				return
			}
		
			c.JSON(200, ca)

		case int(sid) == 0 :
			ca, err := ctl.client.Clubapplication.
				Query().
				WithOwner().
				WithClub().
				WithClubappstatus().
				Where(clubapplication.HasOwnerWith(user.IDEQ(int(id))), clubapplication.HasClubWith(club.IDEQ(int(cid)))).
				All(context.Background())

			if err != nil {
				c.JSON(404, gin.H{
					"error": err.Error(),
				})
				return
			}
		
			c.JSON(200, ca)

		default:
			ca, err := ctl.client.Clubapplication.
				Query().
				WithOwner().
				WithClub().
				WithClubappstatus().
				Where(clubapplication.HasOwnerWith(user.IDEQ(int(id))), clubapplication.HasClubWith(club.IDEQ(int(cid))), clubapplication.HasClubappstatusWith(clubappstatus.IDEQ(int(sid)))).
				All(context.Background())
				
			if err != nil {
				c.JSON(404, gin.H{
					"error": err.Error(),
				})
				return
			}
		
			c.JSON(200, ca)
	}
}

// ListClubapplication handles request to get a list of clubapplication entities
// @Summary List clubapplication entities
// @Description list clubapplication entities
// @ID list-clubapplication
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Clubapplication
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubapplications [get]
func (ctl *ClubapplicationController) ListClubapplication(c *gin.Context) {
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

	clubapplications, err := ctl.client.Clubapplication.
		Query().
		WithOwner().
		WithClub().
		WithClubappstatus().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, clubapplications)
}

// DeleteClubapplication handles DELETE requests to delete a clubapplication entity
// @Summary Delete a clubapplication entity by ID
// @Description get clubapplication by ID
// @ID delete-clubapplication
// @Produce  json
// @Param id path int true "Clubapplication ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubapplications/{id} [delete]
func (ctl *ClubapplicationController) DeleteClubapplication(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Clubapplication.
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

// UpdateClubapplication handles PUT requests to update a clubapplication entity
// @Summary Update a clubapplication entity by ID
// @Description update clubapplication by ID
// @ID update-clubapplication
// @Accept   json
// @Produce  json
// @Param id path int true "Clubapplication ID"
// @Param clubapplication body ent.Clubapplication true "Clubapplication entity"
// @Success 200 {object} ent.Clubapplication
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /clubapplications/{id} [put]
func (ctl *ClubapplicationController) UpdateClubapplication(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Clubapplication{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "clubapplication binding failed",
		})
		return
	}
	obj.ID = int(id)
	ca, err := ctl.client.Clubapplication.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ca)
}

// NewClubapplicationController creates and registers handles for the clubapplication controller
func NewClubapplicationController(router gin.IRouter, client *ent.Client) *ClubapplicationController {
	cac := &ClubapplicationController{
		client: client,
		router: router,
	}
	cac.register()
	return cac
}

// InitClubapplicationController registers routes to the main engine
func (ctl *ClubapplicationController) register() {
	clubapplications := ctl.router.Group("/clubapplications")

	clubapplications.GET("", ctl.ListClubapplication)

	// CRUD
	clubapplications.POST("", ctl.CreateClubapplication)
	clubapplications.GET(":id", ctl.GetClubapplication)
	clubapplications.PUT(":id", ctl.UpdateClubapplication)
	clubapplications.DELETE(":id", ctl.DeleteClubapplication)
}
