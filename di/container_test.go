package di

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	assert.Implements(t, (*Container)(nil), NewContainer())
}

func TestSetInstance(t *testing.T) {
	// Ensure the instance is nil to start
	instance = nil
	c := SetInstance(NewContainer())
	assert.Equal(t, c, instance)
}

func TestForgetInstance(t *testing.T) {
	SetInstance(NewContainer())
	assert.NotNil(t, instance)

	ForgetInstance()
	assert.Nil(t, instance)
}

func TestGetInstance(t *testing.T) {
	// Ensure the instance is nil to start
	ForgetInstance()
	c := SetInstance(NewContainer())
	assert.Equal(t, c, GetInstance())
}

func TestGetInstanceCreatesContainerIfNotSet(t *testing.T) {
	// Ensure the instance is nil to start
	ForgetInstance()
	assert.NotNil(t, GetInstance())
}

func TestHasReturnsFalseIfNothingInContainer(t *testing.T) {
	c := NewContainer()
	assert.False(t, c.Has("Something"))
}

func TestInstanceAddsValueToContainer(t *testing.T) {
	c := NewContainer()

	c.Instance("TheMeaningOfLife", 42)
	assert.True(t, c.Has("TheMeaningOfLife"))
}

func TestResolveReturnsInstanceValueAndNilWhenBound(t *testing.T) {
	c := NewContainer()

	c.Instance("TheMeaningOfLife", 42)

	val, err := c.Resolve("TheMeaningOfLife")
	assert.Equal(t, 42, val)
	assert.Nil(t, err)
}

func TestResolveReturnsNilAndErrorWhenNotBound(t *testing.T) {
	c := NewContainer()

	val, err := c.Resolve("TheMeaningOfLife")
	assert.Nil(t, val)
	assert.EqualError(t, err, "Abstract TheMeaningOfLife does not exist in container.")
}

func TestMustResolveReturnsValueWhenBound(t *testing.T) {
	c := NewContainer()

	c.Instance("TheMeaningOfLife", 42)

	assert.Equal(t, 42, c.MustResolve("TheMeaningOfLife"))
}

func TestMustResolvePanicsWhenNotBound(t *testing.T) {
	c := NewContainer()

	assert.Panics(t, func() {
		c.MustResolve("TheMeaningOfLife")
	})
}

func TestBindAddsResolverToContainer(t *testing.T) {
	c := NewContainer()

	resolver := Resolver(func(container Container) interface{} {
		return 42
	})

	c.Bind("TheMeaningOfLife", resolver)
	assert.True(t, c.Has("TheMeaningOfLife"))
}

func TestResolveReturnsBoundConcreteValue(t *testing.T) {
	c := NewContainer()
	c.Bind("TheMeaningOfLife", func(container Container) interface{} {
		return 42
	})
	assert.Equal(t, 42, c.MustResolve("TheMeaningOfLife"))
}

type TestBindObjectStub struct {
	Value int64
}

func TestBindCreatesASharedInstance(t *testing.T) {
	c := NewContainer()
	c.Bind("ObjectStub", func(container Container) interface{} {
		return &TestBindObjectStub{time.Now().UnixNano()}
	})
	assert.Exactly(t, c.MustResolve("ObjectStub"), c.MustResolve("ObjectStub"))

}

func TestFactoryCreatesANewInstance(t *testing.T) {
	c := NewContainer()
	c.Factory("ObjectStub", func(container Container) interface{} {
		return &TestBindObjectStub{time.Now().UnixNano()}
	})
	assert.NotEqual(t, c.MustResolve("ObjectStub"), c.MustResolve("ObjectStub"))
}

func TestIsShared(t *testing.T) {
	c := NewContainer()
	c.Bind("BindIsShared", func(container Container) interface{} {
		return "Hi"
	})
	assert.True(t, c.IsShared("BindIsShared"))

	c.Instance("InstanceIsShared", "Hello")
	assert.True(t, c.IsShared("BindIsShared"))

	c.Factory("FactoryIsNotShared", func(container Container) interface{} {
		return "Hi"
	})
	assert.False(t, c.IsShared("FactoryIsNotShared"))
}

func TestContainerIsPassedToResolver(t *testing.T) {
	c := NewContainer()
	c.Bind("Container", func(container Container) interface{} {
		return container
	})

	assert.Exactly(t, c, c.MustResolve("Container"))
}

func TestBindingCanBeOverridden(t *testing.T) {
	c := NewContainer()

	c.Instance("Instance", "Hello")
	c.Instance("Instance", "Hi")
	assert.Equal(t, "Hi", c.MustResolve("Instance"))

	c.Bind("Bind", func(container Container) interface{} {
		return "Hello"
	})
	c.Bind("Bind", func(container Container) interface{} {
		return "Hi"
	})
	assert.Equal(t, "Hi", c.MustResolve("Bind"))

	c.Factory("Factory", func(container Container) interface{} {
		return "Hello"
	})
	c.Factory("Factory", func(container Container) interface{} {
		return "Hi"
	})
	assert.Equal(t, "Hi", c.MustResolve("Factory"))
}

type TestNestedObjectStub struct {
	NestedDependency TestBindObjectStub
}

func TestBindNestedDependency(t *testing.T) {
	c := NewContainer()

	c.Bind("ObjectStub", func(container Container) interface{} {
		return TestBindObjectStub{42}
	})
	c.Bind("NestedObjectStub", func(container Container) interface{} {
		return TestNestedObjectStub{container.MustResolve("ObjectStub").(TestBindObjectStub)}
	})

	resolved := c.MustResolve("NestedObjectStub").(TestNestedObjectStub)
	assert.Equal(t, 42, int(resolved.NestedDependency.Value))
}

func TestFactoryNestedDependency(t *testing.T) {
	c := NewContainer()

	c.Factory("ObjectStub", func(container Container) interface{} {
		return TestBindObjectStub{42}
	})
	c.Factory("NestedObjectStub", func(container Container) interface{} {
		return TestNestedObjectStub{container.MustResolve("ObjectStub").(TestBindObjectStub)}
	})

	resolved := c.MustResolve("NestedObjectStub").(TestNestedObjectStub)
	assert.Equal(t, 42, int(resolved.NestedDependency.Value))
}

type TestingServiceProvider struct{}

func (provider *TestingServiceProvider) Register(container Container) {
	container.Instance("ProvidedInstance", 42)
}

func TestRegisterAddsServiceToContainer(t *testing.T) {
	c := NewContainer()
	c.Register(&TestingServiceProvider{})
	assert.Equal(t, 42, c.MustResolve("ProvidedInstance"))
}
