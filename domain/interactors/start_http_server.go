package interactors

import (
	"fmt"
	"net/http"

	"github.com/RobyFerro/go-web-framework/domain/entities"
)

// StartHTTPServer will return a new GoWeb Http server
func StartHTTPServer(config entities.AppConf, router http.Handler) *http.Server {
	serverString := fmt.Sprintf("%s:%d", config.Name, config.Port)

	return &http.Server{
		Addr:    serverString,
		Handler: router,
	}
}
