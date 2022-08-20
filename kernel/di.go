package kernel

import (
	"log"
	"go.uber.org/dig"
	"github.com/RobyFerro/go-web-framework/register"
)

// BuildCustomContainer provides a service container with custom services.
// It returns a container that will only be user on the HTTP controllers.
func BuildCustomContainer(modules []register.DIModule) *dig.Container {
	container := dig.New()

	for _, m := range modules {
		for _, p := range m.Provides {
			if err := container.Provide(p); err != nil {
				log.Fatal(err)
			}
		}
	}

	return container
}

// BuildCommandContainer builds a service container that will be used only to run console commands.
func BuildCommandContainer() *dig.Container {
	container := dig.New()

	for _, s := range CommandServices {
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

	_ = sc.Provide(func() []register.HTTPRouter {
		return Router
	})
}
