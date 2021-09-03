package register

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
