---
layout: documentation
title: Service Container
---
# Service Container

- [Introduction](#introduction)
- [Service Providers](#service-provider)
- [Binding](#binding)
- [Resolving](#resolving)

<a class="anchor" id="introduction"></a>
## Introduction
Gimli provides a basic service container implementation trough the `di` package. The container does not aim to be "magic" 
and is a simple way to manage object dependencies. You must first bind any values you wish to use in the container:

```go
package main

import . "github.com/nickbryan/gimli/di"

func main() {
        container := NewContainer()
        
        container.Bind("name", func(container Container) interface{} {
                return "Nick"
        })
}
```

There a few different ways you can bind values in the container which will be covered in the next section. Once your 
values are bound in the container you can access them as follows:
```go
package main

import . "github.com/nickbryan/gimli/di"

func main() {
        container := NewContainer()
        
        container.Bind("name", func(container Container) interface{} {
                return "Nick"
        })
        
        println(container.MustResolve("name")) // Prints "Nick" or panics of "name" is not bound.
}
```

<a class="anchor" id="service-providers"></a>
## Service Providers
Service providers aim to provide a centralised way of encapsulating and organising your application bootstrapping with 
the container. A simple example of a service provider in action is the routing provider used by the `foundation` package:
```go
package providers

import (
	"github.com/nickbryan/gimli/di"
	"github.com/nickbryan/gimli/routing"
)

// RoutingProvider sets the router in the container.
type RoutingProvider struct{}

// Register a new router in the container.
func (p *RoutingProvider) Register(container di.Container) {
	container.Bind("router", func(container di.Container) interface{} {
		return routing.NewRouter()
	})
}
```

The `foundation` package then registers this with the container as follows:

    container.Register(&providers.RoutingProvider{})
    
If you have created your project from the supplied skeleton you can do all the registering of your applications service 
providers in the `bootstrap/providers.go` file.

<a class="anchor" id="binding"></a>
## Binding
### Bind
The most common way to bind a service on the container is to use the `bind` method. The first parameter should be a string 
that represents the nature of the service; used to resolve the service later. The second parameter should be a closure that 
takes the container as its only parameter and returns an instance of the service.
```go 
container.Bind("Notifier", func(container di.Container) interface{} {
    return &Notifier{
        Mailer: container.MustResolve("EmailService"),
    }
})
```
Receiving the container as an argument to the resolver allows us to build up complex structs that require other services 
bound in the container.

When a service is bound in the container via bind, the same instance of the service will be returned on each resolution.

### Instance
### Factory

<a class="anchor" id="resloving"></a>
## Resolving
### Resolve
### MustResolve