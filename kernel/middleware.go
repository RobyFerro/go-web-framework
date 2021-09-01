package kernel

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Middleware interface {
	Handle(next http.Handler) http.Handler
}

// Parse list of middleware and get an array of []mux.Middleware func
// Required by Gorilla Mux
func parseMiddleware(mwList []Middleware) []mux.MiddlewareFunc {
	var midFunc []mux.MiddlewareFunc

	for _, mw := range mwList {
		midFunc = append(midFunc, mw.Handle)
	}

	return midFunc
}
