package bootstrap

import (
	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/foundation/skeleton/app/providers"
)

func init() {
	container := di.GetInstance()

	container.Register(&providers.ControllerServiceProvider{})
}
