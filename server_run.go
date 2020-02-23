package main

import (
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
func (c *ServerRun) Run(srv *http.Server, conf *Conf) {
	if err := StartServer(srv, conf); err != nil {
		ProcessError(err)
	}
}
