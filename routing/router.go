package routing

import (
	"net/http"
	"path"
)

/*
TODO: Maybe abstract interfaces for RouteAdder, RoutSearcher and RouteAdderSearcher. This can be used in place of trie
for extension?
*/

type Router interface {
	http.Handler

	Dispatch(response http.ResponseWriter, request *http.Request)
	SetNotFoundHandler(handler http.Handler)
}

type router struct {
	notFoundHandler http.Handler
	collection      *RouteCollection
}

func NewRouter(collection *RouteCollection) Router {
	return &router{
		notFoundHandler: http.NotFoundHandler(),
		collection:      collection,
	}
}

func (r *router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	r.Dispatch(response, request)
}

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

func (r *router) SetNotFoundHandler(handler http.Handler) {
	r.notFoundHandler = handler
}
