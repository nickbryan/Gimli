package commands

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var goPath = os.Getenv("GOPATH")

func TestDefaultFileSystemIsSetWhenPassedAsNil(t *testing.T) {
	command := New("github.com/nickbryan/testapp", nil)
	assert.Equal(t, afero.NewOsFs(), command.filesystem)
}

func TestErrorIsReturnedIfProjectAlreadyExists(t *testing.T) {
	filesystem := afero.NewMemMapFs()

	err := filesystem.MkdirAll(goPath+"/src/github.com/nickbryan/testapp", os.ModePerm)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	err = New("github.com/nickbryan/testapp", filesystem).Run()
	assert.EqualError(t, err, "Project already exists at ["+goPath+"/src/github.com/nickbryan/testapp"+"].")
}

func TestSkeletonIsCopiedToPathIfNoError(t *testing.T) {
	base := afero.NewOsFs()
	roBase := afero.NewReadOnlyFs(base)
	filesystem := afero.NewCopyOnWriteFs(roBase, afero.NewMemMapFs())
	path := goPath + "/src/github.com/nickbryan/testapp"

	err := New("github.com/nickbryan/testapp", filesystem).Run()
	assert.Nil(t, err)

	ok, _ := afero.DirExists(filesystem, path)
	assert.True(t, ok, "Skeleton has not been copied to path")
}
