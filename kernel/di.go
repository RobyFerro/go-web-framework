package kernel

import (
	"github.com/RobyFerro/go-web-framework/register"
	"go.uber.org/dig"
)

// Container will provide access to the global Service Container
var (
	Container *dig.Container
)

// BuildContainer provide the global service container
func BuildContainer() *dig.Container {
	container := dig.New()

	for _, s := range Services.List {
		if err := container.Provide(s); err != nil {

		}
	}

	injectBasicEntities(container)
	Container = container

	return container
}

// Inject base entities: controllers, models, commands in service container
func injectBasicEntities(sc *dig.Container) {
	_ = sc.Provide(func() register.ControllerRegister {
		return Controllers
	})

	_ = sc.Provide(func() register.CommandRegister {
		return Commands
	})

	_ = sc.Provide(func() register.ModelRegister {
		return Models
	})
}
