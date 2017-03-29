package commands

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/urfave/cli"
)

const skeletonPath string = "github.com/nickbryan/gimli/foundation/skeleton"

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

func (command *newCommand) Run() error {
	goPath := strings.TrimRight(os.Getenv("GOPATH"), "/") + "/"
	projectPath := goPath + "src/" + command.path

	if ok, _ := afero.DirExists(command.filesystem, projectPath); ok {
		return cli.NewExitError("Project already exists at ["+projectPath+"].", 1)
	}

	err := command.copyDir(goPath+"src/"+skeletonPath, projectPath)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}

func (command *newCommand) copyDir(src string, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = command.filesystem.MkdirAll(dst, si.Mode())
	if err != nil {
		return nil
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return nil
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = command.copyDir(srcPath, dstPath)
			if err != nil {
				return nil
			}
		} else {
			err = command.copyFile(srcPath, dstPath)
			if err != nil {
				return nil
			}
		}
	}

	return nil
}

func (command *newCommand) copyFile(src, dst string) error {
	in, err := command.filesystem.Open(src)
	if err != nil {
		return nil
	}
	defer in.Close()

	out, err := command.filesystem.Create(dst)
	if err != nil {
		return nil
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return nil
	}

	err = out.Sync()
	if err != nil {
		return nil
	}

	si, err := os.Stat(src)
	if err != nil {
		return nil
	}
	err = command.filesystem.Chmod(dst, si.Mode())
	if err != nil {
		return nil
	}

	return nil
}
