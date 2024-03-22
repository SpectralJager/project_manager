package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func HandleIndex(c echo.Context) error {
	return c.Render(200, "index", nil)
}

func HandleLogin(c echo.Context) error {
	return c.Render(200, "login", nil)
}

func HandleLoginCheck(c echo.Context) error {
	fmt.Println(c.FormParams())
	return c.Redirect(302, "/")
}
