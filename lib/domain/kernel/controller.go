package kernel

import (
	"net/http"
	"reflect"

	"github.com/RobyFerro/go-web-framework/lib/domain/entities"
)

// RegisterBaseController parse a controller instance and implement it with the current base controller.
// This operation will give you access to all basic controller properties.
func RegisterBaseController(res http.ResponseWriter, req *http.Request, controller *interface{}) {
	baseController := entities.BaseController{
		Response: res,
		Request:  req,
	}

	c := reflect.ValueOf(*controller).Elem().FieldByName("BaseController")
	c.Set(reflect.ValueOf(baseController))
}
