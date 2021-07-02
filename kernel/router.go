package kernel

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework/tool"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

// Route structure is used to decode all route presents into routing.yml file.
type Route struct {
	Path        string   `yaml:"path"`
	Action      string   `yaml:"action"`
	Method      string   `yaml:"method"`
	Description string   `yaml:"description"`
	Middleware  []string `yaml:"middleware"`
	Prefix      string   `yaml:"prefix"`
}

// Group structure used to decode all groups presents into the routing.yml file.
type Group struct {
	Prefix     string `yaml:"prefix"`
	Routes     map[string]Route
	Middleware []string
}

// Router structure of web router.
type Router struct {
	Routes map[string]Route `yaml:"routes"`
	Groups map[string]Group `yaml:"groups"`
}

// WebRouter parses routing structures and set every route.
// Return a Gorilla Mux router instance with all routes indicated in router.yml file.
func WebRouter() *mux.Router {
	var wg sync.WaitGroup

	wg.Add(3)

	routes, err := RetrieveRoutingConf()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	go func() {
		HandleSingleRoute(routes.Routes, router)
		wg.Done()
	}()

	go func() {
		HandleGroups(routes.Groups, router)
		wg.Done()
	}()

	go func() {
		GiveAccessToPublicFolder(router)
		wg.Done()
	}()

	wg.Wait()

	return router
}

// HandleSingleRoute handles single path parsing.
// This method it's used to parse every single path. If middleware is present a sub-router with will be created
func HandleSingleRoute(routes map[string]Route, router *mux.Router) {
	for _, route := range routes {
		hasMiddleware := len(route.Middleware) > 0
		directive := strings.Split(route.Action, "@")
		if hasMiddleware {
			subRouter := mux.NewRouter()
			subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				executeControllerDirective(directive, writer, request)
			}).Methods(route.Method)

			subRouter.Use(parseMiddleware(route.Middleware, Middleware)...)
			router.Handle(route.Path, subRouter).Methods(route.Method)
		} else {
			router.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
				executeControllerDirective(directive, writer, request)
			}).Methods(route.Method)
		}
	}
}

// HandleGroups parses route groups.
func HandleGroups(groups map[string]Group, router *mux.Router) {
	for _, group := range groups {
		subRouter := router.PathPrefix(group.Prefix).Subrouter()

		for _, route := range group.Routes {
			directive := strings.Split(route.Action, "@")
			if len(route.Middleware) > 0 {
				nestedRouter := mux.NewRouter()
				fullPath := fmt.Sprintf("%s%s", group.Prefix, route.Path)
				nestedRouter.HandleFunc(fullPath, func(writer http.ResponseWriter, request *http.Request) {
					executeControllerDirective(directive, writer, request)
				}).Methods(route.Method)

				nestedRouter.Use(parseMiddleware(route.Middleware, Middleware)...)
				subRouter.Handle(route.Path, nestedRouter).Methods(route.Method)
			} else {
				subRouter.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
					executeControllerDirective(directive, writer, request)
				}).Methods(route.Method)
			}
		}

		subRouter.Use(parseMiddleware(group.Middleware, Middleware)...)
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
	for _, contr := range Controllers.List {
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
func executeControllerDirective(directive []string, w http.ResponseWriter, r *http.Request) {
	container := BuildCustomContainer()
	cc := GetControllerInterface(directive, w, r)
	method := reflect.ValueOf(cc).MethodByName(directive[1])
	if err := container.Invoke(method.Interface()); err != nil {
		log.Fatal(err)
	}
}
