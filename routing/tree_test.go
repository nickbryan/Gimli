package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type routePair struct {
	path  string
	route *Route
}

var routePairs []*routePair = []*routePair{
	{"/path/to/route", NewRoute("/path/to/route", nil)},
	{"/path/to/route/Nick", NewRoute("/path/to/route/:name", nil)},
	{"/path/to/route/123/something", NewRoute("/path/to/route/:name/something", nil)},
	{"/different", NewRoute("/different", nil)},
	{"/different/path", NewRoute("/different/path", nil)},
	{"/another/your-face/or/something/true", NewRoute("/another/:path/or/something/:else", nil)},
}

func TestRoutesCanBeMatchedByPath(t *testing.T) {
	tr := newTree()

	for _, routePair := range routePairs {
		tr.addRoute(routePair.route)

		rt, _ := tr.findRoute(routePair.path)
		assert.Equal(t, routePair.route, rt)
	}
}
