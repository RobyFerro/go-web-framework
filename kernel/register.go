package kernel

import (
	"github.com/RobyFerro/go-web-framework/cli"
	"github.com/RobyFerro/go-web-framework/register"
)

var (
	// Commands will export all registered commands
	// The following map of interfaces expose all available method that can be used by Go-Web CLI tool.
	// The map index determines the command that you've to run to for use the relative method.
	// Example: './goweb migration:up' will run '&command.MigrationUp{}' command.
	Commands = register.CommandRegister{
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
		// Here is where you've to register your custom controller
	}
	SingletonServices = register.ServiceRegister{
		RetrieveAppConf,
		CreateSessionStore,
	}
	CommandServices = register.ServiceRegister{}
	Models          = register.ModelRegister{}
	Controllers     = register.ControllerRegister{}
	Middlewares     = register.MiddlewareRegister{}
	Router          []register.HTTPRouter
)
