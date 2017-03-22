package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestNewRouteCollection(t *testing.T) {
	rc := NewRouteCollection()

	assert.IsType(t, new(RouteCollection), rc)
}

func TestRouteCanBeAddedToCollection(t *testing.T) {
	rc := NewRouteCollection()

	route := NewRoute("/", nil, nil)
	rc.Add(route)

	routes, _ := rc.routes.search("/")
	assert.Contains(t, routes, route)
}

func TestNamedRouteIsAddedToRoutesAndNamedRouteMap(t *testing.T) {
	rc := NewRouteCollection()

	route := NewRoute("/", nil, nil)
	route.SetName("home")

	rc.Add(route)

	routes, _ := rc.routes.search("/")
	assert.Contains(t, routes, route)

	assert.Contains(t, rc.namedRoutes, route.Name())
	assert.Equal(t, route, rc.namedRoutes[route.Name()])
}

// TODO: test that named routes cannot be overridden

func TestNonNamedRouteIsNotAddedToNamedRouteMap(t *testing.T) {
	rc := NewRouteCollection()

	route := NewRoute("/", nil, nil)

	assert.NotContains(t, rc.namedRoutes, route.Name())
}

func TestNamedAndNonNamedRoutesAreAddedToAllRoutes(t *testing.T) {
	rc := NewRouteCollection()

	route := NewRoute("/test", nil, nil)
	namedRoute := NewRoute("/", nil, nil)
	namedRoute.SetName("home")

	rc.Add(route)
	rc.Add(namedRoute)

	assert.Contains(t, rc.allRoutes, route)
	assert.Contains(t, rc.allRoutes, namedRoute)
}

func TestNamedRoutesCanBeUpdatedInCollection(t *testing.T) {
	rc := NewRouteCollection()

	route := NewRoute("/", nil, nil)

	rc.Add(route)

	assert.NotContains(t, rc.namedRoutes, route.Name())

	route.SetName("home")
	rc.RefreshNamedRoutes()

	assert.Contains(t, rc.namedRoutes, route.Name())
	assert.Equal(t, route, rc.namedRoutes[route.Name()])
}

func TestCountReturnsNumberOfRoutesInCollection(t *testing.T) {
	rc := NewRouteCollection()

	route1 := NewRoute("/", nil, nil)
	route2 := NewRoute("/a", nil, nil)
	route3 := NewRoute("/b", nil, nil)

	rc.Add(route1)
	rc.Add(route2)
	rc.Add(route3)

	assert.Equal(t, 3, rc.Count())
}

func TestHas(t *testing.T) {
	rc := NewRouteCollection()
	route := NewRoute("", nil, nil)
	route2 := NewRoute("", []string{http.MethodPost}, nil)

	assert.False(t, rc.Has(route))
	rc.Add(route)
	assert.True(t, rc.Has(route))

	assert.False(t, rc.Has(route2))
	rc.Add(route2)
	assert.True(t, rc.Has(route2))
}
