package cli

import (
	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/interactors"
	"github.com/RobyFerro/go-web-framework/domain/registers"
)

// ShowCommands will show all registered commands
type ShowCommands struct {
	entities.Command
	Commands registers.CommandRegister
}

// Register this command
func (c *ShowCommands) Register() {
	c.Signature = "show:commands"
	c.Description = "Show Go-Web commands list"
}

// Run this command
func (c *ShowCommands) Run() {
	interactors.ShowCommands{Commands: c.Commands}.Call()
}
