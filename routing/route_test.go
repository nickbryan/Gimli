package routing

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouteSetsParamsAsExpected(t *testing.T) {
	r := NewRoute("/path/to/route", []string{"get", "post"})

	assert.IsType(t, new(Route), r)
	assert.Equal(t, "/path/to/route", r.Path())
	assert.Equal(t, []string{"GET", "POST"}, r.Methods())
}

func TestPathIsFormattedWhenSet(t *testing.T) {
	r := NewRoute("   ////path/to/route", nil)
	assert.Equal(t, "/path/to/route", r.Path())

	r.SetPath("path/to/route")
	assert.Equal(t, "/path/to/route", r.Path())
}

func TestMethodsAreConvertedToNetHttpFormatWhenSet(t *testing.T) {
	r := NewRoute("/path/to/route", nil)

	r.SetMethods("get", "head", "post", "PUT", "PATCH", "DELETE")
	assert.Equal(t, []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	}, r.Methods())
}
