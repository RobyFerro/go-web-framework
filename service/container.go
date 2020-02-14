package service

import (
	"github.com/RobyFerro/go-web-framework/config"
	"github.com/RobyFerro/go-web-framework/database"
	"github.com/RobyFerro/go-web-framework/exception"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
)

// Here is where service container is built.
// As you can see the service container is provided by Uber DIG library.
// Se its documentation (https://godoc.org/go.uber.org/dig) to implement extra services.
func BuildContainer(router *mux.Router) *dig.Container {
	container := dig.New()

	err := container.Provide(func() *mux.Router {
		return router
	})

	if err != nil {
		exception.ProcessError(err)
	}

	if err := container.Provide(config.Configuration); err != nil {
		exception.ProcessError(err)
	}

	if err := container.Provide(CreateSessionStore); err != nil {
		exception.ProcessError(err)
	}

	err = container.Invoke(func(conf config.Conf) {
		if len(conf.Redis.Host) > 0 {
			if err := container.Provide(ConnectRedis); err != nil {
				exception.ProcessError(err)
			}
		}

		if len(conf.Database.Host) > 0 {
			if err := container.Provide(database.ConnectDB); err != nil {
				exception.ProcessError(err)
			}
		}

		if len(conf.Mongo.Host) > 0 {
			if err := container.Provide(ConnectMongo); err != nil {
				exception.ProcessError(err)
			}
		}

		if len(conf.Elastic.Hosts) > 0 {
			if err := container.Provide(ConnectElastic); err != nil {
				exception.ProcessError(err)
			}
		}
	})

	if err != nil {
		exception.ProcessError(err)
	}

	if err := container.Provide(GetHttpServer); err != nil {
		exception.ProcessError(err)
	}

	if err := container.Provide(SetAuth); err != nil {
		exception.ProcessError(err)
	}

	return container
}
