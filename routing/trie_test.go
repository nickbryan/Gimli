package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouteTrie(t *testing.T) {
	assert.Implements(t, (*routeTrie)(nil), newRouteTrie())
}

type routePair struct {
	path  string
	route *Route
}

var routePairs []*routePair = []*routePair{
	{"/", NewRoute("/", nil, nil)},
	{"/path/to/route", NewRoute("/path/to/route", nil, nil)},
	{"/path/to/route/Nick", NewRoute("/path/to/route/:name", nil, nil)},
	{"/path/to/route/123/something", NewRoute("/path/to/route/:name/something", nil, nil)},
	{"/path/to/route/123/somethingelse", NewRoute("/path/to/route/:name/somethingelse", nil, nil)},
	{"/different", NewRoute("/different", nil, nil)},
	{"/different/path", NewRoute("/different/path", nil, nil)},
	{"/another/your-face/or/something/true", NewRoute("/another/:path/or/something/:else", nil, nil)},
}

func TestNilAndEmptyMapIsReturnedIfRouteNotFound(t *testing.T) {
	trie := newRouteTrie()

	for _, routePair := range routePairs {
		trie.add(routePair.route)
	}

	routes, params := trie.search("/path/to")
	assert.Nil(t, routes)
	assert.Empty(t, params)
}

func TestRoutesCanBeMatchedByPath(t *testing.T) {
	trie := newRouteTrie()

	for _, routePair := range routePairs {
		trie.add(routePair.route)

		routes, _ := trie.search(routePair.path)
		assert.Equal(t, routePair.route, routes[0])
	}
}

func TestMultipleRoutsCanBeAddedWithTheSamePath(t *testing.T) {
	trie := newRouteTrie()

	route1 := NewRoute("/path/to/route", []string{"GET"}, nil)
	route2 := NewRoute("/path/to/route", []string{"POST"}, nil)

	trie.add(route1)
	trie.add(route2)

	routes, _ := trie.search("/path/to/route")
	assert.Equal(t, []*Route{route1, route2}, routes)
}

func TestNamedParamsAreReturnedWhenRoutesAreFound(t *testing.T) {
	trie := newRouteTrie()

	route := NewRoute("/path/to/route/:name/with/:id", nil, nil)
	trie.add(route)

	_, params := trie.search("/path/to/route/nick/with/123")
	assert.Equal(t, map[string]string{"name": "nick", "id": "123"}, params)
}

func TestSearchCanHandleAPathThatStartsWithoutASlash(t *testing.T) {
	trie := newRouteTrie()

	route := NewRoute("/path/to/route", nil, nil)
	trie.add(route)

	routes, _ := trie.search("path/to/route")
	assert.Equal(t, route, routes[0])
}
