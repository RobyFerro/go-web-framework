package cli

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

type RouterShow struct {
	register.Command
}

func (c *RouterShow) Register() {
	c.Signature = "router:show"
	c.Description = "Show all available routes"
}

func (c *RouterShow) Run(routes []register.HTTPRouter) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"PATH", "METHOD", "DESCRIPTION", "MIDDLEWARES"})
	for _, router := range routes {
		parseRoutes(router.Route, table, nil)
		parseGroups(router.Groups, table)
	}

	table.Render()
}

func parseGroups(groups []register.Group, table *tablewriter.Table) {
	for _, group := range groups {
		parseRoutes(group.Routes, table, &group.Prefix)
	}
}

func parseRoutes(routes []register.Route, table *tablewriter.Table, prefix *string) {
	for _, r := range routes {
		middlewares := getMiddlewareString(&r.Middleware)
		if prefix != nil {
			r.Path = fmt.Sprintf("%s%s", *prefix, r.Path)
		}

		table.Append([]string{r.Path, r.Method, r.Description, middlewares})
	}
}

func getMiddlewareString(middlewares *[]register.Middleware) string {
	var list []string
	for _, m := range *middlewares {
		list = append(list, m.GetName())
	}

	return strings.Join(list, ",")
}
