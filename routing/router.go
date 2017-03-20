package routing

/*
TODO: Maybe abstract interfaces for RouteAdder, RoutSearcher and RouteAdderSearcher. This can be used in place of trie
for extension?
*/

type Router interface{}

type router struct{}

func NewRouter() Router {
	return new(router)
}
