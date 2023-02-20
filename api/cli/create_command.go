package cli

import (
	"fmt"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/interactors"
	"github.com/RobyFerro/go-web-framework/infrastructure/services"
)

// CreateCommand handles new CLI command creation
type CreateCommand struct {
	entities.Command
}

// Register this command
func (c *CreateCommand) Register() {
	c.Signature = "cmd:create <name>"
	c.Description = "Create new command"
}

// Run create command
func (c *CreateCommand) Run() {
	fmt.Println("Creating new CLI command...")
	useCase := interactors.GenerateNewApplicationElement{
		Blueprint:   "command",
		Command:     c.Command,
		Service:     services.CommandServiceImpl{},
		Destination: "",
	}

	useCase.Call()
	fmt.Println("Success!")
}
