package providers

import (
	"errors"
	"path"
	"runtime"
	"testing"

	"github.com/nickbryan/gimli/config"
	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/routing"
	"github.com/stretchr/testify/assert"
)

func TestRoutingProviderBindsRouterInContainer(t *testing.T) {
	container := di.NewContainer()
	(&RoutingProvider{}).Register(container)
	assert.Implements(t, (*routing.Router)(nil), container.MustResolve("router"))
}

func TestConfigProviderReadsValuesFromFiles(t *testing.T) {
	container := di.NewContainer()

	_, file, _, ok := runtime.Caller(0)
	if ok == false {
		assert.Fail(t, "Could not read caller information")
	}
	container.Instance("path.config", path.Dir(file)+"/test_assets/config")

	(&ConfigurationProvider{}).Register(container)
	assert.IsType(t, new(config.Repository), container.MustResolve("config"))

	conf := container.MustResolve("config").(*config.Repository)
	assert.Equal(t, "Hello", conf.Get("val1"))
	assert.Equal(t, "World", conf.Get("val2"))
}

func TestCheckErrPanicsWhenErrorIsNotNil(t *testing.T) {
	assert.Panics(t, func() {
		checkError(errors.New("SomeErrorMessage"))
	})
}
