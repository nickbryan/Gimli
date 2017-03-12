package routing

import (
	"net/http"
	"net/url"
)

type RouteClosure func(request *http.Request, response http.ResponseWriter, params url.Values)

// RequestMethod type is used to identify the supported request methods
type RequestMethod string

// Constants representing supported request methods
const (
	GET    RequestMethod = "GET"
	POST   RequestMethod = "POST"
	PUT    RequestMethod = "PUT"
	DELETE RequestMethod = "DELETE"
)

// Route represents a single application route
type Route struct {
	Method  RequestMethod
	Pattern string
	Handler interface{}
}
