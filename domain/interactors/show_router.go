package interactors

import (
	"fmt"
	"os"
	"strings"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/registers"
	"github.com/olekukonko/tablewriter"
)

// ShowRouter prints applications routers
type ShowRouter struct {
	Routers registers.RouterRegister
}

// Call executes show router interactor
func (c ShowRouter) Call() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"PATH", "METHOD", "DESCRIPTION", "MIDDLEWARES"})
	for _, router := range c.Routers {
		c.parseRoutes(router.Route, table, nil)
		c.parseGroups(router.Groups, table)
	}

	table.Render()
}

func (c ShowRouter) parseRoutes(routes []entities.Route, table *tablewriter.Table, prefix *string) {
	for _, r := range routes {
		middlewares := c.getMiddlewareString(&r.Middleware)
		if prefix != nil {
			r.Path = fmt.Sprintf("%s%s", *prefix, r.Path)
		}

		table.Append([]string{r.Path, r.Method, r.Description, middlewares})
	}
}

func (c ShowRouter) parseGroups(groups []entities.Group, table *tablewriter.Table) {
	for _, group := range groups {
		c.parseRoutes(group.Routes, table, &group.Prefix)
	}
}

func (c ShowRouter) getMiddlewareString(middlewares *[]entities.Middleware) string {
	var list []string
	for _, m := range *middlewares {
		list = append(list, m.GetName())
	}

	return strings.Join(list, ",")
}
