package cli

import (
	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/interactors"
	"github.com/RobyFerro/go-web-framework/domain/registers"
)

// ShowRouters prints all routers to the stdout
type ShowRouters struct {
	entities.Command
	Routers registers.RouterRegister
}

// Register register show router command
func (c *ShowRouters) Register() {
	c.Signature = "router:show"
	c.Description = "Show all available routes"
}

// Run show router command
func (c *ShowRouters) Run() {
	interactors.ShowRouter{Routers: c.Routers}.Call()
}
