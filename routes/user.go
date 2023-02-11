package routes

import (
	controllers "github.com/trinitt/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	user := e.Group("/user")

	user.POST("/signup", controllers.RegisterUser)
	user.POST("/login", controllers.LoginUser)
}
