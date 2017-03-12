package foundation

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/nickbryan/framework/di"
	"github.com/nickbryan/framework/foundation/providers"
)

const VERSION = "0.1.0"

type Application interface {
	di.Container

	BasePath() string
	BootstrapPath() string
	ConfigPath() string
	Path() string
	PublicPath() string
	SetBasePath(basePath string) Application
	Run()
	Environment() string
	IsEnvironment(env string) bool
}

type application struct {
	basePath string

	di.Container
}

func NewApplication(basePath string) Application {
	app := &application{
		Container: di.NewContainer(),
	}

	app.SetBasePath(basePath)

	app.registerBaseBindings()
	app.registerBaseProviders()

	return app
}

func (app *application) registerBaseBindings() {
	di.SetInstance(app)
	app.Instance("app", app)
	app.Instance("container", app)
}

func (app *application) registerBaseProviders() {
	app.Register(&providers.ConfigurationProvider{})
	app.Register(&providers.RoutingProvider{})
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

func (app *application) SetBasePath(basePath string) Application {
	app.basePath = strings.TrimRight(basePath, `\/`)

	app.bindPathsInContainer()

	return app
}

func (app *application) bindPathsInContainer() {
	app.Instance("path", app.Path())
	app.Instance("path.base", app.BasePath())
	app.Instance("path.bootstrap", app.BootstrapPath())
	app.Instance("path.config", app.ConfigPath())
	app.Instance("path.public", app.PublicPath())
}

func (app *application) Environment() string {
	env, err := App().Make("env")
	if err != nil {
		panic(err)
	}

	return env.(string)
}

func (app *application) IsEnvironment(env string) bool {
	return app.Environment() == env
}

func (app *application) Run() {
	host, port := Config().Get("host").(string), Config().Get("port").(string)

	http.ListenAndServe(host+":"+port, http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		Router().Parse(req, res)
	}))
}
