package register

import (
	"net/http"
)

// ServiceRegister defines a controller register type.
// This will be used to resolve this register in service container
type ServiceRegister []interface{}

// ModelRegister defines a controller register type.
// This will be used to resolve this register in service container
type ModelRegister []interface{}

// ControllerRegister defines a controller register type.
// This will be used to resolve this register in service container
type ControllerRegister []interface{}

// CommandRegister defines all registered commands
type CommandRegister map[string]interface{}

// Middleware define an interface used to handle all application middleware
type Middleware interface {
	Handle(next http.Handler) http.Handler
}

// MiddlewareRegister defines all middleware present in your web application
type MiddlewareRegister []interface{}

// Route defines an HTTP Router endpoint
type Route struct {
	Name        string
	Path        string
	Action      string
	Method      string
	Description string
	Validation  interface{}
	Middleware  []Middleware
}

// Group defines a group of HTTP Route
type Group struct {
	Name       string
	Prefix     string
	Routes     []Route
	Middleware []Middleware
}

// HTTPRouter contains Route and Group that defines a complete HTTP Router
type HTTPRouter struct {
	Route  []Route
	Groups []Group
}

// Command is used to define a CLI command
type Command struct {
	Signature   string
	Description string
	Args        string
}
