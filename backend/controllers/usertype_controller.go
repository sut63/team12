package controllers

import (
	"context"
	"strconv"

	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/usertype"
	"github.com/gin-gonic/gin"
)

// UsertypeController defines the struct for the usertype controller
type UsertypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateUsertype handles POST requests for adding usertype entities
// @Summary Create usertype
// @Description Create usertype
// @ID create-usertype
// @Accept   json
// @Produce  json
// @Paramusertype body ent.Usertype true "Usertype entity"
// @Success 200 {object} ent.Usertype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /usertypes [post]
func (ctl *UsertypeController) CreateUsertype(c *gin.Context) {
	obj := ent.Usertype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Usertype binding failed",
		})
		return
	}

	u, err := ctl.client.Usertype.
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

// GetUsertype handles GET requests to retrieve a usertype entity
// @Summary Get a usertype entity by ID
// @Description get usertype by ID
// @ID get-usertype
// @Produce  json
// @Param id path int true "Usertype ID"
// @Success 200 {object} ent.Usertype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /usertypes/{id} [get]
func (ctl *UsertypeController) GetUsertype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	t, err := ctl.client.Usertype.
		Query().
		Where(usertype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, t)
}

// ListUsertype handles request to get a list of usertype entities
// @Summary List usertype entities
// @Description list usertype entities
// @ID list-usertype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Usertype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /usertypes [get]
func (ctl *UsertypeController) ListUsertype(c *gin.Context) {
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

	Usertypes, err := ctl.client.Usertype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Usertypes)
}

// UpdateUsertype handles PUT requests to update a usertype entity
// @Summary Update a usertype entity by ID
// @Description update usertype by ID
// @ID update-usertype
// @Accept   json
// @Produce  json
// @Param id path int true "Usertype ID"
// @Param usertype body ent.Usertype true "Usertype entity"
// @Success 200 {object} ent.Usertype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /usertypes/{id} [put]
func (ctl *UsertypeController) UpdateUsertype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Usertype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Usertype binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Usertype.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewUsertypeController creates and registers handles for the usertype controller
func NewUsertypeController(router gin.IRouter, client *ent.Client) *UsertypeController {
	uc := &UsertypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitUsertypeController registers routes to the main engine
func (ctl *UsertypeController) register() {
	Usertypes := ctl.router.Group("/Usertypes")

	Usertypes.GET("", ctl.ListUsertype)

	 // CRUD
	Usertypes.POST("", ctl.CreateUsertype)
	Usertypes.GET(":id", ctl.GetUsertype)
	Usertypes.PUT(":id", ctl.UpdateUsertype)

}
