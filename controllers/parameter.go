package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trinitt/config"
	"github.com/trinitt/models"
)

type AddParamToEntityRequest struct {
	Name     string               `json:"name"`
	Type     models.ParameterType `json:"type"`
	EntityID uint                 `json:"entityId"`
}

type AddParamToEntityResponse struct {
	Name     string               `json:"name"`
	Type     models.ParameterType `json:"type"`
	EntityID uint                 `json:"entityId"`
}

func AddParamToEntity(c echo.Context) error {
	var req AddParamToEntityRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := config.GetDB()

	userId := uint(1)

	var entity models.Entity

	if err := db.Where("id = ?", req.EntityID).First(&entity).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if entity.CreatedByID != userId {
		return c.JSON(http.StatusUnauthorized, "You are not authorized to add parameters to this entity")
	}

	parameter := models.Parameter{
		KeyName:  req.Name,
		Type:     req.Type,
		EntityID: entity.ID,
	}

	if err := db.Create(&parameter).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res := AddParamToEntityResponse{
		Name:     parameter.KeyName,
		Type:     parameter.Type,
		EntityID: parameter.EntityID,
	}

	return c.JSON(http.StatusOK, res)
}
