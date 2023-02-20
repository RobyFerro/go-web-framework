package interactors

import (
	"net/http"

	register "github.com/RobyFerro/go-web-framework/domain/registers"
	"github.com/RobyFerro/go-web-framework/domain/services"
)

// GetHTTPRouter will generate a new http router
func GetHTTPRouter(routerService services.RouterService, register []register.RouterRegister) http.Handler {
	return routerService.NewRouter(register)
}
