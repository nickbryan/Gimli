package routing

import "net/http"

type Matcher interface {
	Match(route *Route, request *http.Request) bool
}

type MatcherFunc func(route *Route, request *http.Request) bool

func (matcher MatcherFunc) Match(route *Route, request *http.Request) bool {
	return matcher(route, request)
}

func MethodMatcher(route *Route, request *http.Request) bool {
	for _, method := range route.Methods() {
		if method == request.Method {
			return true
		}
	}

	return false
}
