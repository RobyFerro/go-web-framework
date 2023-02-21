package entities

import "net/http"

// Middleware declares an http midleware
type Middleware interface {
	Handle(next http.Handler) http.Handler
	GetName() string
	GetDescription() string
}


