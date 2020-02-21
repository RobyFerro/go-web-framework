package gwf

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/console"
	"github.com/RobyFerro/go-web-framework/service"
	"github.com/common-nighthawk/go-figure"
	"os"
	"reflect"
)

func GoWeb(
	args []string,
	commands CommandRegister,
	controllers ControllerRegister,
	services ServiceRegister,
	middleware interface{},
	models ModelRegister,
) {
	printCLIHeader()
	mergeCommands(commands)

	cmd := console.Commands[args[1]]
	if cmd == nil {
		fmt.Println("Command not found!")
		os.Exit(1)
	}

	v := reflect.ValueOf(cmd).MethodByName("Run")

	// Set args if exists
	if len(args) == 3 {
		v.FieldByName("Args").Set(reflect.ValueOf(args[2]))
	}

	container := service.BuildContainer(controllers, middleware, services, models)
	if err := container.Invoke(v.Interface()); err != nil {
		ProcessError(err)
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
		console.Commands[i] = c
	}
}
