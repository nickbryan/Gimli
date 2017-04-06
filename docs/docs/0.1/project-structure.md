---
layout: documentation
title: Project Structure
---
# Project Structure

- [Introduction](#introduction)
- [The Skeleton Project](#skeleton-project)
  - [The `app` Package](#app-package)
  - [The `bootstrap` Package](#bootstrap-package)
  - [The `config` Directory](#config-directory)
  - [The `public` Directory](#public-directory)

<a class="anchor" id="introduction"></a>
## Introduction
Gimli aims to provide a solid foundation for getting started with a new project by providing a skeleton 
application out of the box. You can [read more on creating a new skeleton project](/docs/0.1/installation#creating-a-project) 
in the installation section. This skeleton is just a starting point so feel free to customise it however you like.
 
Using the skeleton is a nice way to get started but there are times when you will need more flexibility with your project 
structure. Gimli has tried to keep this in mind by taking a modular approach to framework design, splitting each component 
into a self contained package. So you can build your own project skeleton using only the packages that you need.

The foundation package is the heart of the framework and responsible for tying all of the other Gimli components together.
All of the other packages aim to be self contained and do not require any other parts of the framework to function.
        
<a class="anchor" id="skeleton-project"></a>
## The Skeleton Project
The default skeleton project structure looks like this:

    foundation/skeleton
    ├── app
    │   ├── controllers
    │   │   └── welcome.go
    │   ├── printer.go
    │   └── providers
    │       └── controller.go
    ├── bootstrap
    │   ├── bootstrap.go
    │   ├── providers.go
    │   └── routes.go
    ├── config
    │   └── app.json
    ├── main.go
    └── public
        ├── css
        ├── favicon.ico
        ├── img
        ├── js
        └── robots.txt
        
The projects main package consists of a single `main.go` file which includes the `bootstrap` package and runs the application.

<a class="anchor" id="app-package"></a>
### The App Package
The `app` package is where the main application code lives. You can split this up as you see fit but some default 
sub-packages have been created as a starting point.

<a class="anchor" id="bootstrap-package"></a>
### The Bootstrap Package
The bootstrap package is where all application bootstrapping should take place. 

The file `bootstrap.go` is responsible for pulling in the `github.co/nickbryan/gimli/foundation` package and initialising 
the application. This will create the container instance and register the default service providers.
 
The `providers.go` file has been created to provide a convenient place to register custom service providers with the 
container. This should be done in the `init()` function so that they are registered before `main.go` calls `Application.Run()`.

Similarly to `providers.go`, a `routes.go` file is created as a convenient place to register all route callbacks for the 
application. Again, this should be done in the `init()` function.

<a class="anchor" id="config-directory"></a>
### The Config Directory
Currently Gimli only supports `.json` files as a configuration format.

When Gimli is booted a `github.com/nickbryan/gimli/config.Repository` is loaded into the container and all `.json` files 
are loaded from the `config` directory.

You can [read more about configuration](/docs/0.1/configuration) in the Bootstrapping section of the documentation.

<a class="anchor" id="public-directory"></a>
### The Public Directory
The `public` directory is where all application assets should be stored and has been broken down into a common structure.
The router will look in here by default for all static files.
