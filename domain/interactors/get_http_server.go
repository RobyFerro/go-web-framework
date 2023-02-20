package interactors

import (
	"fmt"
	"net/http"

	"github.com/RobyFerro/go-web-framework/domain/entities"
)

// GetHTTPServer handles http server
type GetHTTPServer struct {
	Config entities.Config
	Router http.Handler
}

// Call executes usecase logic
func (c GetHTTPServer) Call() *http.Server {
	serverString := fmt.Sprintf("%s:%d", c.Config.Name, c.Config.Port)

	return &http.Server{
		Addr:    serverString,
		Handler: c.Router,
	}
}
