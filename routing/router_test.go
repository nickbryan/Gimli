package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func TestNewRouter(t *testing.T) {
	assert.Implements(t, (*Router)(nil), NewRouter(NewRouteCollection()))
}

func TestRouterCanBePassedAsHttpHandler(t *testing.T) {
	assert.Implements(t, (*http.Handler)(nil), NewRouter(NewRouteCollection()))
}

func TestDispatchCallsNotFoundHandlerWhenNoRouteFound(t *testing.T) {
	router := NewRouter(NewRouteCollection())
	response := httptest.NewRecorder()
	request := httptest.NewRequest("", "/", nil)

	router.Dispatch(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Contains(t, response.Body.String(), "404 page not found")
}

func TestCustomNotFoundHandlerCanBeSet(t *testing.T) {
	router := NewRouter(NewRouteCollection())
	router.SetNotFoundHandler(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusNotImplemented)
		response.Write([]byte("This is the custom text"))
	}))

	response := httptest.NewRecorder()
	request := httptest.NewRequest("", "/", nil)

	router.Dispatch(response, request)

	assert.Equal(t, http.StatusNotImplemented, response.Code)
	assert.Contains(t, response.Body.String(), "This is the custom text")
}

func TestRouteHandlerIsCalledWhenRouteIsMatched(t *testing.T) {
	routeCollection := NewRouteCollection()
	routeCollection.Add(NewRoute("/test", nil, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Handler Was Called!"))
	})))

	router := NewRouter(routeCollection)

	response := httptest.NewRecorder()
	request := httptest.NewRequest("", "/test", nil)

	router.Dispatch(response, request)
	assert.Contains(t, response.Body.String(), "The Handler Was Called!")
}

func TestMultipleFoundRoutesCanBeFilteredByMethod(t *testing.T) {
	routeCollection := NewRouteCollection()
	routeCollection.Add(NewRoute("/test", []string{http.MethodGet}, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Handler Was Called For GET!"))
	})))
	routeCollection.Add(NewRoute("/test", []string{http.MethodPost}, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Handler Was Called For POST!"))
	})))

	router := NewRouter(routeCollection)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/test", nil)

	router.Dispatch(response, request)
	assert.Contains(t, response.Body.String(), "The Handler Was Called For POST!")
}

// TODO: add has for routes to collection to allow check for multiple of the same route being added.

func TestIfMultipleRoutesMatchedTheFirstFoundIsReturned(t *testing.T) {
	routeCollection := NewRouteCollection()
	routeCollection.Add(NewRoute("/test", []string{http.MethodGet}, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Handler For Route1 Was Called!"))
	})))
	routeCollection.Add(NewRoute("/test", []string{http.MethodGet}, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Handler For Route2 Was Called!"))
	})))

	router := NewRouter(routeCollection)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/test", nil)

	router.Dispatch(response, request)
	assert.Contains(t, response.Body.String(), "The Handler For Route1 Was Called!")
}
