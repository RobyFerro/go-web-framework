package service

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
)

// Here is where service container is built.
// As you can see the service container is provided by Uber DIG library.
// Se its documentation (https://godoc.org/go.uber.org/dig) to implement extra services.
func BuildContainer(controllers []interface{}, middleware interface{}) *dig.Container {
	container := dig.New()

	gwf.Controllers = controllers
	gwf.Middleware = middleware

	err := container.Provide(func() *mux.Router {
		router, err := gwf.WebRouter()
		if err != nil {
			gwf.ProcessError(err)
		}

		return router
	})

	if err != nil {
		gwf.ProcessError(err)
	}

	if err := container.Provide(gwf.Configuration); err != nil {
		gwf.ProcessError(err)
	}

	if err := container.Provide(gwf.CreateSessionStore); err != nil {
		gwf.ProcessError(err)
	}

	err = container.Invoke(func(conf *gwf.Conf) {
		if len(conf.Redis.Host) > 0 {
			if err := container.Provide(gwf.ConnectRedis); err != nil {
				gwf.ProcessError(err)
			}
		}

		if len(conf.Database.Host) > 0 {
			if err := container.Provide(gwf.ConnectDB); err != nil {
				gwf.ProcessError(err)
			}
		}

		if len(conf.Mongo.Host) > 0 {
			if err := container.Provide(gwf.ConnectMongo); err != nil {
				gwf.ProcessError(err)
			}
		}

		if len(conf.Elastic.Hosts) > 0 {
			if err := container.Provide(gwf.ConnectElastic); err != nil {
				gwf.ProcessError(err)
			}
		}
	})

	if err != nil {
		gwf.ProcessError(err)
	}

	if err := container.Provide(gwf.GetHttpServer); err != nil {
		gwf.ProcessError(err)
	}

	if err := container.Provide(gwf.SetAuth); err != nil {
		gwf.ProcessError(err)
	}

	setContainer(container)

	return container
}

func setContainer(c *dig.Container) {
	gwf.Container = c
}
