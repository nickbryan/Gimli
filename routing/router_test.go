package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	assert.Implements(t, (*Router)(nil), NewRouter())
}
