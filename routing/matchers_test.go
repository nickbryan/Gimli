package routing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethodMatcher(t *testing.T) {
	route := NewRoute("/test", []string{http.MethodGet})
	matcher := MatcherFunc(MethodMatcher)

	request := httptest.NewRequest(http.MethodGet, "/test", nil)
	assert.True(t, matcher.Match(route, request))

	request = httptest.NewRequest(http.MethodPost, "/test", nil)
	assert.False(t, matcher.Match(route, request))
}
