package providers

import (
	"github.com/nickbryan/Gimli/di"
	"github.com/nickbryan/Gimli/routing"
)

type RoutingProvider struct{}

func (p *RoutingProvider) Register(container di.Container) {
	container.Bind("router", func(container di.Container) interface{} {
		return routing.NewRouter()
	})
}
