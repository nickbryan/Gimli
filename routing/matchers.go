package routing

import "net/http"

// Matcher is used to check if the given route matches a request based on the conditions set out
// in the Match function.
type Matcher interface {
	// Match should run conditionals to check if the route and request match based on requirements.
	Match(route *Route, request *http.Request) bool
}

// MatcherFunc is an adapter to allow the use of ordinary functions as a Matcher.
type MatcherFunc func(route *Route, request *http.Request) bool

// Match calls matcher(route, request)
func (matcher MatcherFunc) Match(route *Route, request *http.Request) bool {
	return matcher(route, request)
}

// MethodMatcher will check to see if the request method matches on of the routes allowed methods.
func MethodMatcher(route *Route, request *http.Request) bool {
	for _, method := range route.Methods() {
		if method == request.Method {
			return true
		}
	}

	return false
}
