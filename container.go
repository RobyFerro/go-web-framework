package gwf

import (
	gwf "github.com/RobyFerro/go-web-framework"
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
		ProcessError(err)
	}

	if err := container.Provide(Configuration); err != nil {
		ProcessError(err)
	}

	if err := container.Provide(CreateSessionStore); err != nil {
		ProcessError(err)
	}

	err = container.Invoke(func(conf gwf.Conf) {
		if len(conf.Redis.Host) > 0 {
			if err := container.Provide(ConnectRedis); err != nil {
				ProcessError(err)
			}
		}

		if len(conf.Database.Host) > 0 {
			if err := container.Provide(ConnectDB); err != nil {
				ProcessError(err)
			}
		}

		if len(conf.Mongo.Host) > 0 {
			if err := container.Provide(ConnectMongo); err != nil {
				ProcessError(err)
			}
		}

		if len(conf.Elastic.Hosts) > 0 {
			if err := container.Provide(ConnectElastic); err != nil {
				ProcessError(err)
			}
		}
	})

	if err != nil {
		ProcessError(err)
	}

	if err := container.Provide(GetHttpServer); err != nil {
		ProcessError(err)
	}

	if err := container.Provide(SetAuth); err != nil {
		ProcessError(err)
	}

	return container
}
