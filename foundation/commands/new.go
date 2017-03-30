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

func New(path string, filesystem afero.Fs) *newCommand {
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

	if err = command.replacePaths(projectPath); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}

func (command *newCommand) replacePaths(dir string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	command.replaceInFile(dir+"/main.go", "github.com/nickbryan/gimli/foundation/skeleton", command.path)
	command.replaceInFile(dir+"/bootstrap/routes.go", "github.com/nickbryan/gimli/foundation/skeleton", command.path)
	command.replaceInFile(dir+"/bootstrap/providers.go", "github.com/nickbryan/gimli/foundation/skeleton", command.path)
	command.replaceInFile(dir+"/app/controllers/welcome.go", "github.com/nickbryan/gimli/foundation/skeleton", command.path)
	command.replaceInFile(dir+"/app/providers/controller.go", "github.com/nickbryan/gimli/foundation/skeleton", command.path)

	command.replaceInFile(dir+"/bootstrap/bootstrap.go", "/home/mifdev/go/src/github.com/nickbryan/gimli/foundation/skeleton", dir)

	return
}

func (command *newCommand) replaceInFile(filePath, text, replace string) {
	file, err := command.filesystem.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	read, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	newContents := strings.Replace(string(read), text, replace, -1)

	file, err = command.filesystem.Create(filePath)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(newContents))
	if err != nil {
		panic(err)
	}
}

func (command *newCommand) copyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return
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
			err = command.copyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}

func (command *newCommand) copyFile(src, dst string) (err error) {
	in, err := command.filesystem.Open(src)
	if err != nil {
		return
	}
	defer func() {
		if e := in.Close(); e != nil {
			err = e
		}
	}()

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
