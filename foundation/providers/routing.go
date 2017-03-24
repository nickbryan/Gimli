package providers

import (
	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/routing"
)

type RoutingProvider struct{}

func (p *RoutingProvider) Register(container di.Container) {
	container.Bind("router", func(container di.Container) interface{} {
		return routing.NewRouter()
	})
}
