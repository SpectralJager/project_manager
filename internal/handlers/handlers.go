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
	email := c.FormValue("email")
	password := c.FormValue("password")
	if email != "test@projager.io" && password != "test" {
		return c.Redirect(302, "/login")
	}
	return c.Redirect(302, "/profile")
}

func HandleRegister(c echo.Context) error {
	return c.Render(200, "register", nil)
}

func HandleRegisterCheck(c echo.Context) error {
	fmt.Println(c.FormParams())
	return c.Redirect(302, "/login")
}

func HandleProfile(c echo.Context) error {
	return c.Render(200, "profile", nil)
}
