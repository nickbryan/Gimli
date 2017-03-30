package providers

import (
	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/foundation/skeleton/app"
	"github.com/nickbryan/gimli/foundation/skeleton/app/controllers"
)

type ControllerServiceProvider struct{}

func (provider *ControllerServiceProvider) Register(container di.Container) {
	container.Bind("printer", func(container di.Container) interface{} {
		return app.PrinterService{
			Message: "Welcome to the Gimli framework!",
		}
	})

	container.Bind("controllers.welcome", func(container di.Container) interface{} {
		return &controllers.WelcomeController{
			Printer: container.MustResolve("printer").(app.PrinterService),
		}
	})
}
