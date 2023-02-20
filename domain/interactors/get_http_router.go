package interactors

import (
	"net/http"

	register "github.com/RobyFerro/go-web-framework/domain/registers"
	"github.com/RobyFerro/go-web-framework/domain/services"
)

// GetHTTPRouter generates a new application router
type GetHTTPRouter struct {
	Service  services.RouterService
	Register register.RouterRegister
}

// Call executes usecase logic
func (c GetHTTPRouter) Call() http.Handler {
	return c.Service.NewRouter(c.Register)
}
