package commands

import (
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/urfave/cli"
)

type newCommand struct {
	path       string
	filesystem afero.Fs
}

func New(path string, filesystem afero.Fs) Runner {
	n := &newCommand{
		path: path,
	}

	if filesystem == nil {
		filesystem = afero.NewOsFs()
	}

	n.filesystem = filesystem

	return n
}

func (new *newCommand) Run() error {
	goPath := strings.TrimRight(os.Getenv("GOPATH"), "/") + "/"

	if _, err := new.filesystem.Stat(goPath + "src/" + new.path); os.IsNotExist(err) == false {
		return cli.NewExitError("Project [testapp] already exists in [github.com/nickbryan/].", 1)
	}

	return nil
}
