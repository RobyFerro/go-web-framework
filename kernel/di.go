package kernel

import (
	"github.com/RobyFerro/go-web-framework/register"
	"go.uber.org/dig"
	"log"
)

// BuildCustomContainer provides a service container with custom services.
// It returns a container that will only be user on the HTTP controllers.
func BuildCustomContainer() *dig.Container {
	container := dig.New()

	for _, s := range CustomServices.List {
		if err := container.Provide(s); err != nil {
			log.Fatal(err)
		}
	}

	return container
}

// BuildSingletonContainer provide the global service container
func BuildSingletonContainer() *dig.Container {
	container := dig.New()

	for _, s := range SingletonServices.List {
		if err := container.Provide(s); err != nil {
			log.Fatal(err)
		}
	}

	injectBasicEntities(container)

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
