package go_web_framework

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

var (
	// Declaring base controller
	BC BaseController
	// Get app configuration
	BaseConfig, _ = Configuration()
	// Register service container
	Container *dig.Container
)

type HttpKernel struct {
	Models    []interface{}
	Container *dig.Container
}

// Parse routing structures and set every route.
// Return a Gorilla Mux router instance with all routes indicated in router.yml file.
func WebRouter() (*mux.Router, error) {
	var wg sync.WaitGroup
	wg.Add(3)

	routes, err := ConfigurationWeb()
	if err != nil {
		return nil, err
	}

	router := mux.NewRouter()

	go func() {
		handleSingleRoute(routes.Routes, router)
		wg.Done()
	}()

	go func() {
		handleGroups(routes.Groups, router)
		wg.Done()
	}()

	go func() {
		giveAccessToPublicFolder(router)
		wg.Done()
	}()

	wg.Wait()

	return router, nil
}

// Handle single path parsing.
// This method it's used to parse every single path. If middleware is present a sub-router with will be created
func handleSingleRoute(routes map[string]Route, router *mux.Router) {
	var wg sync.WaitGroup
	wg.Add(len(routes))

	for _, route := range routes {
		go func(r Route) {
			hasMiddleware := len(r.Middleware) > 0
			directive := strings.Split(r.Action, "@")
			if hasMiddleware {
				subRouter := mux.NewRouter()
				subRouter.HandleFunc(r.Path, func(writer http.ResponseWriter, request *http.Request) {
					cc := GetControllerInterface(directive, writer, request)
					method := reflect.ValueOf(cc).MethodByName(directive[1])
					method.Call([]reflect.Value{})
				}).Methods(r.Method)

				subRouter.Use(parseMiddleware(r.Middleware)...)
				router.Handle(r.Path, subRouter)
			} else {
				router.HandleFunc(r.Path, func(writer http.ResponseWriter, request *http.Request) {
					cc := GetControllerInterface(directive, writer, request)
					method := reflect.ValueOf(cc).MethodByName(directive[1])
					method.Call([]reflect.Value{})
				}).Methods(r.Method)
			}

			wg.Done()
		}(route)
	}

	wg.Wait()
}

// Parse route groups.
func handleGroups(groups map[string]Group, router *mux.Router) {
	for _, group := range groups {
		subRouter := router.PathPrefix(group.Prefix).Subrouter()
		var wg sync.WaitGroup
		wg.Add(len(group.Routes))

		for _, route := range group.Routes {
			go func(r Route) {
				directive := strings.Split(r.Action, "@")
				if len(r.Middleware) > 0 {
					nestedRouter := mux.NewRouter()
					fullPath := fmt.Sprintf("%s%s", group.Prefix, r.Path)
					nestedRouter.HandleFunc(fullPath, func(writer http.ResponseWriter, request *http.Request) {
						cc := GetControllerInterface(directive, writer, request)
						method := reflect.ValueOf(cc).MethodByName(directive[1])
						method.Call([]reflect.Value{})
					}).Methods(r.Method)

					nestedRouter.Use(parseMiddleware(r.Middleware)...)
					subRouter.Handle(r.Path, nestedRouter)
				} else {
					subRouter.HandleFunc(r.Path, func(writer http.ResponseWriter, request *http.Request) {
						cc := GetControllerInterface(directive, writer, request)
						method := reflect.ValueOf(cc).MethodByName(directive[1])
						method.Call([]reflect.Value{})
					}).Methods(r.Method)
				}

				wg.Done()
			}(route)
		}

		wg.Wait()

		subRouter.Use(parseMiddleware(group.Middleware)...)
	}
}

// Parse list of middleware and get an array of []mux.Middleware func
// Required by Gorilla Mux
func parseMiddleware(mwList []string) []mux.MiddlewareFunc {
	var midFunc []mux.MiddlewareFunc

	for _, mw := range mwList {
		m := reflect.ValueOf(Middleware{})
		method := m.MethodByName(mw)

		callable := method.Interface().(func(handler http.Handler) http.Handler)
		midFunc = append(midFunc, callable)
	}

	return midFunc
}

// Give access to public folder. With the /public prefix you can access to all of your assets.
// This is mandatory to access to public files (.js, .css, images, etc...).
func giveAccessToPublicFolder(router *mux.Router) {
	publicDirectory := http.Dir(GetDynamicPath("public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(publicDirectory)))
}

// Returns a specific controller instance by comparing "directive" parameter with controller name.
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

// Parse a controller instance and implement it with the current base controller.
// This operation will give you access to all basic controller properties.
func registerBaseController(res http.ResponseWriter, req *http.Request, controller *interface{}) *interface{} {
	if err := setBaseController(res, req); err != nil {
		ProcessError(err)
	}
	if err := checkControllerIntegrations(&BC); err != nil {
		ProcessError(err)
	}

	c := reflect.ValueOf(*controller).Elem().FieldByName("BaseController")
	c.Set(reflect.ValueOf(BC))

	return controller
}

// Setting up the base controller instance (defined in conf.go) with the properties/method defined in the Service Container.
// Here you can implement the BaseController content.
// Remember to update even the structure (app/http/controller/controller.go)
func setBaseController(res http.ResponseWriter, req *http.Request) error {
	if err := Container.Invoke(func(db *gorm.DB, c Conf, a *Auth, s *sessions.CookieStore) {
		BC = BaseController{
			DB:       db,
			Response: res,
			Request:  req,
			Config:   c,
			Auth:     a,
			Session:  s,
		}
	}); err != nil {
		return err
	}

	return nil
}

// Check controller integrations
// Es: Redis, Elastic, Mongo connections
func checkControllerIntegrations(base *BaseController) error {

	// If is configured MongoDB will be implemented into service container
	if len(BaseConfig.Mongo.Host) > 0 {
		if err := Container.Invoke(func(m *mongo.Database) {
			base.Mongo = m
		}); err != nil {
			return err
		}
	}

	// If is configured Redis will be implemented into service container
	if len(BaseConfig.Redis.Host) > 0 {
		if err := Container.Invoke(func(r *redis.Client) {
			base.Redis = r
		}); err != nil {
			return err
		}
	}

	// If is configured ElasticSearch will be implemented into service container
	if len(BaseConfig.Elastic.Hosts) > 0 {
		if err := Container.Invoke(func(e *elasticsearch.Client) {
			base.Elastic = e
		}); err != nil {
			return err
		}
	}

	return nil
}
