package kernel

import (
	"log"
	"net/http"
	"reflect"
)

// BaseController represents the controller structure. This structure determines what you can find in the
// controllers instance. Adding something else inside this struct will not directly implement the struct.
// This because is just a part of the controller construction. See the "setBaseController"
// method inside app/kernel/kernel.go
type BaseController struct {
	Response http.ResponseWriter // HTTP response
	Request  *http.Request       // HTTP request
}

var (
	// BC is used to declare base controller
	BC BaseController
)

// Parse a controller instance and implement it with the current base controller.
// This operation will give you access to all basic controller properties.
func registerBaseController(res http.ResponseWriter, req *http.Request, controller *interface{}) *interface{} {
	if err := setBaseController(res, req); err != nil {
		log.Fatal(err)
	}

	c := reflect.ValueOf(*controller).Elem().FieldByName("BaseController")
	c.Set(reflect.ValueOf(BC))

	return controller
}

// Setting up the base controller instance (defined in conf.go) with the properties/method defined in the
// Service Container. Here you can implement the BaseController content.
// Remember to update even the structure (app/http/controller/controller.go)
func setBaseController(res http.ResponseWriter, req *http.Request) error {
	BC = BaseController{
		Response: res,
		Request:  req,
	}

	return nil
}
