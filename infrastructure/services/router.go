package services

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/kernel"
	"github.com/RobyFerro/go-web-framework/domain/registers"
	"github.com/RobyFerro/go-web-framework/helpers"
	"github.com/julienschmidt/httprouter"
)

// RouterServiceImpl implements RouterService interfaces ans handles http router methods.
type RouterServiceImpl struct {
	Controllers registers.ControllerRegister
}

// NewRouter returns a new HTTP Router
func (r RouterServiceImpl) NewRouter(register []registers.RouterRegister) http.Handler {
	router := httprouter.New()

	for _, registeredRouter := range register {
		r.parseSingleRoutes(registeredRouter.Route, router)
		r.parseGroupRoutes(registeredRouter.Groups, router)
	}

	return router
}

func (r RouterServiceImpl) parseSingleRoutes(routes []entities.Route, router *httprouter.Router) {
	for _, route := range routes {
		r.registerHandler(route, router, nil)
	}
}

func (r RouterServiceImpl) parseGroupRoutes(groups []entities.Group, router *httprouter.Router) {
	for _, group := range groups {
		for _, route := range group.Routes {
			r.registerHandler(route, router, &group.Prefix)
		}
	}
}

func (r RouterServiceImpl) registerHandler(route entities.Route, router *httprouter.Router, prefix *string) {
	router.HandlerFunc(route.Method, route.Path, func(resp http.ResponseWriter, request *http.Request) {
		if err := r.validateRequest(route.Validation, request); err != nil {
			resp.WriteHeader(http.StatusUnprocessableEntity)
			_, _ = resp.Write([]byte(err.Error()))

			return
		}

		r.executeControllerDirective(route, resp, request)
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

func (r RouterServiceImpl) executeControllerDirective(
	route entities.Route,
	res http.ResponseWriter,
	req *http.Request) {
	controllerData := strings.Split(route.Action, "@")
	item := r.getControllerItem(controllerData[0])

	cc := kernel.RegisterBaseController(res, req, &item.Controller)
	method := reflect.ValueOf(cc).MethodByName(controllerData[1])

	method.Interface()
}

// GetControllerName returns a ControllerRegisterItem structure
func (r RouterServiceImpl) getControllerItem(itemName string) registers.ControllerRegisterItem {
	var result registers.ControllerRegisterItem
	for _, contr := range r.Controllers {
		controllerName := reflect.Indirect(reflect.ValueOf(contr.Controller)).Type().Name()
		if controllerName == itemName {
			result = contr
		}
	}

	return result
}
