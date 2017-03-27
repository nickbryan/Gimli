package commands

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var goPath = "/home/nickbryan/go"

func TestErrorIsReturnedIfProjectAlreadyExists(t *testing.T) {
	os.Setenv("GOPATH", goPath)
	filesystem := afero.NewMemMapFs()

	err := filesystem.MkdirAll(goPath+"/src/github.com/nickbryan/testapp", os.ModePerm)
	if err != nil {
		panic(err)
	}

	command := New("github.com/nickbryan/testapp", filesystem).Run()
	assert.EqualError(t, command, "Project [testapp] already exists in [github.com/nickbryan/].")
}

func TestSkeletonIsCopiedToPathIfNoError(t *testing.T) {
	os.Setenv("GOPATH", goPath)
	base := afero.NewOsFs()
	roBase := afero.NewReadOnlyFs(base)
	filesystem := afero.NewCopyOnWriteFs(roBase, afero.NewMemMapFs())
	path := goPath + "/src/github.com/nickbryan/testapp"

	command := New("github.com/nickbryan/testapp", filesystem).Run()
	assert.Nil(t, command)

	ok, _ := afero.DirExists(filesystem, path)
	assert.True(t, ok, "Skeleton has not been copied to path")
}
