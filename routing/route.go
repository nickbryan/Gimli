package routing

import "strings"

type Route struct {
	path    string
	methods []string
}

func NewRoute(path string, methods []string) *Route {
	r := &Route{}

	r.SetPath(path)
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
