package foundation

import (
	"fmt"
	"os"
	"reflect"

	"github.com/RobyFerro/go-web-framework/domain/interactors"
	"github.com/RobyFerro/go-web-framework/domain/kernel"
	"github.com/RobyFerro/go-web-framework/domain/registers"
	"github.com/RobyFerro/go-web-framework/infrastructure/services"
	"github.com/common-nighthawk/go-figure"
)

type BaseEntities struct {
	Controllers registers.ControllerRegister
	Commands    registers.CommandRegister
	Middlewares registers.MiddlewareRegister
	Models      registers.ModelRegister
	Router      []registers.RouterRegister
}

// Start will run the HTTP web server
func Start(e BaseEntities) {
	startup(e)
	routerService := services.RouterServiceImpl{}

	router := interactors.GetHTTPRouter(routerService, e.Router)
	conf := interactors.GetAppConfig()
	server := interactors.GetHTTPServer(conf, router)

	interactors.StartHTTPServer(*server, conf)
}

// StartCommand method runs specific CLI command
func StartCommand(args []string, e BaseEntities) {
	startup(e)

	cmd := kernel.Commands[args[0]]
	if cmd == nil {
		fmt.Println("Command not found!")
		os.Exit(1)
	}

	rc := reflect.ValueOf(cmd)
	if len(args) == 2 {
		reflect.Indirect(rc).FieldByName("Args").SetString(args[1])
	}

	rc.MethodByName("Run").Interface()
}

func startup(e BaseEntities) {
	myFigure := figure.NewFigure("Go-Web", "graffiti", true)
	myFigure.Print()
	fmt.Println("Go-Web CLI tool - Author: roberto.ferro@ikdev.it")

	RegisterBaseEntities(e)
}

// RegisterBaseEntities base entities in Go-Web kernel
// This method will register: Controllers, Models, CLI commands, Services and middleware
func RegisterBaseEntities(entities BaseEntities) {
	kernel.Controllers = entities.Controllers
	kernel.Middlewares = entities.Middlewares
	kernel.Models = entities.Models
	kernel.Router = entities.Router

	mergeCommands(entities.Commands)
	mergeMiddleware(entities.Middlewares)
}

// MergeCommands will merge system command with customs
func mergeCommands(commands registers.CommandRegister) {
	for i, c := range commands {
		kernel.Commands[i] = c
	}
}

// MergeCommands will merge system command with customs
func mergeMiddleware(mw registers.MiddlewareRegister) {
	for i, c := range mw {
		kernel.Middlewares[i] = c
	}
}
