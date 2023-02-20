package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/go-git/go-git/v5"
)

// ServiceCreate will generate Go-Web application key in main config.yml file
type ServiceCreate struct {
	entities.Command
}

// Register this command
func (c *ServiceCreate) Register() {
	c.Signature = "service:create [service-name]" // Change command signature
	c.Description = "Create new Go-Web service"   // Change command description
}

// Run this command
func (c *ServiceCreate) Run() {
	if len(c.Args) == 0 {
		c.Help()
		return
	}

	fmt.Printf("Creating service %s...\n", c.Args)
	if err := c.clone(c.Args); err != nil {
		log.Fatalf("Error: %s", err)
	}

	if err := c.resetGit(); err != nil {
		log.Fatalf("Error: %s", err)
	}

	if err := c.update(); err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Println("Service created successfully!")
}

// Help will show help for this command
func (c *ServiceCreate) Help() {
	log.Println("Usage: create-service [service-name]")
}

// Clones Go-Web repository in destination folder
func (c *ServiceCreate) clone(destination string) error {
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:      "https://github.com/RobyFerro/go-web.git",
		Progress: nil,
	})

	return err
}

// Reset git repository
func (c *ServiceCreate) resetGit() error {
	path := fmt.Sprintf("%s/.git", c.Args)
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = c.Args
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// Updates Go-Web Framework to the latest minor version
func (c *ServiceCreate) update() error {
	cmd := exec.Command("go", "get", "-u", "github.com/RobyFerro/go-web-framework")
	cmd.Dir = c.Args

	return cmd.Run()
}
