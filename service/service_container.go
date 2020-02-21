package service

import (
	gwf "github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web-framework/console"
	"go.uber.org/dig"
)

// BuildContainer provide the global service container
func BuildContainer(
	controllers gwf.ControllerRegister,
	middleware interface{},
	services gwf.ServiceRegister,
	models gwf.ModelRegister,
) *dig.Container {
	container := dig.New()
	bindServices(services)

	gwf.Controllers = controllers
	gwf.Middleware = middleware
	gwf.Models = models

	for _, s := range Services {
		if err := container.Provide(s); err != nil {
			gwf.ProcessError(err)
		}
	}

	injectBasicEntities(container)
	gwf.Container = container

	return container
}

// Inject base entities: controllers, models, commands in service container
func injectBasicEntities(sc *dig.Container) {
	_ = sc.Provide(func() gwf.ControllerRegister {
		return gwf.Controllers
	})

	_ = sc.Provide(func() gwf.CommandRegister {
		return console.Commands
	})

	_ = sc.Provide(func() gwf.ModelRegister {
		return gwf.Models
	})
}

// Merge custom services with defaults
func bindServices(services []interface{}) {
	for _, s := range services {
		Services = append(Services, s)
	}
}

// Services declares all framework services.
var Services = []interface{}{
	gwf.Configuration,
	gwf.CreateSessionStore,
	gwf.GetHttpServer,
	gwf.WebRouter,
}
