package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trinitt/config"
	"github.com/trinitt/models"
)

type CreateEntityRequest struct {
	Name string `json:"name"`
}

type CreateEntityResponse struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

func CreateEntity(c echo.Context) {
	var req CreateEntityRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	db := config.GetDB()

	userId := c.Get("user").(uint)

	entity := models.Entity{
		Name:        req.Name,
		CreatedByID: userId,
	}

	if err := db.Create(&entity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var response CreateEntityResponse

	response.Name = entity.Name
	response.ID = entity.ID

	c.JSON(http.StatusOK, response)
}
