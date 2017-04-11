---
layout: documentation
title: Defining Routes
---
# Defining Routes

- [Introduction](#introduction)
- [Basic Routing](#basic-routing)

<a class="anchor" id="introduction"></a>
## Introduction
Currently, the simplest way to define your application routes is in the `bootstrap/routes.go` file. There are a number of 
helper functions that make adding routes to the router a breeze which we will cover soon. The 
[routing package](/docs/0.1/routing) provides other ways of defining routes that give you more control if you need it but 
most applications will be fine with the following.

<a class="anchor" id="basic-routing"></a>
## Basic Routing
The routing package exposes a helper method for each type of request method, along with other helpers if you need some 
more control:

```go
package yours

import (
        "net/http"
        
        . "github.com/nickbryan/gimli/routing"
)

func main() {
    callback := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
            rw.Write([]byte("The" + r.Method + "handler was called."))
    })
    
    router := NewRouter()
    
    router.Get("/", callback) // prints "The GET handler was called."
    router.Post("/", callback) // prints "The POST handler was called."
    router.Put("/", callback) // prints "The PUT handler was called."
    router.Patch("/", callback) // prints "The PATCH handler was called."
    router.Delete("/", callback) // prints "The DELETE handler was called."
    
    router.Any("/any", callback) // Responds to any request method
    router.Match("/match", callback, http.MethodGet, http.MethodDelete) // Only responds to GET and DELETE request methods
}

```