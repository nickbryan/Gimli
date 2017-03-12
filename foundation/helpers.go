package foundation

import (
	"github.com/nickbryan/framework/config"
	"github.com/nickbryan/framework/di"
	"github.com/nickbryan/framework/routing"
)

func App() Application {
	if app, ok := di.GetInstance().(Application); ok {
		return app
	}

	panic("Container singleton is not an instance of foundation.Application")
}

func Config() *config.Repository {
	conf, err := App().Make("config")
	if err != nil {
		panic(err)
	}

	return conf.(*config.Repository)
}

func Router() routing.Router {
	router, err := App().Make("router")
	if err != nil {
		panic(err)
	}

	return router.(routing.Router)
}
