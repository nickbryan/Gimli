package routing

import (
	"net/http"
	"strings"
)

type Route struct {
	path    string
	methods []string
	name    string
}

func NewRoute(path string, methods []string) *Route {
	r := &Route{}

	r.SetPath(path)

	if methods == nil {
		methods = []string{http.MethodGet}
	}

	r.SetMethods(methods...)

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
