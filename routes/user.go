package routes

import (

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	e.Group("/user")
	
}
