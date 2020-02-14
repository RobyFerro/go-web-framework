package gwf

import (
	"github.com/gorilla/mux"
	"go.uber.org/dig"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

// Structure used to decode all route presents into routing.yml file.
type Route struct {
	Path        string   `yaml:"path"`
	Action      string   `yaml:"action"`
	Method      string   `yaml:"method"`
	Description string   `yaml:"description"`
	Middleware  []string `yaml:"middleware"`
	Prefix      string   `yaml:"prefix"`
}

// Structure used to decode all groups presents into the routing.yml file.
type Group struct {
	Prefix     string `yaml:"prefix"`
	Routes     map[string]Route
	Middleware []string
}

// Main structure of web router.
type Router struct {
	Routes map[string]Route `yaml:"routes"`
	Groups map[string]Group `yaml:"groups"`
}

// Parse routing structures and set every route.
// Return a Gorilla Mux router instance with all routes indicated in router.yml file.
func WebRouter(controllers []interface{}, middleware interface{}, sc *dig.Container) (*mux.Router, error) {
	var wg sync.WaitGroup
	wg.Add(3)

	Controllers = controllers
	Middleware = middleware
	Container = sc

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

// Parse router.yml file (present in Go-Web root dir) and return a Router structure.
// This structure will be used by the HTTP kernel to setup every routes.
func ConfigurationWeb() (*Router, error) {
	var conf Router
	routePath := GetDynamicPath("routing.yml")
	c, err := os.Open(routePath)

	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(c)

	if err := decoder.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
