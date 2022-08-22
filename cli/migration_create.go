package cli

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web-framework/tool"
)

// MigrationCreate will create a new migration
// This component will create two file UP and DOWN.
// UP: Work only for migrate operation
// DOWN: Work only for rollback operation
type MigrationCreate struct {
	register.Command
}

// Register this command
func (c *MigrationCreate) Register() {
	c.Signature = "migration:create <name>"
	c.Description = "Create new migration files"
}

// Run this command
func (c *MigrationCreate) Run() {
	fmt.Println("Creating new migrations...")
	date := time.Now().Unix()
	path := tool.GetDynamicPath("database/migration")

	filenameUp := fmt.Sprintf("%s/%d_%s.up.sql", path, date, c.Args)
	filenameDown := fmt.Sprintf("%s/%d_%s.down.sql", path, date, c.Args)

	fmt.Printf("\nCreating new '%s'...\n", filenameUp)

	if err := os.WriteFile(filenameUp, []byte("/* MIGRATION UP */"), 0755); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created new up migration: %s\n", filenameUp)
	fmt.Printf("Creating new '%s'...\n", filenameDown)

	if err := os.WriteFile(filenameDown, []byte("/* MIGRATION DOWN */"), 0755); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nCreated new down migration: %s", filenameDown)
	fmt.Printf("\nDo not forget to register it!")
}
