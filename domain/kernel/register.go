package kernel

import (
	"github.com/RobyFerro/go-web-framework/cli"
	"github.com/RobyFerro/go-web-framework/domain/registers"
)

var (
	// Commands will export all registered commands
	// The following map of interfaces expose all available method that can be used by Go-Web CLI tool.
	// The map index determines the command that you've to run to for use the relative method.
	// Example: './goweb migration:up' will run '&command.MigrationUp{}' command.
	Commands = registers.CommandRegister{
		"database:seed":      &cli.Seeder{},
		"show:commands":      &cli.ShowCommands{},
		"cmd:create":         &cli.CmdCreate{},
		"controller:create":  &cli.ControllerCreate{},
		"generate:key":       &cli.GenerateKey{},
		"middleware:create":  &cli.MiddlewareCreate{},
		"migration:create":   &cli.MigrationCreate{},
		"migration:rollback": &cli.MigrateRollback{},
		"migration:up":       &cli.MigrationUp{},
		"model:create":       &cli.ModelCreate{},
		"router:show":        &cli.RouterShow{},
		"service:create":     &cli.ServiceCreate{},
		"update":             &cli.UpdateAlfred{},
		// Here is where you've to register your custom controller
	}
	// Models defines all registered models
	Models = registers.ModelRegister{}
	// Controllers defines all registered applicazion controller
	Controllers = registers.ControllerRegister{}
	// Middlewares defines all registered middlewares
	Middlewares = registers.MiddlewareRegister{}
	// Router defines all application routers
	Router []registers.RouterRegister
)
