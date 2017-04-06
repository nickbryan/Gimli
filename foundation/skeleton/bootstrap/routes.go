package bootstrap

import (
	"net/http"

	"github.com/nickbryan/gimli/foundation/skeleton/app/controllers"
	"github.com/nickbryan/gimli/routing"
)

func init() {
	container := Application.Container()
	router := container.MustResolve("router").(routing.Router)

	router.Get("/", http.HandlerFunc(
		container.MustResolve("controllers.welcome").(*controllers.WelcomeController).Welcome,
	))
}
