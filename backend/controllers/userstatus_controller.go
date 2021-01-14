package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/OMENX/app/ent"
	"github.com/OMENX/app/ent/userstatus"
)

// UserStatusController is ...
type UserStatusController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateUserStatus handles POST requests for adding userstatus entities
// @Summary Create userstatus
// @Description Create userstatus
// @ID create-userstatus
// @Accept   json
// @Produce  json
// @Param userstatus body ent.UserStatus true "UserStatus entity"
// @Success 200 {object} ent.UserStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatuss [post]
func (ctl *UserStatusController) CreateUserStatus(c *gin.Context) {
	obj := ent.UserStatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "userstatus binding failed",
		})
		return
	}

	us, err := ctl.client.UserStatus.
		Create().
		SetUserstatus(obj.Userstatus).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, us)
}

// GetUserStatus handles GET requests to retrieve a userstatus entity
// @Summary Get a userstatus entity by ID
// @Description get userstatus by ID
// @ID get-userstatus
// @Produce  json
// @Param id path int true "UserStatus ID"
// @Success 200 {object} ent.UserStatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatuss/{id} [get]
func (ctl *UserStatusController) GetUserStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	us, err := ctl.client.UserStatus.
		Query().
		Where(userstatus.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, us)
}

// ListUserStatus handles request to get a list of userstatus entities
// @Summary List userstatus entities
// @Description list userstatus entities
// @ID list-userstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.UserStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatuss [get]
func (ctl *UserStatusController) ListUserStatus(c *gin.Context) {
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

	userstatuss, err := ctl.client.UserStatus.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, userstatuss)
}

// DeleteUserStatus handles DELETE requests to delete a userstatus entity
// @Summary Delete a userstatus entity by ID
// @Description get userstatus by ID
// @ID delete-userstatus
// @Produce  json
// @Param id path int true "UserStatus ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatuss/{id} [delete]
func (ctl *UserStatusController) DeleteUserStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.UserStatus.
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

// UpdateUserStatus handles PUT requests to update a userstatus entity
// @Summary Update a userstatus entity by ID
// @Description update userstatus by ID
// @ID update-userstatus
// @Accept   json
// @Produce  json
// @Param id path int true "UserStatus ID"
// @Param userstatus body ent.UserStatus true "UserStatus entity"
// @Success 200 {object} ent.UserStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatuss/{id} [put]
func (ctl *UserStatusController) UpdateUserStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.UserStatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "userstatus binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	usr, err := ctl.client.UserStatus.
		UpdateOneID(int(id)).
		SetUserstatus(obj.Userstatus).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, usr)
}

// NewUserstatusController creates and registers handles for the userstatus controller
func NewUserstatusController(router gin.IRouter, client *ent.Client) *UserStatusController {
	usrc := &UserStatusController{
		client: client,
		router: router,
	}

	usrc.register()
	return usrc

}

func (ctl *UserStatusController) register() {
	userstatuss := ctl.router.Group("/userstatuss")

	userstatuss.GET("", ctl.ListUserStatus)

	//CRUD
	userstatuss.POST("", ctl.CreateUserStatus)
	userstatuss.GET(":id", ctl.GetUserStatus)
	userstatuss.PUT(":id", ctl.UpdateUserStatus)
	userstatuss.DELETE(":id", ctl.DeleteUserStatus)
}
