package http

import (
	"github.com/RobyFerro/go-web-framework/config"
	"go.uber.org/dig"
	"log"
	"net"
	"net/http"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

var ServiceContainer *dig.Container
var appConf config.Conf

// Start HTTP server
func StartServer(sc *dig.Container) error {
	ServiceContainer = sc
	var serveError error

	if err := sc.Invoke(func(s *http.Server, conf config.Conf) {
		appConf = conf
		webListener, _ := net.Listen("tcp4", ":"+strconv.Itoa(conf.Server.Port))

		if err := changeRunningUser(); err != nil {
			serveError = err
			return
		}

		if appConf.Server.Ssl {
			serveError = s.ServeTLS(webListener, appConf.Server.SslCert, appConf.Server.SslKey)
			return
		} else {
			serveError = s.Serve(webListener)
			return
		}
	}); err != nil {
		return err
	}

	if serveError != nil {
		return serveError
	}

	return nil
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
