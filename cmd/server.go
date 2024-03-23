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

	v1 := app.Group("/api/v1")
	v1.POST("/login/check", handlers.HandleLoginCheck)
	v1.POST("/register/check", handlers.HandleRegisterCheck)

	app.GET("/", handlers.HandleIndex)
	app.GET("/login", handlers.HandleLogin)
	app.GET("/register", handlers.HandleRegister)
	app.GET("/profile", handlers.HandleProfile)

	if err := app.Start(":8080"); err != nil {
		log.Fatalf("unexpected app error: %v", err)
	}
}
