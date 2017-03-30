package bootstrap

import (
	"net/http"

	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/foundation/skeleton/app/controllers"
	"github.com/nickbryan/gimli/routing"
)

func init() {
	container := di.GetInstance()
	router := container.MustResolve("router").(routing.Router)

	router.Get("/", http.HandlerFunc(
		container.MustResolve("controllers.welcome").(*controllers.WelcomeController).Welcome,
	))
}
