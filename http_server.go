package gwf

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var appConf Conf

// StartServer will run the Go HTTP web server
func StartServer(srv *http.Server, conf *Conf) error {
	appConf = *conf
	webListener, _ := net.Listen("tcp4", ":"+strconv.Itoa(conf.Server.Port))

	if runtime.GOOS == "linux" {
		if err := changeRunningUser(); err != nil {
			return err
		}
	}

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

// Prepare HTTP server for Service Container
func GetHttpServer(router *mux.Router, cfg *Conf) *http.Server {
	serverString := fmt.Sprintf("%s:%d", cfg.Server.Name, cfg.Server.Port)

	var httpServerConf = http.Server{}

	if cfg.Server.Ssl {
		sslCfg := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}

		// Add TLS configuration to http server

		httpServerConf = http.Server{
			Addr:    serverString,
			Handler: router,
			//			ReadTimeout:  time.Duration(agentconfig.Ag.Agent.HttpRTimeout) * time.Second,
			//			WriteTimeout: time.Duration(agentconfig.Ag.Agent.HttpWTimeout) * time.Second,
			TLSConfig:    sslCfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}

	} else {
		httpServerConf = http.Server{
			Addr:    serverString,
			Handler: router,
		}
	}

	return &httpServerConf
}

// Create session CookieStore
func CreateSessionStore(cfg *Conf) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(os.Getenv(cfg.App.Key)))
}

// Change running user. This method works only on Linux systems
// If you'd like to run go-web on Windows or OSX system you should avoid the following code
func changeRunningUser() error {
	var numUID int
	var numGID int

	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	if len(appConf.Server.RunUser) > 0 && len(appConf.Server.RunGroup) > 0 {
		// Check the way UID is written (digits or string)
		if strings.IndexFunc(appConf.Server.RunUser, isNotDigit) == 0 {
			// If UID is a string, we lookup it to an int
			_uid, _ := user.Lookup(appConf.Server.RunUser)

			if _uid == nil {
				return nil
			}

			if numUID, err := strconv.Atoi(_uid.Uid); err != nil {
				return err
			} else {
				if err := changeUID(numUID); err != nil {
					return err
				}
			}

		} else {
			numUID, _ = strconv.Atoi(appConf.Server.RunUser)
		}

		// Check the way GID is written (digits or string)
		if strings.IndexFunc(appConf.Server.RunGroup, isNotDigit) == 0 {
			// If UID is a string, we lookup it to an int
			_gid, _ := user.LookupGroup(appConf.Server.RunGroup)

			if _gid == nil {
				return nil
			}

			if numGID, err := strconv.Atoi(_gid.Gid); err != nil {
				return err
			} else {
				if err := changeGID(numGID); err != nil {
					return err
				}
			}

		} else {
			numGID, _ = strconv.Atoi(appConf.Server.RunGroup)
		}

		log.Printf("Changing running user to %d:%d\n", uint32(numUID), uint32(numGID))
	}

	return nil
}

// Execute a syscall to change active user.
func changeUID(uid int) error {
	if syscall.Getuid() == uid {
		return nil
	}

	if err := syscall.Setuid(uid); err != nil {
		return err
	}

	return nil
}

// Execute a syscall to change active group
func changeGID(gid int) error {
	if syscall.Getgid() == gid {
		return nil
	}

	if err := syscall.Setgid(gid); err != nil {
		return err
	}

	return nil
}
