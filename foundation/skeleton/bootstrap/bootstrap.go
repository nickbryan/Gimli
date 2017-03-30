package bootstrap

import (
	"github.com/nickbryan/gimli/foundation"
)

const BasePath = "/home/mifdev/go/src/github.com/nickbryan/gimli/foundation/skeleton"

var application foundation.Application

func init() {
	application = foundation.NewApplication(BasePath)
}

func Application() foundation.Application {
	return application
}
