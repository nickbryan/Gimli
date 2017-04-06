package bootstrap

import (
	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/foundation/skeleton/app/providers"
)

func init() {
	container := Application.Container()

	container.Register(&providers.ControllerServiceProvider{})
}
