package main

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// ShowRoute will shows all routes registerer in your Go-Web application
type ShowRoute struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *ShowRoute) Register() {
	c.Signature = "show:route"
	c.Description = "Show active Go-Web routes"
}

// Run this command
func (c *ShowRoute) Run() {
	var data [][]string
	routes, err := ConfigurationWeb()
	if err != nil {
		ProcessError(err)
	}

	// Parse single route
	showSingleRoute(routes.Routes, &data)

	// Parse groups
	showGroupRoutes(routes.Groups, &data)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"METHOD", "ACTION", "PREFIX", "PATH", "MIDDLEWARE", "DESCRIPTION"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}

// Show single routes
func showSingleRoute(routes map[string]Route, data *[][]string) {
	for _, r := range routes {
		*data = append(*data, []string{
			r.Method,
			r.Action,
			r.Prefix,
			r.Path,
			strings.ToLower(strings.Join(r.Middleware, ", ")),
			r.Description,
		})
	}
}

// Show groups
func showGroupRoutes(routes map[string]Group, data *[][]string) {
	for _, g := range routes {
		var middleware []string
		middleware = append(middleware, g.Middleware...)
		for _, gr := range g.Routes {
			middleware = append(middleware, gr.Middleware...)
			*data = append(*data, []string{
				gr.Method,
				gr.Action,
				g.Prefix,
				gr.Path,
				strings.ToLower(strings.Join(middleware, ", ")),
				gr.Description,
			})
		}
	}
}
