package commands

import (
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/urfave/cli"
)

const skeletonPath string = ""

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

	if ok, _ := afero.DirExists(new.filesystem, goPath+"src/"+new.path); ok {
		return cli.NewExitError("Project [testapp] already exists in [github.com/nickbryan/].", 1)
	}

	return nil
}
