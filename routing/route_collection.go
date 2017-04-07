package routing

// RouteCollection provides helpful ways of dealing with collections of Routes.
type RouteCollection struct {
	routes      routeTrie
	allRoutes   []*Route
	namedRoutes map[string]*Route
}

// RouteMatchGroup is used to encapsulate found routes, when looked up by path, with the parsed named params
// from the url.
type RouteMatchGroup struct {
	Routes    []*Route
	UrlParams map[string]string
}

// NewRouteCollection creates an empty RouteCollection.
func NewRouteCollection() *RouteCollection {
	return &RouteCollection{
		routes:      newRouteTrie(),
		allRoutes:   []*Route{},
		namedRoutes: make(map[string]*Route),
	}
}

// Add a route to the collection. If the route has a name assigned it will be added to the
// list of named routes.
func (collection *RouteCollection) Add(route *Route) {
	if route.Name() != "" {
		collection.namedRoutes[route.Name()] = route
	}

	collection.routes.add(route)
	collection.allRoutes = append(collection.allRoutes, route)
}

// RouteByName can be used to lookup a route by its name.
func (collection *RouteCollection) RouteByName(name string) *Route {
	return collection.namedRoutes[name]
}

// RoutesByPath will lookup routes in the route trie and return a RoutMatchGroup
// containing any matched routes and the parsed params.
func (collection *RouteCollection) RoutesByPath(path string) *RouteMatchGroup {
	routes, params := collection.routes.search(path)

	if len(routes) == 0 {
		return nil
	}

	return &RouteMatchGroup{routes, params}
}

// RefreshNamedRoutes will clear the named routes list and add all named routes back from the all routes list.
// This is useful if a name is assigned to a route after it has been added to the collection.
func (collection *RouteCollection) RefreshNamedRoutes() {
	collection.namedRoutes = map[string]*Route{}

	for _, route := range collection.allRoutes {
		if route.Name() != "" {
			collection.namedRoutes[route.Name()] = route
		}
	}
}

// Count will return the total number of routes in the collection.
func (collection *RouteCollection) Count() int {
	return len(collection.allRoutes)
}

func (collection *RouteCollection) Prefix(prefix string) {
	if prefix == "" {
		return
	}

	collection.routes = newRouteTrie()

	for _, route := range collection.allRoutes {
		route.SetPath(prefix + route.Path())

		collection.Add(route)
	}
}

func (collection *RouteCollection) AllRoutes() []*Route {
	return collection.allRoutes
}

func (collection *RouteCollection) AddCollection(routeCollection *RouteCollection) {
	for _, route := range routeCollection.AllRoutes() {
		collection.Add(route)
	}
}
