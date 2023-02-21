package cli

import (
	"fmt"
	"log"

	"github.com/RobyFerro/go-web-framework/lib/domain/entities"
	"github.com/RobyFerro/go-web-framework/lib/domain/interactors"
)

// UpdateAlfred will generate Go-Web application key in main config.yml file
type UpdateAlfred struct {
	entities.Command
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
		err := interactors.UpdateAlfred{}.Call()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Update aborted!")
	}
}
