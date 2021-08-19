package kernel

import (
	"github.com/RobyFerro/dig"
	"github.com/RobyFerro/go-web-framework/register"
	"log"
)

// BuildCustomContainer provides a service container with custom services.
// It returns a container that will only be user on the HTTP controllers.
func BuildCustomContainer() *dig.Container {
	container := dig.New()

	for _, s := range Services.List {
		if err := container.Provide(s); err != nil {
			log.Fatal(err)
		}
	}

	return container
}

// BuildCommandContainer builds a service container that will be used only to run console commands.
func BuildCommandContainer() *dig.Container {
	container := dig.New()

	for _, s := range CommandServices.List {
		if err := container.Provide(s); err != nil {
			log.Fatal(err)
		}
	}
	injectBasicEntities(container)

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
