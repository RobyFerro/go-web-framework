package command

import (
	"fmt"
	"io/ioutil"
	"time"

	gwf "github.com/RobyFerro/go-web-framework"
)

// MigrationCreate will create a new migration
// This component will create two file UP and DOWN.
// UP: Work only for migrate operation
// DOWN: Work only for rollback operation
type MigrationCreate struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *MigrationCreate) Register() {
	c.Signature = "migration:create <name>"
	c.Description = "Create new migration files"
}

// Run this command
func (c *MigrationCreate) Run() {
	date := time.Now().Unix()
	path := gwf.GetDynamicPath("database/migration")

	filenameUp := fmt.Sprintf("%s/%d_%s.up.sql", path, date, c.Args)
	filenameDown := fmt.Sprintf("%s/%d_%s.down.sql", path, date, c.Args)

	fmt.Printf("\nCreating new '%s'...\n", filenameUp)

	if err := ioutil.WriteFile(filenameUp, []byte("/* MIGRATION UP */"), 0755); err != nil {
		gwf.ProcessError(err)
	}

	fmt.Printf("Created new up migration: %s\n", filenameUp)
	fmt.Printf("Creating new '%s'...\n", filenameDown)

	if err := ioutil.WriteFile(filenameDown, []byte("/* MIGRATION DOWN */"), 0755); err != nil {
		gwf.ProcessError(err)
	}

	fmt.Printf("Created new down migration: %s\n", filenameDown)
	fmt.Printf("Do not forget to register it!")
}
