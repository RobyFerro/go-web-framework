package foundation

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/common-nighthawk/go-figure"
	"log"
	"os"
	"reflect"
)

// Start method will start the main Go-Web HTTP Server.
func Start(args []string, cm register.CommandRegister, c register.ControllerRegister, s register.ServiceRegister, mw interface{},
	m register.ModelRegister) {
	printCLIHeader()
	registerBaseEntities(c, m, s, cm, mw)

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

	// Build service container.
	// This container will used to invoke the requested command.
	container := kernel.BuildContainer()
	if err := container.Invoke(rc.MethodByName("Run").Interface()); err != nil {
		log.Fatal(err)
	}
}

// Register base entities in Go-Web kernel
// This method will register: Controllers, Models, CLI commands, Services and middleware
func registerBaseEntities(c register.ControllerRegister, m register.ModelRegister, s register.ServiceRegister,
	cm register.CommandRegister, mw interface{}) {
	kernel.Controllers = c
	kernel.Middleware = mw
	kernel.Models = m

	mergeCommands(cm)
	bindServices(s.List)
}

// Merge custom services with defaults
func bindServices(services []interface{}) {
	for _, s := range services {
		kernel.Services.List = append(kernel.Services.List, s)
	}
}

// Print Go-Web CLI header
func printCLIHeader() {
	myFigure := figure.NewFigure("Go-Web", "graffiti", true)
	myFigure.Print()

	fmt.Println("Go-Web CLI tool - Author: roberto.ferro@ikdev.eu")
}

// MergeCommands will merge system command with customs
func mergeCommands(commands register.CommandRegister) {
	for i, c := range commands.List {
		kernel.Commands.List[i] = c
	}
}
