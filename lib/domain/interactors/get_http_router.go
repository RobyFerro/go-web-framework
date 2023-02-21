package interactors

import (
	"net/http"

	"github.com/RobyFerro/go-web-framework/lib/domain/registers"
	"github.com/RobyFerro/go-web-framework/lib/domain/services"
)

// GetHTTPRouter generates a new application router
type GetHTTPRouter struct {
	Service  services.RouterService
	Register registers.RouterRegister
}

// Call executes usecase logic
func (c GetHTTPRouter) Call() http.Handler {
	return c.Service.NewRouter(c.Register)
}
