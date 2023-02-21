package kernel

import (
	"github.com/RobyFerro/go-web-framework/lib/domain/interactors"
	"github.com/RobyFerro/go-web-framework/lib/infrastructure/services"
)

// Start will run the HTTP web server
func Start(e *BaseEntities) {
	//startup(e)

	router := interactors.GetHTTPRouter{
		Service: services.RouterServiceImpl{
			Controllers: e.Controllers,
		},
		Register: e.Router,
	}.Call()

	config := interactors.GetAppConfig{}.Call()
	server := interactors.GetHTTPServer{
		Config: config,
		Router: router,
	}.Call()

	interactors.StartHTTPServer{
		Server: *server,
		Config: config,
	}.Call()
}

// StartCommand method runs specific CLI command
// func StartCommand(args []string, e BaseEntities) {
// 	startup(e)

// 	cmd := Commands[args[0]]
// 	if cmd == nil {
// 		fmt.Println("Command not found!")
// 		os.Exit(1)
// 	}

// 	rc := reflect.ValueOf(cmd)
// 	if len(args) == 2 {
// 		reflect.Indirect(rc).FieldByName("Args").SetString(args[1])
// 	}

// 	rc.MethodByName("Run").Interface()
// }

// func startup(e config.BaseEntities) {
// 	myFigure := figure.NewFigure("Go-Web", "graffiti", true)
// 	myFigure.Print()
// 	fmt.Println("Go-Web CLI tool - Author: roberto.ferro@ikdev.it")

// 	RegisterBaseEntities(e)
// }

// RegisterBaseEntities base entities in Go-Web kernel
// This method will register: Controllers, Models, CLI commands, Services and middleware
// func RegisterBaseEntities(entities config.BaseEntities) {
// 	kernel.Controllers = entities.Controllers
// 	kernel.Middlewares = entities.Middlewares
// 	kernel.Models = entities.Models
// 	kernel.Router = entities.Router

// 	mergeCommands(entities.Commands)
// 	mergeMiddleware(entities.Middlewares)
// }

// MergeCommands will merge system command with customs
// func mergeCommands(commands registers.CommandRegister) {
// 	for i, c := range commands {
// 		kernel.Commands[i] = c
// 	}
// }

// MergeCommands will merge system command with customs
// func mergeMiddleware(mw registers.MiddlewareRegister) {
// 	for i, c := range mw {
// 		kernel.Middlewares[i] = c
// 	}
// }
