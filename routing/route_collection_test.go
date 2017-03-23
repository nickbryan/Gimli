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

	route := NewRoute("/", nil, nil)
	rc.Add(route)

	compiledRoutes := rc.RoutesByPath("/")
	assert.Contains(t, compiledRoutes.Routes, route)
}

func TestRouteByNameReturnsNilIfRouteNotFound(t *testing.T) {
	rc := NewRouteCollection()
	assert.Nil(t, rc.RouteByName("routeName"))
}

func TestNamedRouteIsAddedToRoutesAndNamedRouteMap(t *testing.T) {
	rc := NewRouteCollection()

	route := NewRoute("/", nil, nil)
	route.SetName("home")

	rc.Add(route)

	compiledRoutes := rc.RoutesByPath("/")
	assert.Contains(t, compiledRoutes.Routes, route)

	assert.Equal(t, route, rc.RouteByName(route.Name()))
}

func TestNonNamedRouteIsNotAddedToNamedRouteMap(t *testing.T) {
	rc := NewRouteCollection()
	route := NewRoute("/", nil, nil)

	assert.Nil(t, rc.RouteByName(route.Name()))
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

	assert.Nil(t, rc.RouteByName(route.Name()))

	route.SetName("home")
	rc.RefreshNamedRoutes()

	assert.Equal(t, route, rc.RouteByName(route.Name()))
}

func TestCountReturnsNumberOfRoutesInCollection(t *testing.T) {
	rc := NewRouteCollection()

	rc.Add(NewRoute("/a", nil, nil))
	rc.Add(NewRoute("/b", nil, nil))
	rc.Add(NewRoute("/c", nil, nil))

	assert.Equal(t, 3, rc.Count())
}
