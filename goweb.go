package gwf

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/foundation"
	"github.com/RobyFerro/go-web-framework/helper"
	"github.com/RobyFerro/go-web-framework/types"
	"github.com/common-nighthawk/go-figure"
	"os"
	"reflect"
)

// Start method will start the main Go-Web HTTP Server.
func Start(args []string, cm types.CommandRegister, c types.ControllerRegister, s types.ServiceRegister, mw interface{}, m types.ModelRegister) {
	printCLIHeader()
	registerBaseEntities(c, m, s, cm, mw)

	cmd := foundation.Commands.List[args[0]]
	if cmd == nil {
		fmt.Println("Command not found!")
		os.Exit(1)
	}

	rc := reflect.ValueOf(cmd)
	// Set args if exists
	if len(args) == 2 {
		reflect.Indirect(rc).FieldByName("Args").SetString(args[1])
	}

	// Build service container.
	// This container will used to invoke the requested command.
	container := BuildContainer()
	if err := container.Invoke(rc.MethodByName("Run").Interface()); err != nil {
		helper.ProcessError(err)
	}
}

// Register base entities in Go-Web kernel
// This method will register: Controllers, Models, CLI commands, Services and middleware
func registerBaseEntities(c types.ControllerRegister, m types.ModelRegister, s types.ServiceRegister, cm types.CommandRegister, mw interface{}) {
	foundation.Controllers = c
	foundation.Middleware = mw
	foundation.Models = m

	mergeCommands(cm)
	bindServices(s.List)
}

// Merge custom services with defaults
func bindServices(services []interface{}) {

	for _, s := range services {
		foundation.Services.List = append(foundation.Services.List, s)
	}
}

// Print Go-Web CLI header
func printCLIHeader() {
	myFigure := figure.NewFigure("Go-Web", "graffiti", true)
	myFigure.Print()

	fmt.Println("Go-Web CLI tool - Author: roberto.ferro@ikdev.eu")
}

// MergeCommands will merge system command with customs
func mergeCommands(commands types.CommandRegister) {
	for i, c := range commands.List {
		foundation.Commands.List[i] = c
	}
}
