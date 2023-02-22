package kernel

import (
	"fmt"
	"sync"

	"github.com/RobyFerro/go-web-framework/lib/api/cli"
	"github.com/RobyFerro/go-web-framework/lib/domain/registers"
)

var lock = &sync.Mutex{}

var baseEntities *BaseEntities

// GetBaseEntities initialize application base entities
// returns a singleton instance
func GetBaseEntities() *BaseEntities {
	if baseEntities == nil {
		lock.Lock()
		defer lock.Unlock()
		if baseEntities == nil {
			baseEntities = &BaseEntities{}
			baseEntities.registerCommands()
		} else {
			fmt.Println("BaseEntities instance already created.")
		}
	} else {
		fmt.Println("BaseEntities instance already created.")
	}

	return baseEntities
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
	c.Commands.Add("show:routes", &cli.ShowRouters{Routers: c.Router})
	c.Commands.Add("show:commands", &cli.ShowCommands{Commands: c.Commands})
}
