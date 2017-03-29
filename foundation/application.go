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

func NewApplication(basePath string) Application {
	app := &application{
		container: di.NewContainer(),
	}

	app.SetBasePath(basePath)

	app.registerBaseBindings()
	app.registerBaseProviders()

	return app
}

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

func (app *application) SetBasePath(basePath string) {
	// Test path.clean
	app.basePath = strings.TrimRight(basePath, `\/`)

	app.bindPathsInContainer()
}

func (app *application) BasePath() string {
	return app.basePath
}

func (app *application) BootstrapPath() string {
	return app.basePath + string(filepath.Separator) + "bootstrap"
}

func (app *application) ConfigPath() string {
	return app.basePath + string(filepath.Separator) + "config"
}

func (app *application) Path() string {
	return app.basePath + string(filepath.Separator) + "app"
}

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

func (app *application) Environment() string {
	return app.container.MustResolve("env").(string)
}

func (app *application) IsEnvironment(env string) bool {
	return app.Environment() == env
}
