package services

import (
	"net/http"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/tool"
	"github.com/julienschmidt/httprouter"
)

// RouterServiceImpl implements RouterService interfaces ans handles http router methods.
type RouterServiceImpl struct{}

// NewRouter returns a new HTTP Router
func (r RouterServiceImpl) NewRouter(register entities.RouterRegister) http.Handler {
	router := httprouter.New()

	parseSingleRoutes(register.Route, router)
	parseGroupRoutes(register.Groups, router)

	return router
}

func parseSingleRoutes(routes []entities.Route, router *httprouter.Router) {
	for _, route := range routes {
		registerHandler(route, router, nil)
	}
}

func parseGroupRoutes(groups []entities.Group, router *httprouter.Router) {
	for _, group := range groups {
		for _, route := range group.Routes {
			registerHandler(route, router, &group.Prefix)
		}
	}
}

func registerHandler(route entities.Route, router *httprouter.Router, prefix *string) {
	router.HandlerFunc(route.Method, route.Path, func(writer http.ResponseWriter, request *http.Request) {
		if err := validateRequest(route.Validation, request); err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			_, _ = writer.Write([]byte(err.Error()))

			return
		}

		// TODO: Execute controller directive
	})
}

func validateRequest(data interface{}, r *http.Request) error {
	if data != nil {
		if err := tool.DecodeJsonRequest(r, &data); err != nil {
			return err
		}

		if err := tool.ValidateRequest(data); err != nil {
			return err
		}
	}

	return nil
}
