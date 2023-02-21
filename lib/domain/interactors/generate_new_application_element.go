package interactors

import (
	"github.com/RobyFerro/go-web-framework/lib/domain/entities"
	"github.com/RobyFerro/go-web-framework/lib/domain/services"
)

// GenerateNewApplicationElement interactors is used to generate new application controller
type GenerateNewApplicationElement struct {
	Blueprint   string
	Command     entities.Command
	Service     services.CommandService
	Destination string
}

// Call executes usecase logic
func (c GenerateNewApplicationElement) Call() {
	commandName := c.Service.SnakeToCamelCase(c.Command.Args)
	blueprint := c.Service.ReadBlueprint(c.Blueprint)

	content := c.Service.ReplaceContent(commandName, blueprint)
	c.Service.Create(content, c.Destination, c.Command.Args)
}
