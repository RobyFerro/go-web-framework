package cli

import (
	"fmt"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/interactors"
	"github.com/RobyFerro/go-web-framework/infrastructure/services"
)

// CreateMiddleware handles new CLI command creation
type CreateMiddleware struct {
	entities.Command
}

// Register this command
func (c *CreateMiddleware) Register() {
	c.Signature = "cmd:create <name>"
	c.Description = "Create new command"
}

// Run create command
func (c *CreateMiddleware) Run() {
	fmt.Println("Creating new application middleware...")
	useCase := interactors.GenerateNewApplicationElement{
		Blueprint:   "middleware",
		Command:     c.Command,
		Service:     services.CommandServiceImpl{},
		Destination: "app/http/middleware",
	}

	useCase.Call()
	fmt.Println("Success!")
}
