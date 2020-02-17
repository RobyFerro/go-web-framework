package service

import (
	"github.com/RobyFerro/go-web-framework"
	"go.uber.org/dig"
)

// Here is where service container is built.
// As you can see the service container is provided by Uber DIG library.
// Se its documentation (https://godoc.org/go.uber.org/dig) to implement extra services.
func BuildContainer(controllers []interface{}, middleware interface{}, services []interface{}) *dig.Container {
	container := dig.New()
	bindServices(services)

	gwf.Controllers = controllers
	gwf.Middleware = middleware

	for _, s := range Services {
		if err := container.Provide(s); err != nil {
			gwf.ProcessError(err)
		}
	}

	gwf.Container = container

	return container
}

// Merge custom services with defaults
func bindServices(services []interface{}) {
	for _, s := range services {
		Services = append(Services, s)
	}
}

// Declare framework services.
var Services = []interface{}{
	gwf.Configuration,
	gwf.CreateSessionStore,
	gwf.GetHttpServer,
	gwf.SetAuth,
	gwf.WebRouter,
}
