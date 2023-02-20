package cli

import (
	"fmt"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/interactors"
	"github.com/RobyFerro/go-web-framework/infrastructure/services"
)

// CreateController handles new CLI command creation
type CreateController struct {
	entities.Command
}

// Register this command
func (c *CreateController) Register() {
	c.Signature = "cmd:create <name>"
	c.Description = "Create new command"
}

// Run create command
func (c *CreateController) Run() {
	fmt.Println("Creating new application controller...")
	useCase := interactors.GenerateNewApplicationElement{
		Blueprint:   "controller",
		Command:     c.Command,
		Service:     services.CommandServiceImpl{},
		Destination: "app/http/controller",
	}

	useCase.Call()
	fmt.Println("Success!")
}
