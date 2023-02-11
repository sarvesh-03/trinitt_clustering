package controllers

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/trinitt/config"
	"github.com/trinitt/models"
	"github.com/trinitt/utils"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func LoginUser(c echo.Context) error {
	db := config.GetDB()

	var login LoginRequest

	if err := c.Bind(&login); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	// Get user from database
	var user models.User

	if err := db.Where("email = ?", login.Email).First(&user).Error; err != nil {
		return c.String(http.StatusUnauthorized, "Unauthorized - 1")
	}

	fmt.Println(user)

	fmt.Println(user.Password, "user password")

	if !utils.VerifyPassword(user.Password, login.Password) {
		return c.String(http.StatusUnauthorized, "Unauthorized - 2")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func RegisterUser(c echo.Context) error {
	db := config.GetDB()

	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Bad request - 1")
	}

	if req.Password != req.ConfirmPassword {
		return c.String(http.StatusBadRequest, "Bad request - 2")
	}

	// Check if user already exists
	var user models.User

	if err := db.Where("email = ?", req.Email).First(&user).Error; err == nil {
		return c.String(http.StatusBadRequest, "Bad request - 3")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error - 1")
	}

	fmt.Println(hashedPassword, "hashed password")

	user.Password = hashedPassword
	user.Email = req.Email

	if err := db.Create(&user).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error - 2")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error - 3, Try logging in")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
