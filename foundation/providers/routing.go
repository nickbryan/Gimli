package providers

import (
	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/routing"
)

// RoutingProvider sets the router in the container.
type RoutingProvider struct{}

// Register a new router in the container.
func (p *RoutingProvider) Register(container di.Container) {
	container.Bind("router", func(container di.Container) interface{} {
		return routing.NewRouter()
	})
}
