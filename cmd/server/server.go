package main

import (
	"log"
	"project_manager/handlers"
	"project_manager/templates"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	server.File("/static/main.css", "./static/main.css")

	state, err := handlers.InitState("./db/tmp.db", "./sql/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	server.Renderer = &templates.Template{
		Templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	server.GET("/", state.IndexRouteHandler)
	server.GET("/user", state.GetUserRouteHandler)
	server.POST("/user", state.AddUserRouterHandler)

	err = server.Start("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
