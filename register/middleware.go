package register

import "net/http"

// Middleware define an interface used to handle all application middleware
type Middleware interface {
	Handle(next http.Handler) http.Handler
	GetName() string
	GetDescription() string
}

// MiddlewareRegister defines all middleware present in your web application
type MiddlewareRegister []interface{}
