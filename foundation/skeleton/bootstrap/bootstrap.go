package bootstrap

import (
	"github.com/nickbryan/gimli/foundation"
	_ "github.com/nickbryan/gimli/foundation/skeleton/app"
)

const BasePath = "YOUR_APPLICATION_BASE_PATH"

var application foundation.Application

func init() {
	application = foundation.NewApplication(BasePath)
}

func Application() foundation.Application {
	return application
}
