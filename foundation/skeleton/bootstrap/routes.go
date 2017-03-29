package bootstrap

import (
	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/foundation/skeleton/app"
	"github.com/nickbryan/gimli/routing"
)

func init() {
	router := di.GetInstance().MustResolve("router").(routing.Router)

	router.Get("/", &app.WelcomeHandler{"Welcome to the Gimli Framework!"})
}
