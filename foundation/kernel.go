package foundation

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/cli"
	"github.com/RobyFerro/go-web-framework/helper"
	http2 "github.com/RobyFerro/go-web-framework/http"
	"github.com/RobyFerro/go-web-framework/service"
	"github.com/RobyFerro/go-web-framework/types"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
	"net/http"
	"reflect"
	"strings"
)

var (
	// BC is used to declare base controller
	BC http2.BaseController
	// Container will provide access to the global Service Container
	Container *dig.Container
	// Controllers will handle every application controller
	Controllers types.ControllerRegister
	// Middleware will handle every application middleware
	Middleware interface{}
	// Models will handle every application middleware
	Models types.ModelRegister
	// Services will handle every application service
	Services = types.ServiceRegister{
		List: []interface{}{
			service.Configuration,
			CreateSessionStore,
			GetHttpServer,
			WebRouter,
		},
	}
	// Commands will export all registered commands
	// The following map of interfaces expose all available method that can be used by Go-Web CLI tool.
	// The map index determines the command that you've to run to for use the relative method.
	// Example: './goweb migration:up' will run '&command.MigrationUp{}' command.
	Commands = types.CommandRegister{
		List: map[string]interface{}{
			"database:seed": &cli.Seeder{},
			"server:daemon": &cli.ServerDaemon{},
			"server:run":    &cli.ServerRun{},
			"show:commands": &cli.ShowCommands{},
			// Here is where you've to register your custom controller
		},
	}
)

// Handle single path parsing.
// This method it's used to parse every single path. If middleware is present a sub-router with will be created
func HandleSingleRoute(routes map[string]Route, router *mux.Router) {
	for _, route := range routes {
		hasMiddleware := len(route.Middleware) > 0
		directive := strings.Split(route.Action, "@")
		if hasMiddleware {
			subRouter := mux.NewRouter()
			subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				cc := GetControllerInterface(directive, writer, request)
				method := reflect.ValueOf(cc).MethodByName(directive[1])
				if err := Container.Invoke(method.Interface()); err != nil {
					helper.ProcessError(err)
				}
			}).Methods(route.Method)

			subRouter.Use(parseMiddleware(route.Middleware, Middleware)...)
			router.Handle(route.Path, subRouter).Methods(route.Method)
		} else {
			router.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				cc := GetControllerInterface(directive, writer, request)
				method := reflect.ValueOf(cc).MethodByName(directive[1])
				if err := Container.Invoke(method.Interface()); err != nil {
					helper.ProcessError(err)
				}
			}).Methods(route.Method)
		}
	}
}

// Parse route groups.
func HandleGroups(groups map[string]Group, router *mux.Router) {
	for _, group := range groups {
		subRouter := router.PathPrefix(group.Prefix).Subrouter()

		for _, route := range group.Routes {
			directive := strings.Split(route.Action, "@")
			if len(route.Middleware) > 0 {
				nestedRouter := mux.NewRouter()
				fullPath := fmt.Sprintf("%s%s", group.Prefix, route.Path)
				nestedRouter.HandleFunc(fullPath, func(writer http.ResponseWriter, request *http.Request) {
					cc := GetControllerInterface(directive, writer, request)
					method := reflect.ValueOf(cc).MethodByName(directive[1])
					if err := Container.Invoke(method.Interface()); err != nil {
						helper.ProcessError(err)
					}
				}).Methods(route.Method)

				nestedRouter.Use(parseMiddleware(route.Middleware, Middleware)...)
				subRouter.Handle(route.Path, nestedRouter).Methods(route.Method)
			} else {
				subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
					cc := GetControllerInterface(directive, writer, request)
					method := reflect.ValueOf(cc).MethodByName(directive[1])
					if err := Container.Invoke(method.Interface()); err != nil {
						helper.ProcessError(err)
					}
				}).Methods(route.Method)
			}
		}

		subRouter.Use(parseMiddleware(group.Middleware, Middleware)...)
	}
}

// Parse list of middleware and get an array of []mux.Middleware func
// Required by Gorilla Mux
func parseMiddleware(mwList []string, middleware interface{}) []mux.MiddlewareFunc {
	var midFunc []mux.MiddlewareFunc

	for _, mw := range mwList {
		m := reflect.ValueOf(middleware)
		method := m.MethodByName(mw)

		callable := method.Interface().(func(handler http.Handler) http.Handler)
		midFunc = append(midFunc, callable)
	}

	return midFunc
}

// Give access to public folder. With the /public prefix you can access to all of your assets.
// This is mandatory to access to public files (.js, .css, images, etc...).
func GiveAccessToPublicFolder(router *mux.Router) {
	publicDirectory := http.Dir(helper.GetDynamicPath("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(publicDirectory)))
}

// GetControllerInterface  will returns a specific controller instance by comparing "directive" parameter with controller name.
func GetControllerInterface(directive []string, w http.ResponseWriter, r *http.Request) interface{} {
	var result interface{}

	// Find the right controller
	for _, contr := range Controllers.List {
		controllerName := reflect.Indirect(reflect.ValueOf(contr)).Type().Name()
		if controllerName == directive[0] {
			registerBaseController(w, r, &contr)
			result = contr
		}
	}

	return result
}

// Parse a controller instance and implement it with the current base controller.
// This operation will give you access to all basic controller properties.
func registerBaseController(res http.ResponseWriter, req *http.Request, controller *interface{}) *interface{} {
	if err := setBaseController(res, req); err != nil {
		helper.ProcessError(err)
	}

	c := reflect.ValueOf(*controller).Elem().FieldByName("BaseController")
	c.Set(reflect.ValueOf(BC))

	return controller
}

// Setting up the base controller instance (defined in conf.go) with the properties/method defined in the Service Container.
// Here you can implement the BaseController content.
// Remember to update even the structure (app/http/controller/controller.go)
func setBaseController(res http.ResponseWriter, req *http.Request) error {
	BC = http2.BaseController{
		Response: res,
		Request:  req,
	}

	return nil
}
