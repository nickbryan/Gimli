package routing

import (
	"net/http"
	pkgPath "path"
	"strings"
)

// Route should contain all information needed to match a url path to a handler bound on the route.
type Route struct {
	matchers []Matcher
	path     string
	methods  []string
	name     string
	handler  http.Handler
}

// NewRoute will create a new Route and set a default MethodMatcher.
func NewRoute(path string, methods []string, handler http.Handler) *Route {
	r := &Route{}

	r.matchers = []Matcher{MatcherFunc(MethodMatcher)}

	r.SetPath(path)

	if methods == nil {
		methods = []string{http.MethodGet}
	}

	r.SetMethods(methods...)

	r.SetHandler(handler)

	return r
}

// Path will return the required pattern used to match the route.
func (r *Route) Path() string {
	return r.path
}

// SetPath will normalise and set the pattern used to match the route.
func (r *Route) SetPath(path string) {
	r.path = pkgPath.Clean("/" + strings.TrimLeft(strings.TrimSpace(path), "/"))
}

// Methods will return a list of request methods that this route will respond to.
func (r *Route) Methods() []string {
	return r.methods
}

// SetMethods will normalise and set the request methods for this route based on http.Method*.
func (r *Route) SetMethods(methods ...string) {
	formatted := []string{}
	for _, method := range methods {
		formatted = append(formatted, strings.ToUpper(method))
	}

	r.methods = formatted
}

// Name can be used to lookup a route by its name.
func (r *Route) Name() string {
	return r.name
}

// SetName will normalise and set the name of the route.
func (r *Route) SetName(name string) {
	r.name = strings.ToLower(strings.TrimSpace(name))
}

// SetHandler will set the handler that will be called if the route is matched.
func (r *Route) SetHandler(handler http.Handler) {
	r.handler = handler
}

// AddMatcher will add a Matcher to the list. These are used to check if the route matches a specific request criteria.
func (r *Route) AddMatcher(matcher Matcher) {
	r.matchers = append(r.matchers, matcher)
}

// Matches will run all matchers to check if this route matches the passed in request.
func (r *Route) Matches(request *http.Request) bool {
	for _, matcher := range r.matchers {
		if matcher.Match(r, request) == false {
			return false
		}
	}

	return true
}
