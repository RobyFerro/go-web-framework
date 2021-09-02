package kernel

import (
	"fmt"
	"github.com/RobyFerro/dig"
	"github.com/RobyFerro/go-web-framework/register"
	"github.com/RobyFerro/go-web-framework/tool"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"strings"
)

var SingletonIOC *dig.Container

// WebRouter parses routing structures and set every route.
// Return a Gorilla Mux router instance with all routes indicated in router.yml file.
func WebRouter(routes []register.HTTPRouter) *mux.Router {
	SingletonIOC = BuildSingletonContainer()
	router := mux.NewRouter()
	router.Use(gzipMiddleware)

	for _, r := range routes {
		if len(r.Route) > 0 {
			HandleSingleRoute(r.Route, router)
		}

		if len(r.Groups) > 0 {
			HandleGroups(r.Groups, router)
		}

		GiveAccessToPublicFolder(router)
	}

	return router
}

// HandleSingleRoute handles single path parsing.
// This method it's used to parse every single path. If middleware is present a sub-router with will be created
func HandleSingleRoute(routes []register.Route, router *mux.Router) {
	for _, route := range routes {
		hasMiddleware := len(route.Middleware) > 0
		directive := strings.Split(route.Action, "@")
		validation := route.Validation
		if hasMiddleware {
			subRouter := mux.NewRouter()
			subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				if err := validateRequest(validation, request); err != nil {
					writer.WriteHeader(http.StatusUnprocessableEntity)
					_, _ = writer.Write([]byte(err.Error()))

					return
				}

				executeControllerDirective(directive, writer, request)
			}).Methods(route.Method)

			subRouter.Use(parseMiddleware(route.Middleware)...)
			router.Handle(route.Path, subRouter).Methods(route.Method)
		} else {
			router.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				if err := validateRequest(validation, request); err != nil {
					writer.WriteHeader(http.StatusUnprocessableEntity)
					_, _ = writer.Write([]byte(err.Error()))

					return
				}

				executeControllerDirective(directive, writer, request)
			}).Methods(route.Method)
		}
	}
}

// HandleGroups parses route groups.
func HandleGroups(groups []register.Group, router *mux.Router) {
	for _, group := range groups {
		subRouter := router.PathPrefix(group.Prefix).Subrouter()

		for _, route := range group.Routes {
			directive := strings.Split(route.Action, "@")
			validation := route.Validation
			if len(route.Middleware) > 0 {
				nestedRouter := mux.NewRouter()
				fullPath := fmt.Sprintf("%s%s", group.Prefix, route.Path)
				nestedRouter.HandleFunc(fullPath, func(writer http.ResponseWriter, request *http.Request) {
					if err := validateRequest(validation, request); err != nil {
						writer.WriteHeader(http.StatusUnprocessableEntity)
						_, _ = writer.Write([]byte(err.Error()))

						return
					}

					executeControllerDirective(directive, writer, request)
				}).Methods(route.Method)

				nestedRouter.Use(parseMiddleware(route.Middleware)...)
				subRouter.Handle(route.Path, nestedRouter).Methods(route.Method)
			} else {
				subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
					if err := validateRequest(validation, request); err != nil {
						writer.WriteHeader(http.StatusUnprocessableEntity)
						_, _ = writer.Write([]byte(err.Error()))

						return
					}

					executeControllerDirective(directive, writer, request)
				}).Methods(route.Method)
			}
		}

		subRouter.Use(parseMiddleware(group.Middleware)...)
	}
}

// GiveAccessToPublicFolder gives access to public folder. With the /public prefix you can access to all of your assets.
// This is mandatory to access to public files (.js, .css, images, etc...).
func GiveAccessToPublicFolder(router *mux.Router) {
	publicDirectory := http.Dir(tool.GetDynamicPath("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(publicDirectory)))
}

// GetControllerInterface  will returns a specific controller instance by comparing "directive" parameter with controller name.
func GetControllerInterface(directive []string, w http.ResponseWriter, r *http.Request) interface{} {
	var result interface{}

	// Find the right controller
	for _, contr := range Controllers {
		controllerName := reflect.Indirect(reflect.ValueOf(contr)).Type().Name()
		if controllerName == directive[0] {
			registerBaseController(w, r, &contr)
			result = contr
		}
	}

	return result
}

// Executes controller string directives.
// Example: MainController@main
// 	executes the main method from MainController
// It build the CUSTOM SERVICE CONTAINER and invoke the selected directive inside them.
func executeControllerDirective(d []string, w http.ResponseWriter, r *http.Request) {
	container := BuildCustomContainer()
	cc := GetControllerInterface(d, w, r)
	method := reflect.ValueOf(cc).MethodByName(d[1])

	if err := dig.GroupInvoke(method.Interface(), container, SingletonIOC); err != nil {
		log.Fatal(err)
	}
}

func validateRequest(data interface{}, r *http.Request) error {
	if data != nil {
		if err := tool.DecodeJsonRequest(r, &data); err != nil {
			return err
		}

		if err := tool.ValidateRequest(data); err != nil {
			return err
		}
	}

	return nil
}
