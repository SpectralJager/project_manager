package handlers

import (
	"context"
	"database/sql"
	_ "embed"
	"os"
	"project_manager/db"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type State struct {
	database *db.Queries
}

func InitState(dbPath string, schemaPath string) (*State, error) {
	ctx := context.Background()
	d, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, err
	}

	_, err = d.ExecContext(ctx, string(schema))
	if err != nil {
		return nil, err
	}

	return &State{
		database: db.New(d),
	}, nil
}

func (s *State) IndexRouteHandler(c echo.Context) error {
	users, err := s.database.ListUsers(c.Request().Context())
	if err != nil {
		return c.JSON(403, map[string]any{
			"error": "can't get list of user",
		})
	}
	return c.JSON(200, users)
}

func (s *State) AddUserRouterHandler(c echo.Context) error {
	var user db.CreateUserParams
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(403, err)
	}
	err = s.database.CreateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(403, err)
	}
	return c.JSON(200, map[string]any{"msg": "success"})

}

func (s *State) GetUserRouteHandler(c echo.Context) error {
	username := c.QueryParam("username")
	if username == "" {
		return c.JSON(403, map[string]any{"error": "expect username param"})
	}
	user, err := s.database.GetUser(c.Request().Context(), username)
	if err != nil {
		return c.JSON(403, map[string]any{"error": "can't get username with name " + username})
	}
	return c.JSON(200, user)
}
