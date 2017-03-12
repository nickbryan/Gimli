package providers

import (
	"github.com/nickbryan/framework/di"
	"github.com/nickbryan/framework/routing"
)

type RoutingProvider struct{}

func (p *RoutingProvider) Register(container di.Container) {
	container.Singleton("router", func(container di.Container) interface{} {
		return routing.NewRouter()
	})
}
