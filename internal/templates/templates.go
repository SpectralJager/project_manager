package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplatesRenderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseFiles(
			"./internal/templates/base.html",
			"./internal/templates/index.html",
			"./internal/templates/login.html",
			"./internal/templates/register.html",
			"./internal/templates/profile.html",
		)),
	}
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
