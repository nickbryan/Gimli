package routing

type RouteCollection struct {
	routes routeTrie
}

func NewRouteCollection() *RouteCollection {
	return &RouteCollection{
		routes: newRouteTrie(),
	}
}

func (collection *RouteCollection) Add(route *Route) {
	collection.routes.add(route)
}
