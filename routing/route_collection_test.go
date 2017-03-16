package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouteCollection(t *testing.T) {
	rc := NewRouteCollection()

	assert.IsType(t, new(RouteCollection), rc)
}

func TestRouteCanBeAddedToCollection(t *testing.T) {
	rc := NewRouteCollection()

	route := NewRoute("", nil)
	rc.Add(route)

	routes, _ := rc.routes.search("/")
	assert.Contains(t, routes, route)
}
