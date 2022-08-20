package foundation

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/RobyFerro/dig"
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/common-nighthawk/go-figure"
)

type BaseEntities struct {
	Controllers       register.ControllerRegister
	Commands          register.CommandRegister
	Services          register.ServiceRegister
	SingletonServices register.ServiceRegister
	CommandServices   register.ServiceRegister
	Middlewares       register.MiddlewareRegister
	Models            register.ModelRegister
	Router            []register.HTTPRouter
}

// Start will run the HTTP web server
func Start(e BaseEntities, c kernel.ServerConf) {
	startup(e)
	kernel.RunServer(c, e.Router)
}

// StartCommand method runs specific CLI command
func StartCommand(args []string, e BaseEntities) {
	startup(e)

	c := kernel.BuildCommandContainer()
	cmd := kernel.Commands[args[0]]
	if cmd == nil {
		fmt.Println("Command not found!")
		os.Exit(1)
	}

	rc := reflect.ValueOf(cmd)
	if len(args) == 2 {
		reflect.Indirect(rc).FieldByName("Args").SetString(args[1])
	}

	err := dig.GroupInvoke(rc.MethodByName("Run").Interface(), c)
	if err != nil {
		log.Fatal(err)
	}
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

	mergeSingletonServices(entities.SingletonServices)
	mergeCommandServices(entities.CommandServices)
}

// Merge singleton services with defaults
func mergeSingletonServices(services []interface{}) {
	for _, s := range services {
		kernel.SingletonServices = append(kernel.SingletonServices, s)
	}
}

// MergeCommands will merge system command with customs
func mergeCommands(commands register.CommandRegister) {
	for i, c := range commands {
		kernel.Commands[i] = c
	}
}

// MergeCommands will merge system command with customs
func mergeMiddleware(mw register.MiddlewareRegister) {
	for i, c := range mw {
		kernel.Middlewares[i] = c
	}
}

// MergeCommands will merge system command with customs
func mergeCommandServices(services []interface{}) {
	for _, s := range services {
		kernel.CommandServices = append(kernel.CommandServices, s)
	}
}
