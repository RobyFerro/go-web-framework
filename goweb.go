package gwf

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
	"reflect"
)

// Start method will start the main Go-Web HTTP Server.
func Start(args []string, cm CommandRegister, c ControllerRegister, s ServiceRegister, mw interface{}, m ModelRegister) {
	printCLIHeader()
	registerBaseEntities(c, m, s, cm, mw)

	cmd := Commands[args[0]]
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
		ProcessError(err)
	}
}

// Register base entities in Go-Web kernel
// This method will register: Controllers, Models, CLI commands, Services and middleware
func registerBaseEntities(c ControllerRegister, m ModelRegister, s ServiceRegister, cm CommandRegister, mw interface{}) {
	Controllers = c
	Middleware = mw
	Models = m

	mergeCommands(cm)
	bindServices(s)
}

// Merge custom services with defaults
func bindServices(services []interface{}) {

	for _, s := range services {
		Services = append(Services, s)
	}
}

// Print Go-Web CLI header
func printCLIHeader() {
	myFigure := figure.NewFigure("Go-Web", "graffiti", true)
	myFigure.Print()

	fmt.Println("Go-Web CLI tool - Author: roberto.ferro@ikdev.eu")
}

// MergeCommands will merge system command with customs
func mergeCommands(commands CommandRegister) {
	for i, c := range commands {
		Commands[i] = c
	}
}
