package main

import (
	"log"
	"project_manager/internal/handlers"
	"project_manager/internal/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var renderer = templates.NewTemplatesRenderer()

func main() {
	app := echo.New()

	app.Renderer = renderer
	app.Static("static", "public/static")

	app.Use(middleware.Logger())

	app.GET("/", handlers.HandleIndex)
	app.GET("/login", handlers.HandleLogin)
	app.POST("/login/check", handlers.HandleLoginCheck)

	if err := app.Start(":8080"); err != nil {
		log.Fatalf("unexpected app error: %v", err)
	}
}
