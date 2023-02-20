package cli

import (
	"fmt"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/interactors"
)

// MigrationCreate will create a new migration
// This component will create two file UP and DOWN.
// UP: Work only for migrate operation
// DOWN: Work only for rollback operation
type MigrationCreate struct {
	entities.Command
}

// Register this command
func (c *MigrationCreate) Register() {
	c.Signature = "migration:create <name>"
	c.Description = "Create new migration files"
}

// Run this command
func (c *MigrationCreate) Run() {
	fmt.Println("Creating new migrations...")
	interactors.GenerateNewMigration{
		Name: c.Args,
	}.Call()

	fmt.Println("Complete!")
}
