package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouteTree(t *testing.T) {
	assert.Implements(t, (*routeTree)(nil), newRouteTree())
}

type routePair struct {
	path  string
	route *Route
}

var routePairs []*routePair = []*routePair{
	{"/", NewRoute("/", nil)},
	{"/path/to/route", NewRoute("/path/to/route", nil)},
	{"/path/to/route/Nick", NewRoute("/path/to/route/:name", nil)},
	{"/path/to/route/123/something", NewRoute("/path/to/route/:name/something", nil)},
	{"/path/to/route/123/somethingelse", NewRoute("/path/to/route/:name/somethingelse", nil)},
	{"/different", NewRoute("/different", nil)},
	{"/different/path", NewRoute("/different/path", nil)},
	{"/another/your-face/or/something/true", NewRoute("/another/:path/or/something/:else", nil)},
}

func TestNilAndEmptyMapIsReturnedIfRouteNotFound(t *testing.T) {
	tree := newRouteTree()

	for _, routePair := range routePairs {
		tree.add(routePair.route)
	}

	routes, params := tree.search("/path/to")
	assert.Nil(t, routes)
	assert.Empty(t, params)
}

func TestRoutesCanBeMatchedByPath(t *testing.T) {
	tree := newRouteTree()

	for _, routePair := range routePairs {
		tree.add(routePair.route)

		routes, _ := tree.search(routePair.path)
		assert.Equal(t, routePair.route, routes[0])
	}
}

func TestMultipleRoutsCanBeAddedWithTheSamePath(t *testing.T) {
	tree := newRouteTree()

	route1 := NewRoute("/path/to/route", []string{"GET"})
	route2 := NewRoute("/path/to/route", []string{"POST"})

	tree.add(route1)
	tree.add(route2)

	routes, _ := tree.search("/path/to/route")
	assert.Equal(t, []*Route{route1, route2}, routes)
}

func TestNamedParamsAreReturnedWhenRoutesAreFound(t *testing.T) {
	tree := newRouteTree()

	route := NewRoute("/path/to/route/:name/with/:id", []string{"GET"})
	tree.add(route)

	_, params := tree.search("/path/to/route/nick/with/123")
	assert.Equal(t, map[string]string{"name": "nick", "id": "123"}, params)
}
