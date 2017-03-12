package routing

import (
	"net/http"
	"path"
	"strings"
)

type Router interface {
	Map(method RequestMethod, pattern string, handler interface{})
	Get(pattern string, handler RouteClosure)
	Post(pattern string, handler RouteClosure)
	Put(pattern string, handler RouteClosure)
	Delete(pattern string, handler RouteClosure)
	Parse(request *http.Request, response http.ResponseWriter)
}

// Router handles all routing
type router struct {
	routes *Tree
}

// NewRouter creates a new router and initialises it
func NewRouter() Router {
	tree := Tree{segment: "/", isNamedParam: false, methods: make(map[RequestMethod]*Route)}

	return &router{routes: &tree}
}

// Map a handler to the given method and uri pattern
func (router *router) Map(method RequestMethod, pattern string, handler interface{}) {
	if pattern[0] != '/' {
		panic("Path must start with a '/'.")
	}

	route := &Route{method, pattern, handler}

	router.routes.addNode(method, pattern, route)
}

// Get is a convenient wrapper around map for get requests
func (router *router) Get(pattern string, handler RouteClosure) {
	router.Map(GET, pattern, handler)
}

// Post is a convenient wrapper around map for post requests
func (router *router) Post(pattern string, handler RouteClosure) {
	router.Map(POST, pattern, handler)
}

// Put is a convenient wrapper around map for put requests
func (router *router) Put(pattern string, handler RouteClosure) {
	router.Map(PUT, pattern, handler)
}

// Delete is a convenient wrapper around map for delete requests
func (router *router) Delete(pattern string, handler RouteClosure) {
	router.Map(DELETE, pattern, handler)
}

// Parse is used to parse the request and response
func (router *router) Parse(request *http.Request, response http.ResponseWriter) {
	url := path.Clean(request.URL.Path)
	// TODO: handle as http.Error(rw, "Method Not Allowed", 405)
	method := methodStringToRequestMethod(request.Method)

	// Handle static files
	// TODO: add css, js, img etc
	if url == "/favicon.ico" || url == "/robots.txt" {
		//file := config.PublicPath + url
		//http.ServeFile(response, request, file)
		return
	}

	request.ParseForm()
	params := request.Form
	tree, _ := router.routes.traverse(strings.Split(url, "/")[1:], params)

	// Check if the route is found
	if tree.methods[method] != nil {
		switch tree.methods[method].Handler.(type) {
		case RouteClosure:
			tree.methods[method].Handler.(RouteClosure)(request, response, params)
			return
		case ControllerInterface:
		// do stuff
		default:
			// TODO: add type to this string
			panic("Route Handler type not supported")
		}
	}

	response.WriteHeader(http.StatusNotFound)
}

func methodStringToRequestMethod(method string) RequestMethod {
	switch method {
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	default:
		panic("Request Method not supported: " + method)
	}
}
