package cli

import (
	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/interactors"
)

// GenerateKey will generate Go-Web application key in main config.yml file
type GenerateKey struct {
	entities.Command
}

// Register this command
func (c *GenerateKey) Register() {
	c.Signature = "generate:key"               // Change command signature
	c.Description = "Generate application key" // Change command description
}

// Run this command
func (c *GenerateKey) Run() {
	interactors.GenerateAppKey{}.Call()
}
