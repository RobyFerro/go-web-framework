package entities

// Route declares a specific http route
type Route struct {
	Name        string
	Path        string
	Action      string
	Method      string
	Description string
	Validation  interface{}
	Middleware  []Middleware
}
