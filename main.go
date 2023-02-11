package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/trinitt/config"
	"github.com/trinitt/routes"
	"github.com/trinitt/utils"
)

func main() {

	config.InitConfig()

	server := echo.New()
	utils.InitLogger(server)
	server.Use(middleware.CORS())
	routes.Init(server)

	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
