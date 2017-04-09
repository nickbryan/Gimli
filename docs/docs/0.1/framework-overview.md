---
layout: documentation
title: Framework Overview
---
# Framework Overview

- [Introduction](#introduction)
- [Application Initialisation](#application-initialisation)
- [Request Lifecycle](#request-lifecycle)

<a class="anchor" id="introduction"></a>
## Introduction
It is a good idea to get familiar with whats happening behind the public interface of a framework before you dive
into using it. This section of the documentation aims to provide a high level overview of how Gimli was put together and 
what happens during the lifecycle of a request.

<a class="anchor" id="application-initialisation"></a>
## Application Initialisation
When you run an application created using `gimli new <project>`, Go finds the `main` package in the `main.go` file and calls 
the `main()` function as the entry point to the application. In `main.go` the `bootstrap` package is imported which triggers 
a series of initialisation events as follows:

  * A new `Application` instance is created via a call to `foundation.NewApplication(basePath string)` which causes 
the `foundation` component to do the following:
    - Instantiate a new `di.Container` instance (accessible via `Application.Container()`).
    - Set all project paths in the container, based on the passed in basePath.
    - Register the default container bindings and set the global container instance in the `di.container` package.
    - Register the default service providers (`foundation.providers.routing` and `foundation.providers.config`).
 * Next, Go calls the `init()` functions in `routes.go` and `providers.go` causing your application routes to be registered 
 with the router and your service providers to be registered in the container.
 
 **Note:** *We can guarantee that the application and container will have been initialised by the time the `init()` functions 
 are called due to Go evaluating variable assignment before triggering `init()`.*
 
Once the `bootstrap` package has been imported, the `main()` function is called which in turn calls `Application.Run()`. This 
starts a HTTP server running via a call to `http.ListenAndServe`, passing in the router created by the routing service 
provider when the application was initialised.

<a class="anchor" id="request-lifecycle"></a>
## Request Lifecycle