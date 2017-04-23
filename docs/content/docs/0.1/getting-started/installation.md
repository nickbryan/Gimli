+++
date = "2017-04-03T20:46:58+01:00"
draft = false 
title = "Installation"
[menu.main]
  parent = "Getting Started"
  identifier = "docs/0.1/getting-started/installation"
  weight = -10
+++

- [Server Requirements](#server-requirements)
- [Installing Gimli](#installing-gimli)
- [Creating a Project](#creating-a-project)

<a class="anchor" id="server-requirements"></a>
## Server Requirements
* Go version >= 1.7

<a class="anchor" id="installing-gimli"></a>
## Installing Gimli
You can use `go get` to install the Gimli framework and cli tool:

    $ go get github.com/nickbryan/gimli

Installation will make the following packages available:

    github.com/nickbryan/gimli/config
    github.com/nickbryan/gimli/di
    github.com/nickbryan/gimli/foundation
    github.com/nickbryan/gimli/routing
    
Alternately you can `go get` an individual package:

    $ go get github.com/nickbryan/gimli/routing
    
And import the package as below:
{{<highlight go>}}
package yours

import (
  "github.com/nickbryan/gimli/routing"
)
{{</highlight>}}

<a class="anchor" id="creating-a-project"></a>
## Creating a Project
Gimli comes with a command line utility for working with the framework. When you `go get` the package the cli tool 
is installed in the `bin/` directory of your `$GOPATH`.

If you have added the `bin/` directory to your `$PATH` you can run the following to validate the installation:

    $ gimli
    
You should see an output similar to the following:

    $ gimli
    NAME:
       gimli - A cli utility for managing gimli applications
    
    USAGE:
       gimli [global options] command [command options] [arguments...]
    
    VERSION:
       0.1.0
    
    DESCRIPTION:
       The gimli cli tool should be used to aid in the development of applications using the gimli framework.
    
    AUTHOR:
       Nick Bryan
    
    COMMANDS:
         new      creates a new gimli project
         help, h  Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
       --help, -h     show help
       --version, -v  print the version

If you don't have `bin/` in your path you can run:

    $ $GOPATH/bin/gimli

You can now create a new project with the following command:

    $ gimli new github.com/yourusername/projectname

This will create a new skeleton project in your `$GOPATH` under the specified import path. You should use the import path 
relative to where your project will be kept on version control.
