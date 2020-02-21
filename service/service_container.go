package service

import (
	gwf "github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web-framework/console"
	"go.uber.org/dig"
)

// BuildContainer provide the global service container
func BuildContainer() *dig.Container {
	container := dig.New()
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

// Services declares all framework services.
var Services = []interface{}{
	gwf.Configuration,
	gwf.CreateSessionStore,
	gwf.GetHttpServer,
	gwf.WebRouter,
}
