package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func TestNewRouter(t *testing.T) {
	assert.Implements(t, (*Router)(nil), NewRouter())
}

func TestRouterCanBePassedAsHttpHandler(t *testing.T) {
	assert.Implements(t, (*http.Handler)(nil), NewRouter())
}

func TestDispatchCallsNotFoundHandlerWhenNoRouteFound(t *testing.T) {
	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest("", "/", nil)

	router.Dispatch(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Contains(t, response.Body.String(), "404 page not found")
}

func TestCustomNotFoundHandlerCanBeSet(t *testing.T) {
	router := NewRouter()
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

func TestStatusOKIsReturnedByDefault(t *testing.T) {
	routeCollection := NewRouteCollection()
	routeCollection.Add(NewRoute("/test", nil, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Handler Was Called!"))
	})))

	router := NewRouterFromCollection(routeCollection)

	response := httptest.NewRecorder()
	request := httptest.NewRequest("", "/test", nil)

	router.Dispatch(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestRouteHandlerIsCalledWhenRouteIsMatched(t *testing.T) {
	routeCollection := NewRouteCollection()
	routeCollection.Add(NewRoute("/test", nil, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Handler Was Called!"))
	})))

	router := NewRouterFromCollection(routeCollection)

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

	router := NewRouterFromCollection(routeCollection)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/test", nil)

	router.Dispatch(response, request)
	assert.Contains(t, response.Body.String(), "The Handler Was Called For POST!")
}

func TestIfMultipleRoutesMatchedTheFirstFoundIsReturned(t *testing.T) {
	routeCollection := NewRouteCollection()
	routeCollection.Add(NewRoute("/test", []string{http.MethodGet}, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Handler For Route1 Was Called!"))
	})))
	routeCollection.Add(NewRoute("/test", []string{http.MethodGet}, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Handler For Route2 Was Called!"))
	})))

	router := NewRouterFromCollection(routeCollection)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/test", nil)

	router.Dispatch(response, request)
	assert.Contains(t, response.Body.String(), "The Handler For Route1 Was Called!")
}

func TestRequestMethodHelperFunctions(t *testing.T) {
	router := NewRouter()

	router.Get("", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The GET handler was called."))
	}))

	router.Post("", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The POST handler was called."))
	}))

	router.Put("", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The PUT handler was called."))
	}))

	router.Patch("", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The PATCH handler was called."))
	}))

	router.Delete("", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The DELETE handler was called."))
	}))

	router.Any("/any", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Any handler was called."))
	}))

	router.Match("/match", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The Match handler was called."))
	}), http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete)

	assert.Contains(t, runRequest(http.MethodGet, "/", router), "The GET handler was called.")
	assert.Contains(t, runRequest(http.MethodPost, "/", router), "The POST handler was called.")
	assert.Contains(t, runRequest(http.MethodPut, "/", router), "The PUT handler was called.")
	assert.Contains(t, runRequest(http.MethodPatch, "/", router), "The PATCH handler was called.")
	assert.Contains(t, runRequest(http.MethodDelete, "/", router), "The DELETE handler was called.")

	assert.Contains(t, runRequest(http.MethodGet, "/any", router), "The Any handler was called.")
	assert.Contains(t, runRequest(http.MethodPost, "/any", router), "The Any handler was called.")
	assert.Contains(t, runRequest(http.MethodPut, "/any", router), "The Any handler was called.")
	assert.Contains(t, runRequest(http.MethodPatch, "/any", router), "The Any handler was called.")
	assert.Contains(t, runRequest(http.MethodDelete, "/any", router), "The Any handler was called.")

	assert.Contains(t, runRequest(http.MethodGet, "/match", router), "The Match handler was called.")
	assert.Contains(t, runRequest(http.MethodPost, "/match", router), "The Match handler was called.")
	assert.Contains(t, runRequest(http.MethodPut, "/match", router), "The Match handler was called.")
	assert.Contains(t, runRequest(http.MethodPatch, "/match", router), "The Match handler was called.")
	assert.Contains(t, runRequest(http.MethodDelete, "/match", router), "The Match handler was called.")
}

func TestRouterCanGroupRoutesByPrefix(t *testing.T) {
	router := NewRouter()

	router.Get("test", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("The not grouped test handler was called."))
	}))

	router.Group("prefix", func(router Router) {

		router.Get("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("The prefix base handler was called."))
		}))

		router.Get("test", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("The prefix test handler was called."))
		}))

	})

	assert.Contains(t, runRequest(http.MethodGet, "/test", router), "The not grouped test handler was called.")
	assert.Contains(t, runRequest(http.MethodGet, "/prefix", router), "The prefix base handler was called.")
	assert.Contains(t, runRequest(http.MethodGet, "/prefix/test", router), "The prefix test handler was called.")
}

func runRequest(method, path string, router Router) string {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(method, path, nil)

	router.Dispatch(response, request)

	return response.Body.String()
}
