package cli

import (
	"fmt"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/interactors"
	"github.com/RobyFerro/go-web-framework/infrastructure/services"
)

// CreateNewApplication will generate Go-Web application key in main config.yml file
type CreateNewApplication struct {
	entities.Command
}

// Register this command
func (c *CreateNewApplication) Register() {
	c.Signature = "service:create [service-name]" // Change command signature
	c.Description = "Create new Go-Web service"   // Change command description
}

// Run this command
func (c *CreateNewApplication) Run() {
	fmt.Printf("Creating service %s...\n", c.Args)
	interactors.GenerateNewApp{
		Service:     services.GitServicesImpl{},
		Repository:  "github.com/RobyFerro/go-web-framework",
		Destination: c.Args,
	}.Call()

	fmt.Println("Service created successfully!")
}
