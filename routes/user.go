package routes

import (
	controllers "github.com/trinitt/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	user := e.Group("/user")
	
	user.GET("/kafka", controllers.SignupUser)
}
