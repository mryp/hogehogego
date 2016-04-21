package main

import (
	"./api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiGroup := e.Group("/api")
	apiGroup.Get("/connect", api.ConnectHandler)

	e.Run(standard.New(":8080"))
}
