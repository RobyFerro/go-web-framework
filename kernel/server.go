package kernel

import (
	"log"
	"net"
	"net/http"
	"strconv"
)

var appConf Conf

// RunServer this command
func RunServer() {
	router := WebRouter()
	conf, _ := RetrieveAppConf()
	server := GetHttpServer(router, conf)

	if err := startServer(server, conf); err != nil {
		log.Fatal(err)
	}
}

// startServer will run the Go HTTP web server
func startServer(srv *http.Server, conf *Conf) error {
	appConf = *conf
	webListener, _ := net.Listen("tcp4", ":"+strconv.Itoa(conf.Server.Port))

	if appConf.Server.Ssl {
		if err := srv.ServeTLS(webListener, appConf.Server.SslCert, appConf.Server.SslKey); err != nil {
			return err
		}
	} else {
		if err := srv.Serve(webListener); err != nil {
			return err
		}
	}

	return nil
}
