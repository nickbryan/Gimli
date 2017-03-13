package di

import (
	"errors"
	"sync"
)

// Container is a simple (thread safe) dependency injection container.
type Container interface {
	Has(id string) bool
	Instance(id string, instance interface{})
	Resolve(id string) (interface{}, error)
	MustResolve(id string) interface{}
	Bind(id string, concrete Resolver)
	IsShared(id string) bool
	Factory(id string, concrete Resolver)
	Register(provider ServiceProvider)
}

// ServiceProvider acts as a way of encapsulating multiple or complex service binding logic
// and gives a convenient way to register services with the container.
type ServiceProvider interface {
	Register(container Container)
}

type binding struct {
	concrete Resolver
	shared   bool
}

type container struct {
	mux       sync.RWMutex
	bindings  map[string]*binding
	instances map[string]interface{}
}

// NewContainer will return an empty container.
func NewContainer() Container {
	return &container{
		bindings:  make(map[string]*binding),
		instances: make(map[string]interface{}),
	}
}

var instance Container

// SetInstance is a convenience function for setting a global instance
// of the container. You do not have to use the container this way.
func SetInstance(container Container) Container {
	instance = container

	return instance
}

// ForgetInstance allows resetting of the global container instance.
func ForgetInstance() {
	instance = nil
}

// GetInstance will return a global instance of the container if one has been set.
// If the instance has not been set one will be created and returned.
func GetInstance() Container {
	if instance == nil {
		instance = NewContainer()
	}

	return instance
}

// Has will return true if a service with the given id is set in the container or false if not.
func (c *container) Has(id string) bool {
	c.mux.RLock()
	defer c.mux.RUnlock()

	if _, ok := c.instances[id]; ok {
		return true
	}

	if _, ok := c.bindings[id]; ok {
		return true
	}

	return false
}

// Resolver should be passed into Bind and Factory. All service building logic
// should be encapsulated within the closure.
type Resolver func(container Container) interface{}

// Instance provides a way of adding an already built property in the container.
func (c *container) Instance(id string, instance interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.instances[id] = instance
}

// Bind can be used to create a shared service within the container.
func (c *container) Bind(id string, concrete Resolver) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.bindings[id] = &binding{
		concrete: concrete,
		shared:   true,
	}
}

// Factory can be used to bind a service to the container that will be built on each call.
func (c *container) Factory(id string, concrete Resolver) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.bindings[id] = &binding{
		concrete: concrete,
		shared:   false,
	}
}

// Resolve will return the service with the given id from the container if bound or an error on failure.
func (c *container) Resolve(id string) (interface{}, error) {
	c.mux.RLock()

	if instance, ok := c.instances[id]; ok {
		c.mux.RUnlock()
		return instance, nil
	}

	if _, ok := c.bindings[id]; ok == false {
		c.mux.RUnlock()
		return nil, errors.New("Abstract " + id + " does not exist in container.")
	}

	bound := c.bindings[id]
	c.mux.RUnlock()

	instance := bound.concrete(c)

	if c.IsShared(id) {
		c.mux.Lock()
		c.instances[id] = instance
		c.mux.Unlock()
	}

	return instance, nil
}

// MustResolve will panic if the service with the given id could not be resolved
// from the container.
func (c *container) MustResolve(id string) interface{} {
	instance, err := c.Resolve(id)

	if err != nil {
		panic(err)
	}

	return instance
}

// IsShared can be used to identify if a service was bound as a singleton
// or a factory (Bind or Factory).
func (c *container) IsShared(id string) bool {
	c.mux.RLock()
	defer c.mux.RUnlock()

	if _, ok := c.instances[id]; ok {
		return true
	}

	if binding, ok := c.bindings[id]; ok {
		return binding.shared
	}

	return false
}

// Register allows a service provider to be bound in the container.
func (c *container) Register(provider ServiceProvider) {
	provider.Register(c)
}
