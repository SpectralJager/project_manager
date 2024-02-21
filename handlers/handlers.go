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

	err = d.Ping()
	if err != nil {
		return nil, err
	}

	return &State{
		database: db.New(d),
	}, nil
}

func (s *State) IndexRouteHandler(c echo.Context) error {
	return c.Render(200, "index", nil)
}
