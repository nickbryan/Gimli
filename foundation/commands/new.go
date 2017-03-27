package commands

import (
	"os"
	"strings"

	//"fmt"
	"github.com/spf13/afero"
	"github.com/urfave/cli"
	"io"
	//"io/ioutil"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
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

func (command *newCommand) Run() error {
	goPath := strings.TrimRight(os.Getenv("GOPATH"), "/") + "/"
	projectPath := goPath + "src/" + command.path

	if ok, _ := afero.DirExists(command.filesystem, projectPath); ok {
		return cli.NewExitError("Project [testapp] already exists in [github.com/nickbryan/].", 1)
	}

	_, file, _, ok := runtime.Caller(0)
	if ok == false {
		return cli.NewExitError("Unable to get calling file information.", 1)
	}

	err := command.copyDir(path.Dir(path.Dir(file))+"/skeleton", projectPath)
	if err != nil {
		cli.NewExitError(err, 1)
	}

	return nil
}

func (command *newCommand) copyFile(src, dst string) (err error) {
	in, err := command.filesystem.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := command.filesystem.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = command.filesystem.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

func (command *newCommand) copyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = command.filesystem.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = command.filesystem.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = command.copyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = command.copyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}
