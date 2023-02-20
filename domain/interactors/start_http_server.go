package interactors

import (
	"net"
	"net/http"
	"strconv"

	"github.com/RobyFerro/go-web-framework/domain/entities"
)

// StartHTTPServer starts a new HTTP server
type StartHTTPServer struct {
	Server http.Server
	Config entities.Config
}

// Call executes usecase logic
func (c StartHTTPServer) Call() error {
	webListener, _ := net.Listen("tcp4", ":"+strconv.Itoa(c.Config.Port))
	if err := c.Server.Serve(webListener); err != nil {
		return err
	}

	return nil
}
