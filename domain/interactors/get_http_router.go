package interactors

import (
	"net/http"

	"github.com/RobyFerro/go-web-framework/domain/entities"
	"github.com/RobyFerro/go-web-framework/domain/services"
)

// GetHTTPRouter will generate a new http router
func GetHTTPRouter(routerService services.RouterService, register entities.RouterRegister) http.Handler {
	return routerService.NewRouter(register)
}
