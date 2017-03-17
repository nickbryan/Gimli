package routing

type RouteCollection struct {
	routes      routeTrie
	allRoutes   []*Route
	namedRoutes map[string]*Route
}

func NewRouteCollection() *RouteCollection {
	return &RouteCollection{
		routes:      newRouteTrie(),
		allRoutes:   []*Route{},
		namedRoutes: make(map[string]*Route),
	}
}

func (collection *RouteCollection) Add(route *Route) {
	if route.Name() != "" {
		collection.namedRoutes[route.Name()] = route
	}

	collection.routes.add(route)
	collection.allRoutes = append(collection.allRoutes, route)
}

func (collection *RouteCollection) RefreshNamedRoutes() {
	collection.namedRoutes = map[string]*Route{}

	for _, route := range collection.allRoutes {
		if route.Name() != "" {
			collection.namedRoutes[route.Name()] = route
		}
	}
}

func (collection *RouteCollection) Count() int {
	return len(collection.allRoutes)
}
