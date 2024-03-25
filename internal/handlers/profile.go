package handlers

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProfileProject struct {
	Id     int
	Name   string
	Status string
	Tasks  int
	Owner  string
}

var projects = []ProfileProject{
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "active",
		Tasks:  12,
		Owner:  "you",
	},
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "active",
		Tasks:  12,
		Owner:  "you",
	},
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "active",
		Tasks:  12,
		Owner:  "you",
	},
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "active",
		Tasks:  12,
		Owner:  "John Smith",
	},
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "done",
		Tasks:  12,
		Owner:  "you",
	},
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "done",
		Tasks:  12,
		Owner:  "John Smith",
	},
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "active",
		Tasks:  12,
		Owner:  "John Smith",
	},
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "active",
		Tasks:  12,
		Owner:  "you",
	},
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "done",
		Tasks:  12,
		Owner:  "you",
	},
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "active",
		Tasks:  12,
		Owner:  "John Smith",
	},
	{
		Id:     0,
		Name:   "Lorem ipsum dore sid amen? core el do sirom",
		Status: "done",
		Tasks:  12,
		Owner:  "you",
	},
}

func HandleProfileOverview(c echo.Context) error {
	return c.Render(200, "profile_overview", echo.Map{
		"Projects": projects,
	})
}

func HandleProfileOverviewSearch(c echo.Context) error {
	fmt.Println(c.FormParams())
	prjcts := make([]ProfileProject, 0, len(projects))
	title := c.FormValue("query")
	status := c.FormValue("status")
	owner := c.FormValue("owner")
	for _, p := range projects {
		if strings.Contains(p.Name, title) &&
			strings.Contains(p.Status, status) &&
			strings.Contains(p.Owner, owner) {
			prjcts = append(prjcts, p)
		}
	}
	return c.Render(200, "profile_overview_data", prjcts)
}
