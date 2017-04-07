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

func TestRoutesByPathReturnsNilIfNoRoutesFound(t *testing.T) {
	rc := NewRouteCollection()

	assert.Nil(t, rc.RoutesByPath("some/test/route"))
}

func TestPrefixCanBeAddedToAllRoutesInCollection(t *testing.T) {
	rc := NewRouteCollection()

	routeA := NewRoute("/a", nil, nil)
	routeB := NewRoute("/b", nil, nil)
	rc.Add(routeA)
	rc.Add(routeB)

	rc.Prefix("prefix")

	assert.Equal(t, routeA, rc.RoutesByPath("/prefix/a").Routes[0])
	assert.Equal(t, routeB, rc.RoutesByPath("/prefix/b").Routes[0])
}

func TestCollectionsCanBeAddedToCollections(t *testing.T) {
	rc1 := NewRouteCollection()
	rc1.Add(NewRoute("/a", nil, nil))
	rc1.Add(NewRoute("/b", nil, nil))

	rc2 := NewRouteCollection()
	rc2.Add(NewRoute("/c", nil, nil))
	rc2.Add(NewRoute("/d", nil, nil))

	rc1.AddCollection(rc2)

	assert.NotNil(t, rc1.RoutesByPath("/a"))
	assert.NotNil(t, rc1.RoutesByPath("/b"))
	assert.NotNil(t, rc1.RoutesByPath("/c"))
	assert.NotNil(t, rc1.RoutesByPath("/d"))
}
