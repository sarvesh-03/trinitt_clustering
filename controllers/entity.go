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

func CreateEntity(c echo.Context) error {
	var req CreateEntityRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Error-1")
	}

	db := config.GetDB()

	userId := uint(1)

	entity := models.Entity{
		Name:        req.Name,
		CreatedByID: userId,
	}

	if err := db.Create(&entity).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Error-2")

	}

	var response CreateEntityResponse

	response.Name = entity.Name
	response.ID = entity.ID

	return c.JSON(http.StatusOK, response)
}
