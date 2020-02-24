package gwf

import (
	"go.uber.org/dig"
)

// BuildContainer provide the global service container
func BuildContainer() *dig.Container {
	container := dig.New()

	for _, s := range Services {
		if err := container.Provide(s); err != nil {
			ProcessError(err)
		}
	}

	injectBasicEntities(container)
	Container = container

	return container
}

// Inject base entities: controllers, models, commands in service container
func injectBasicEntities(sc *dig.Container) {
	_ = sc.Provide(func() ControllerRegister {
		return Controllers
	})

	_ = sc.Provide(func() CommandRegister {
		return Commands
	})

	_ = sc.Provide(func() ModelRegister {
		return Models
	})
}
