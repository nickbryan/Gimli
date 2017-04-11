---
layout: documentation
title: Service Container
---
# Service Container

- [Introduction](#introduction)
- [Service Providers](#service-providers)
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
The most common way to bind a service on the container is to use the `Bind` method. The first parameter should be a string 
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
You can add an existing object instance into the container via the `Instance` method. Once bound, resolving will always 
return the same instance:
```go
notifier := NewNotifier()
container.Instance("Notifier", notifier)
```
### Factory
As mentioned above, values are shared by default when bound to the container. Sometimes you may want a new instance of 
an object on each resolution. This can be achieved with the `Factory` method:
```go
container.Factory("SessionStore", func(container di.Container) interface{} {
        return NewSessionStorage()
})
```

<a class="anchor" id="resloving"></a>
## Resolving
### Resolve
Once a value has been bound to the container you can retrieve it via the `Resolve` method. The requested value and `nil` 
will be returned if the value exists in the container, otherwise `nil` and an `error` will be returned:
```go
package yours

import . "github.com/nickbryan/gimli/di"

func main() {
        container := NewContainer()
      
        container.Instance("TheMeaningOfLife", 42)
  
        val, err := container.Resolve("TheMeaningOfLife")
        if err != nil {
                panic(err)
        }
    
        println("The meaning of life is: " + val.(string))
}
```

### MustResolve
There is a short hand way of resolving from the container using `MustResolve`. If the value does not exist, `MustResolve` 
will `panic`. You can use this as follows:
```go
package yours

import . "github.com/nickbryan/gimli/di"

func main() {
        container := NewContainer()
      
        container.Instance("TheMeaningOfLife", 42)
    
        println("The meaning of life is: " + container.MustResolve("TheMeaningOfLife").(string))
}
```

Or you could catch the panic as follows:
```go
package yours

import . "github.com/nickbryan/gimli/di"

func main() {
        defer func() {
                if r := recover(); r != nil {
                    println("Recovered from panic", r)
                }
        }()

        container := NewContainer()
        
        println("The meaning of life is: " + container.MustResolve("TheMeaningOfLife").(string))
}
```