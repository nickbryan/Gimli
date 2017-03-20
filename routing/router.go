package routing

import "net/http"

/*
TODO: Maybe abstract interfaces for RouteAdder, RoutSearcher and RouteAdderSearcher. This can be used in place of trie
for extension?
*/

type Router interface {
	Dispatch(response http.ResponseWriter, request *http.Request)
}

type router struct{}

func NewRouter() Router {
	return new(router)
}

func (r *router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	r.Dispatch(response, request)
}

func (r *router) Dispatch(response http.ResponseWriter, request *http.Request) {

}
