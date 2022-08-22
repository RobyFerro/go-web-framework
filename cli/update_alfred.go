package cli

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/RobyFerro/go-web-framework/register"
)

// GenerateKey will generate Go-Web application key in main config.yml file
type UpdateAlfred struct {
	register.Command
}

// Register this command
func (c *UpdateAlfred) Register() {
	c.Signature = "update"                                         // Change command signature
	c.Description = "Updates Alfred CLI with your custom commands" // Change command description
}

// Run this command
func (c *UpdateAlfred) Run() {
	fmt.Println("Are you sure you want to update Alfred? (y/n)")
	var answer string
	fmt.Scanln(&answer)
	if answer == "y" {
		if err := c.update_alfred(); err != nil {
			log.Fatalf("Error: %s", err)
		}
	} else {
		fmt.Println("Update aborted!")
	}
}

// Help will show help for this command
func (c *UpdateAlfred) Help() {
	log.Println("Usage: create-service [service-name]")
}

func (c *UpdateAlfred) update_alfred() error {
	cmd := exec.Command("go", "install", "./cmd/alfred/...")
	return cmd.Run()
}
