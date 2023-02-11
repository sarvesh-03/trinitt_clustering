package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/trinitt/config"
	"github.com/trinitt/controllers"
	"github.com/trinitt/routes"
	"github.com/trinitt/utils"
)

func main() {

	config.InitConfig()
	config.CreateProducer()
	config.CreateConsumer()
	server := echo.New()
	utils.InitLogger(server)
	server.Use(middleware.CORS())
	routes.Init(server)
	controllers.InitSetup()
	controllers.Consume()

	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
