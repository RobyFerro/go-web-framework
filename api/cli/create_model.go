package cli

import (
	"fmt"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/interactors"
	"github.com/RobyFerro/go-web-framework/infrastructure/services"
)

// CreateModel handles new CLI command creation
type CreateModel struct {
	entities.Command
}

// Register this command
func (c *CreateModel) Register() {
	c.Signature = "cmd:create <name>"
	c.Description = "Create new command"
}

// Run create command
func (c *CreateModel) Run() {
	fmt.Println("Creating new application model...")
	useCase := interactors.GenerateNewApplicationElement{
		Blueprint:   "model",
		Command:     c.Command,
		Service:     services.CommandServiceImpl{},
		Destination: "database/model",
	}

	useCase.Call()
	fmt.Println("Success!")
}
