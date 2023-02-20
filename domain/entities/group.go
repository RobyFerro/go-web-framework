package entities

// Group declares a group of http routes.
type Group struct {
	Name       string
	Prefix     string
	Routes     []Route
	Middleware []Middleware
}
