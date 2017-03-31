package foundation

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/nickbryan/gimli/config"
	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/foundation/providers"
	"github.com/nickbryan/gimli/routing"
)

// VERSION of the application.
const VERSION = "0.1.1"

// Application is the heart of the Gimli framework. It is responsible for all bootstrapping and general
// functionality.
type Application interface {
	Container() di.Container
	Run()

	SetBasePath(basePath string)
	BasePath() string
	BootstrapPath() string
	ConfigPath() string
	Path() string
	PublicPath() string

	Environment() string
	IsEnvironment(env string) bool
}

type application struct {
	container di.Container
	basePath  string
}

// NewApplication creates a new Application instance, set the relevant paths in the container and
// register base bindings and providers.
func NewApplication(basePath string) Application {
	app := &application{
		container: di.NewContainer(),
	}

	app.SetBasePath(basePath)

	app.registerBaseBindings()
	app.registerBaseProviders()

	return app
}

// Run starts a http server running by calling http.ListenAndServe. It uses the host and port
// set in the app.json config.
func (app *application) Run() {
	conf := app.container.MustResolve("config").(*config.Repository)
	host, port := conf.Get("host").(string), conf.Get("port").(string)

	http.ListenAndServe(host+":"+port, app.container.MustResolve("router").(routing.Router))
}

func (app *application) registerBaseBindings() {
	app.container.Instance("app", app)
	app.container.Instance("container", app.container)
	di.SetInstance(app.container)
}

func (app *application) registerBaseProviders() {
	app.container.Register(&providers.ConfigurationProvider{})
	app.container.Register(&providers.RoutingProvider{})
}

func (app *application) Container() di.Container {
	return app.container
}

// SetBasePath sets the application base path and registers all other application paths
// in the container.
func (app *application) SetBasePath(basePath string) {
	// TODO: Test path.clean
	app.basePath = strings.TrimRight(basePath, `\/`)

	app.bindPathsInContainer()
}

// BasePath is the path to the root of the application.
func (app *application) BasePath() string {
	return app.basePath
}

// BootstrapPath is the path to the bootstrap package.
func (app *application) BootstrapPath() string {
	return app.basePath + string(filepath.Separator) + "bootstrap"
}

// BootstrapPath is the path to the config directory.
func (app *application) ConfigPath() string {
	return app.basePath + string(filepath.Separator) + "config"
}

// Path is the path to the app package.
func (app *application) Path() string {
	return app.basePath + string(filepath.Separator) + "app"
}

// PublicPath is the path to the public directory.
func (app *application) PublicPath() string {
	return app.basePath + string(filepath.Separator) + "public"
}

func (app *application) bindPathsInContainer() {
	app.container.Instance("path", app.Path())
	app.container.Instance("path.base", app.BasePath())
	app.container.Instance("path.bootstrap", app.BootstrapPath())
	app.container.Instance("path.config", app.ConfigPath())
	app.container.Instance("path.public", app.PublicPath())
}

// Environment will indicate the current environment.
func (app *application) Environment() string {
	return app.container.MustResolve("env").(string)
}

// IsEnvironment can be used to check for a specific environment.
func (app *application) IsEnvironment(env string) bool {
	return app.Environment() == env
}
