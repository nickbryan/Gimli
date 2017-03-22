package routing

import (
	"net/http"
	"strings"
)

type Route struct {
	matchers []Matcher
	path     string
	methods  []string
	name     string
	handler  http.Handler
}

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

func (r *Route) Path() string {
	return r.path
}

func (r *Route) SetPath(path string) {
	r.path = "/" + strings.TrimLeft(strings.TrimSpace(path), "/")
}

func (r *Route) Methods() []string {
	return r.methods
}

func (r *Route) SetMethods(methods ...string) {
	formatted := []string{}
	for _, method := range methods {
		formatted = append(formatted, strings.ToUpper(method))
	}

	r.methods = formatted
}

func (r *Route) Name() string {
	return r.name
}

func (r *Route) SetName(name string) {
	r.name = strings.ToLower(strings.TrimSpace(name))
}

func (r *Route) SetHandler(handler http.Handler) {
	r.handler = handler
}

func (r *Route) AddMatcher(matcher Matcher) {
	r.matchers = append(r.matchers, matcher)
}

func (r *Route) Matches(request *http.Request) bool {
	for _, matcher := range r.matchers {
		if matcher.Match(r, request) == false {
			return false
		}
	}

	return true
}
