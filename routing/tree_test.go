package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var routes map[string]*Route = map[string]*Route{
	"/path/to/route":               NewRoute("/path/to/route", nil),
	"/path/to/route/Nick":          NewRoute("/path/to/route/:name", nil),
	"/path/to/route/123/something": NewRoute("/path/to/route/:name/something", nil),
}

//func TestRoutesCanBeMatchedByPath(t *testing.T) {
//	tr := newTree()
//
//	for path, route := range routes {
//		tr.addRoute(route)
//
//		rt, _ := tr.findRoute(path)
//		assert.Equal(t, route, rt)
//	}
//}

func TestRoutesCanBeMatchedByPath(t *testing.T) {
	tr := newTree()

	r1 := NewRoute("/path/to/route", nil)
	r2 := NewRoute("/path/to/route/:name", nil)
	r3 := NewRoute("/path/to/route/:name/something", nil)
	r4 := NewRoute("/base/in/face", nil)
	r5 := NewRoute("/base/no/place", nil)

	tr.addRoute(r1)
	tr.addRoute(r2)
	tr.addRoute(r3)
	tr.addRoute(r4)
	tr.addRoute(r5)

	rt, _ := tr.findRoute("/path/to/route")
	assert.Equal(t, r1, rt)

	rt, _ = tr.findRoute("/path/to/route/Nick")
	assert.Equal(t, r2, rt)

	rt, _ = tr.findRoute("/path/to/route/123/something")
	assert.Equal(t, r3, rt)

	rt, _ = tr.findRoute("/base/in/face")
	assert.Equal(t, r4, rt)

	rt, _ = tr.findRoute("/base/no/place")
	assert.Equal(t, r5, rt)s
}
