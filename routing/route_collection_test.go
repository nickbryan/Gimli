package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouteCollection(t *testing.T) {
	rc := NewRouteCollection()

	assert.IsType(t, new(RouteCollection), rc)
}
