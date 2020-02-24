package gwf

import (
	"log"
	"net/http"

	daemon "github.com/sevlyar/go-daemon"
)

// ServerDaemon will run Go-Web HTTP server in daemon "mode"
type ServerDaemon struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *ServerDaemon) Register() {
	c.Signature = "server:daemon"
	c.Description = "Run Go-Web server as a daemon"
}

// Run Go-Web as a daemon
func (c *ServerDaemon) Run(conf *Conf, srv *http.Server) {
	// Simple way to check is a string contains only digits
	cntxt := &daemon.Context{
		PidFileName: "storage/log/go-web.pid",
		PidFilePerm: 0754,
		LogFileName: "storage/log/go-webd.log",
		LogFilePerm: 0754,
		Umask:       027,
	}

	daemonBg, daemonError := cntxt.Reborn()

	if daemonError != nil {
		log.Fatal("Troubles to start daemon ", daemonError, daemonError.Error())
	}

	if daemonBg != nil {
		return
	}

	defer func() {
		_ = cntxt.Release()
	}()

	if err := StartServer(srv, conf); err != nil {
		ProcessError(err)
	}
}
