package config

import (
	"github.com/RobyFerro/go-web-framework/api/cli"
	"github.com/RobyFerro/go-web-framework/domain/registers"
)

// BaseEntities declare application base entities
type BaseEntities struct {
	Controllers registers.ControllerRegister
	Commands    registers.CommandRegister
	Middlewares registers.MiddlewareRegister
	Models      registers.ModelRegister
	Router      []registers.RouterRegister
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
}
