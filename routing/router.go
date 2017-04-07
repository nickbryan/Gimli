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

	Get(path string, handler http.Handler)
	Post(path string, handler http.Handler)
	Put(path string, handler http.Handler)
	Patch(path string, handler http.Handler)
	Delete(path string, handler http.Handler)
	Any(path string, handler http.Handler)
	Match(path string, handler http.Handler, methods ...string)
	Group(prefix string, group func(router Router))
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

	handler := r.notFoundHandler

	if routeCollection := r.collection.RoutesByPath(url); routeCollection != nil {
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

// Get is a helper that adds a route to the collection that will match the request the method GET.
func (r *router) Get(path string, handler http.Handler) {
	r.collection.Add(NewRoute(path, []string{http.MethodGet}, handler))
}

// Post is a helper that adds a route to the collection that will match the request the method POST.
func (r *router) Post(path string, handler http.Handler) {
	r.collection.Add(NewRoute(path, []string{http.MethodPost}, handler))
}

// Put is a helper that adds a route to the collection that will match the request the method PUT.
func (r *router) Put(path string, handler http.Handler) {
	r.collection.Add(NewRoute(path, []string{http.MethodPut}, handler))
}

// Patch is a helper that adds a route to the collection that will match the request the method PATCH.
func (r *router) Patch(path string, handler http.Handler) {
	r.collection.Add(NewRoute(path, []string{http.MethodPatch}, handler))
}

// Delete is a helper that adds a route to the collection that will match the request method DELETE.
func (r *router) Delete(path string, handler http.Handler) {
	r.collection.Add(NewRoute(path, []string{http.MethodDelete}, handler))
}

// Any is a helper that adds a route to the collection that will match any request method.
func (r *router) Any(path string, handler http.Handler) {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete}

	r.collection.Add(NewRoute(path, methods, handler))
}

// Match is a helper that adds a route to the collection that will match the given request methods.
func (r *router) Match(path string, handler http.Handler, methods ...string) {
	r.collection.Add(NewRoute(path, methods, handler))
}

func (r *router) Group(prefix string, group func(router Router)) {
	routeCollection := NewRouteCollection()

	group(NewRouterFromCollection(routeCollection))

	routeCollection.Prefix(prefix)

	r.collection.AddCollection(routeCollection)
}
