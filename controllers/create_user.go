package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trinitt/utils"
)

type SignupRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsSeller  bool   `json:"is_seller"`
}

func SignupUser(c echo.Context) error {

	return utils.SendResponse(c, http.StatusOK, "User created successfully")
}
