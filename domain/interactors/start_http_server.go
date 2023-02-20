package interactors

import (
	"net"
	"net/http"
	"strconv"

	"github.com/RobyFerro/go-web-framework/domain/entities"
)

// StartHTTPServer runs a new HTTP server
func StartHTTPServer(server http.Server, conf entities.AppConf) error {
	webListener, _ := net.Listen("tcp4", ":"+strconv.Itoa(conf.Port))
	if err := server.Serve(webListener); err != nil {
		return err
	}

	return nil
}
