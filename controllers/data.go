package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trinitt/config"
	"github.com/trinitt/models"
)

type Data struct {
	Value string `json:"value"`
}

type AddDataRequest struct {
	ParamID uint   `json:"paramId"`
	Dataset []Data `json:"data"`
}

func AddData(c echo.Context) error {
	var req AddDataRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := config.GetDB()

	userId := c.Get("user").(uint)

	var parameter models.Parameter

	if err := db.Preload("Entity").Where("id = ?", req.ParamID).First(&parameter).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if parameter.Entity.CreatedByID != userId {
		return c.JSON(http.StatusUnauthorized, "You are not authorized to add data to this parameter")
	}

	var data []models.Data

	for _, d := range req.Dataset {
		data = append(data, models.Data{
			Value:       d.Value,
			ParameterID: parameter.ID,
		})
	}

	if err := db.Create(&data).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Error adding data")
	}

	return c.JSON(http.StatusOK, "Data added successfully")
}
