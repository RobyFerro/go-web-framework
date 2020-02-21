package command

import (
	"net/http"

	gwf "github.com/RobyFerro/go-web-framework"
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
func (c *ServerRun) Run(srv *http.Server, conf *gwf.Conf) {
	if err := gwf.StartServer(srv, conf); err != nil {
		gwf.ProcessError(err)
	}
}
