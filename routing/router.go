package routing

import (
	"net/http"
	"path"
)

// Router manages dispatching of requests to route handlers.
type Router interface {
	http.Handler

	Dispatch(response http.ResponseWriter, request *http.Request)
	SetNotFoundHandler(handler http.Handler)
}

type router struct {
	notFoundHandler http.Handler
	collection      *RouteCollection
}

// NewRouter will create a new router instance with an empty collection and a default NotFoundHandler.
func NewRouter() Router {
	return &router{
		notFoundHandler: http.NotFoundHandler(),
		collection:      NewRouteCollection(),
	}
}

// NewRouterFromCollection will create a new router instance with a default NotFoundHandler and set the collection.
func NewRouterFromCollection(collection *RouteCollection) Router {
	return &router{
		notFoundHandler: http.NotFoundHandler(),
		collection:      collection,
	}
}

// ServeHttp allows the router to be passed into http.ListenAndServe.
func (r *router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	r.Dispatch(response, request)
}

// Dispatch is the heart of the router.
func (r *router) Dispatch(response http.ResponseWriter, request *http.Request) {
	url := path.Clean(request.URL.Path)

	routeCollection := r.collection.RoutesByPath(url)

	handler := r.notFoundHandler

	if len(routeCollection.Routes) > 0 {
		for _, route := range routeCollection.Routes {
			if route.Matches(request) {
				handler = route.handler
				break
			}
		}
	}

	handler.ServeHTTP(response, request)
}

// SetNotFoundHandler sets the handler to be called when no routes are matched. This is http.NotFoundHandler
// by default.
func (r *router) SetNotFoundHandler(handler http.Handler) {
	r.notFoundHandler = handler
}
