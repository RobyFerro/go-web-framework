package command

import (
	gwf "github.com/RobyFerro/go-web-framework"
)

// ServerRun will run Go-Web HTTP server
type ServerRun struct {
	Signature   string
	Description string
}

// Register this command
func (c *ServerRun) Register() {
	c.Signature = "server:run"
	c.Description = "Run Go-Web server"
}

// Run this command
func (c *ServerRun) Run(kernel *gwf.HttpKernel, args string, console map[string]interface{}) {
	if err := gwf.StartServer(kernel.Container); err != nil {
		gwf.ProcessError(err)
	}
}
