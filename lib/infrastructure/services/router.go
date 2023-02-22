package services

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/RobyFerro/go-web-framework/helpers"
	"github.com/RobyFerro/go-web-framework/lib/domain/entities"
	"github.com/RobyFerro/go-web-framework/lib/domain/kernel"
	"github.com/RobyFerro/go-web-framework/lib/domain/registers"
	"github.com/julienschmidt/httprouter"
)

// RouterServiceImpl implements RouterService interfaces ans handles http router methods.
type RouterServiceImpl struct {
	Controllers registers.ControllerRegister
}

// NewRouter returns a new HTTP Router
func (r RouterServiceImpl) NewRouter(register registers.RouterRegister) http.Handler {
	router := httprouter.New()

	for _, registeredRouter := range register {
		r.injectRoutesInRouter(registeredRouter.Route, router)
		r.injectGroupsInRouter(registeredRouter.Groups, router)
	}

	return router
}

func (r RouterServiceImpl) injectRoutesInRouter(routes []entities.Route, router *httprouter.Router) {
	for _, route := range routes {
		controllerHandler := r.generateControllerHandler(route)
		handlerWithMiddlewares := r.injectRouteMiddlewares(route.Middleware, controllerHandler)

		router.HandlerFunc(route.Method, route.Path, handlerWithMiddlewares)
	}
}

func (r RouterServiceImpl) injectGroupsInRouter(groups []entities.Group, router *httprouter.Router) {
	for _, group := range groups {
		for _, route := range group.Routes {
			controllerHandler := r.generateControllerHandler(route)
			pathWithPrefix := fmt.Sprintf("%s%s", group.Prefix, route.Path)
			joinedMiddlewares := append(group.Middleware, route.Middleware...)

			handlerWithMiddlewares := r.injectRouteMiddlewares(joinedMiddlewares, controllerHandler)
			router.HandlerFunc(route.Method, pathWithPrefix, handlerWithMiddlewares)
		}
	}
}

func (r RouterServiceImpl) joinGroupMiddlewares(group []entities.Middleware, path []entities.Middleware) []entities.Middleware {
	return append(group, path...)
}

func (r RouterServiceImpl) injectRouteMiddlewares(middlewares []entities.Middleware, handler http.HandlerFunc) http.HandlerFunc {
	var result http.Handler
	for _, middleware := range middlewares {
		if result != nil {
			result = middleware.Handle(result)
		} else {
			result = middleware.Handle(handler)
		}
	}

	return http.HandlerFunc(func(resp http.ResponseWriter, request *http.Request) {
		result.ServeHTTP(resp, request)
	})
}

func (r RouterServiceImpl) generateControllerHandler(route entities.Route) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, request *http.Request) {
		if err := r.validateRequest(route.Validation, request); err != nil {
			resp.WriteHeader(http.StatusUnprocessableEntity)
			_, _ = resp.Write([]byte(err.Error()))

			return
		}

		r.invokeAction(route.Action, resp, request)
	})
}

func (r RouterServiceImpl) validateRequest(data interface{}, req *http.Request) error {
	if data != nil {
		if err := helpers.DecodeJSONRequest(req, &data); err != nil {
			return err
		}

		if err := helpers.ValidateRequest(data); err != nil {
			return err
		}
	}

	return nil
}

func (r RouterServiceImpl) invokeAction(
	routeAction string,
	res http.ResponseWriter,
	req *http.Request) {
	controllerData := strings.Split(routeAction, "@")
	item := r.getControllerInterface(controllerData[0])

	kernel.RegisterBaseController(res, req, &item)
	method := reflect.ValueOf(item).MethodByName(controllerData[1])
	method.Call([]reflect.Value{})
}

func (r RouterServiceImpl) getControllerInterface(itemName string) interface{} {
	var result interface{}
	for _, c := range r.Controllers {
		controllerName := reflect.Indirect(reflect.ValueOf(c)).Type().Name()
		if controllerName == itemName {
			result = c
		}
	}

	return result
}
