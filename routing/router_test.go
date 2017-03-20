package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestNewRouter(t *testing.T) {
	assert.Implements(t, (*Router)(nil), NewRouter())
}

func TestRouterCanBePassedAsHttpHandler(t *testing.T) {
	assert.Implements(t, (*http.Handler)(nil), NewRouter())
}
