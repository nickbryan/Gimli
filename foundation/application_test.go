package foundation

import (
	"testing"

	"github.com/nickbryan/gimli/config"
	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/routing"
	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	assert.Implements(t, (*Application)(nil), NewApplication(""))
}

func TestBaseBindingsAreRegisteredOnApplicationInstantiation(t *testing.T) {
	app := NewApplication("")
	assert.True(t, app.Container().Has("app"))
	assert.True(t, app.Container().Has("container"))
	assert.Exactly(t, app.Container(), di.GetInstance())
}

func TestPathsAreSetInContainer(t *testing.T) {
	basePath := "/path/to/app"
	app := NewApplication(basePath)
	assert.Equal(t, basePath, app.Container().MustResolve("path.base"))
	assert.Equal(t, basePath+"/app", app.Container().MustResolve("path"))
	assert.Equal(t, basePath+"/bootstrap", app.Container().MustResolve("path.bootstrap"))
	assert.Equal(t, basePath+"/config", app.Container().MustResolve("path.config"))
	assert.Equal(t, basePath+"/public", app.Container().MustResolve("path.public"))

}

func TestBaseProvidersAreRegisteredOnApplicationInstantiation(t *testing.T) {
	app := NewApplication("")
	assert.IsType(t, new(config.Repository), app.Container().MustResolve("config"))
	assert.Implements(t, (*routing.Router)(nil), app.Container().MustResolve("router"))
}

func TestApplicationEnvironmentIsSetToProductionByDefault(t *testing.T) {
	app := NewApplication("")
	assert.Equal(t, "production", app.Environment())
	assert.True(t, app.IsEnvironment("production"))
}
