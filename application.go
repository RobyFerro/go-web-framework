package foundation

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/common-nighthawk/go-figure"
	"os"
	"reflect"
)

type BaseEntities struct {
	Controllers       register.ControllerRegister
	Commands          register.CommandRegister
	Services          register.ServiceRegister
	SingletonServices register.ServiceRegister
	Middlewares       interface{}
	Models            register.ModelRegister
}

// Start method will start the main Go-Web HTTP Server.
func Start(args []string, entities BaseEntities) {
	myFigure := figure.NewFigure("Go-Web", "graffiti", true)
	myFigure.Print()

	fmt.Println("Go-Web CLI tool - Author: roberto.ferro@ikdev.eu")

	registerBaseEntities(entities)

	cmd := kernel.Commands.List[args[0]]
	if cmd == nil {
		fmt.Println("Command not found!")
		os.Exit(1)
	}

	rc := reflect.ValueOf(cmd)
	// Set args if exists
	if len(args) == 2 {
		reflect.Indirect(rc).FieldByName("Args").SetString(args[1])
	}

	singletonIOC = kernel.BuildSingletonContainer()
	rc.MethodByName("Run").Call([]reflect.Value{})
}

// Register base entities in Go-Web kernel
// This method will register: Controllers, Models, CLI commands, Services and middleware
func registerBaseEntities(entities BaseEntities) {
	kernel.Controllers = entities.Controllers
	kernel.Middleware = entities.Middlewares
	kernel.Models = entities.Models

	mergeCommands(entities.Commands)
	mergeServices(entities.Services.List)
	mergeSingletonServices(entities.SingletonServices.List)
}

// Merge services with defaults
func mergeServices(services []interface{}) {
	for _, s := range services {
		kernel.Services.List = append(kernel.Services.List, s)
	}
}

// Merge singleton services with defaults
func mergeSingletonServices(services []interface{}) {
	for _, s := range services {
		kernel.SingletonServices.List = append(kernel.SingletonServices.List, s)
	}
}

// MergeCommands will merge system command with customs
func mergeCommands(commands register.CommandRegister) {
	for i, c := range commands.List {
		kernel.Commands.List[i] = c
	}
}
