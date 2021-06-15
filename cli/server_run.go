package cli

import (
	"github.com/RobyFerro/go-web-framework/foundation"
	"github.com/RobyFerro/go-web-framework/helper"
	"github.com/RobyFerro/go-web-framework/service"
	"net/http"
)

// ServerRun will run Go-Web HTTP server
type ServerRun struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *ServerRun) Register() {
	c.Signature = "server:run"
	c.Description = "Run Go-Web server"
}

// Run this command
func (c *ServerRun) Run(srv *http.Server, conf *service.Conf) {
	if err := foundation.StartServer(srv, conf); err != nil {
		helper.ProcessError(err)
	}
}
