package routes

import (


	"github.com/labstack/echo/v4"
	"github.com/trinitt/controllers"
)



func Init(e *echo.Echo) {
	api := e.Group("/api")
	UserRoutes(api)
	api.GET("/kafka", controllers.SignupUser)
	api.GET("/consume", controllers.Consume)

}
