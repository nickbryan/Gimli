package commands

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var goPath = "/home/nickbryan/go"

func TestErrorIsReturnedIfProjectAlreadyExists(t *testing.T) {
	os.Setenv("GOPATH", "/home/nickbryan/go")

	filesystem := afero.NewMemMapFs()

	err := filesystem.MkdirAll(goPath+"/src/github.com/nickbryan/testapp", os.ModePerm)
	if err != nil {
		panic(err)
	}

	command := New("github.com/nickbryan/testapp", filesystem).Run()
	assert.EqualError(t, command, "Project [testapp] already exists in [github.com/nickbryan/].")
}
