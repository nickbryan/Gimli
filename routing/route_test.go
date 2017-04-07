package routing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouteSetsParamsAsExpected(t *testing.T) {
	r := NewRoute("/path/to/route", []string{"get", "post"}, nil)

	assert.IsType(t, new(Route), r)
	assert.Equal(t, "/path/to/route", r.Path())
	assert.Equal(t, []string{"GET", "POST"}, r.Methods())
}

func TestMethodIsSetToGETIfMethodsPassedAsNil(t *testing.T) {
	r := NewRoute("", nil, nil)
	assert.Equal(t, []string{"GET"}, r.Methods())
}

func TestPathIsFormattedWhenSet(t *testing.T) {
	r := NewRoute("   ////path/to/route", nil, nil)
	assert.Equal(t, "/path/to/route", r.Path())

	r.SetPath("path/to/route")
	assert.Equal(t, "/path/to/route", r.Path())

	r.SetPath("")
	assert.Equal(t, "/", r.Path())
}

func TestMethodsAreConvertedToNetHttpFormatWhenSet(t *testing.T) {
	r := NewRoute("", nil, nil)

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

func TestNameIsNormalisedWhenSet(t *testing.T) {
	r := NewRoute("", nil, nil)

	r.SetName(" HoMe ")
	assert.Equal(t, "home", r.Name())
}

func TestRouteCaBeFilteredThroughMatchers(t *testing.T) {
	r := NewRoute("/test", []string{http.MethodGet}, nil)
	r.AddMatcher(MatcherFunc(MethodMatcher))

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	assert.True(t, r.Matches(request))

	request = httptest.NewRequest(http.MethodPost, "/", nil)
	assert.False(t, r.Matches(request))
}

func TestMethodMatcherIsSetByDefault(t *testing.T) {
	r := NewRoute("/test", []string{http.MethodGet}, nil)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	assert.True(t, r.Matches(request))

	request = httptest.NewRequest(http.MethodPost, "/", nil)
	assert.False(t, r.Matches(request))
}

func TestSlashesAreRemovedFromEndOfPathWhenSet(t *testing.T) {
	r := NewRoute("test/", nil, nil)
	assert.Equal(t, "/test", r.Path())
}
