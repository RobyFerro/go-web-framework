package kernel

import (
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

var Middleware interface{}

// Parse list of middleware and get an array of []mux.Middleware func
// Required by Gorilla Mux
func parseMiddleware(mwList []string, middleware interface{}) []mux.MiddlewareFunc {
	var midFunc []mux.MiddlewareFunc

	for _, mw := range mwList {
		m := reflect.ValueOf(middleware)
		method := m.MethodByName(mw)

		callable := method.Interface().(func(handler http.Handler) http.Handler)
		midFunc = append(midFunc, callable)
	}

	return midFunc
}
