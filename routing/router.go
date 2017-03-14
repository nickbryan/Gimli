package routing

type Router interface{}

type router struct{}

func NewRouter() Router {
	return new(router)
}
