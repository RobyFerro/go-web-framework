package go_web_framework

import "net/http"

// Middleware struct is extended by every middleware.
type Middleware struct {
	Handler http.Handler
}
