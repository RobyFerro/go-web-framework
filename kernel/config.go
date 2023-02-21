package kernel

import (
	"fmt"
	"sync"

	"github.com/RobyFerro/go-web-framework/lib/api/cli"
	"github.com/RobyFerro/go-web-framework/lib/domain/registers"
)

var lock = &sync.Mutex{}

var entities *BaseEntities

// GetBaseEntities initialize application base entities
// returns a singleton instance
func GetBaseEntities() *BaseEntities {
	if entities == nil {
		lock.Lock()
		defer lock.Unlock()
		if entities == nil {
			entities = &BaseEntities{}
			entities.registerCommands()
		} else {
			fmt.Println("BaseEntities instance already created.")
		}
	} else {
		fmt.Println("BaseEntities instance already created.")
	}

	return entities
}

// BaseEntities declare application base entities
type BaseEntities struct {
	Controllers registers.ControllerRegister
	Commands    registers.CommandRegister
	Middlewares registers.MiddlewareRegister
	Models      registers.ModelRegister
	Router      registers.RouterRegister
}

// Register base entities
func (c BaseEntities) Register() {
	c.registerCommands()
}

func (c BaseEntities) registerCommands() {
	c.Commands.Add("create:command", &cli.CreateCommand{})
	c.Commands.Add("create:controller", &cli.CreateCommand{})
	c.Commands.Add("create:middleware", &cli.CreateMiddleware{})
	c.Commands.Add("create:model", &cli.CreateModel{})
	c.Commands.Add("create:app-key", &cli.GenerateKey{})
	c.Commands.Add("create:migration", &cli.MigrationCreate{})
	c.Commands.Add("show:routes", &cli.ShowRouters{Routers: c.Router})
	c.Commands.Add("show:commands", &cli.ShowCommands{Commands: c.Commands})
}
