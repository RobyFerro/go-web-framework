package kernel

import (
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

// Parse list of middleware and get an array of []mux.Middleware func
// Required by Gorilla Mux
func parseMiddleware(mwList []string) []mux.MiddlewareFunc {
	var midFunc []mux.MiddlewareFunc

	for _, name := range mwList {
		for _, mw := range Middlewares.List {
			rName := reflect.ValueOf(mw).Elem().FieldByName("Name").String()

			if name == rName {
				m := reflect.ValueOf(mw)
				method := m.MethodByName("Handle")

				callable := method.Interface().(func(handler http.Handler) http.Handler)
				midFunc = append(midFunc, callable)
			}
		}
	}

	return midFunc
}
