package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMap = map[string]interface{}{
	"string_value": "Hello World",
	"float_value":  4.20,
	"bool_value":   true,
	"int_value":    420,
	"nil_value":    nil,
	"slice_value":  [2]string{"hello", "world"},
}

func TestNewRepository(t *testing.T) {
	r := NewRepository()

	assert.IsType(t, new(Repository), r)
}

func TestNewPopulatedRepository(t *testing.T) {
	r := NewPopulatedRepository(testMap)

	assert.IsType(t, new(Repository), r)
	assert.Equal(t, &Repository{items: testMap}, r)
}

var repo = NewPopulatedRepository(testMap)

func TestHas(t *testing.T) {
	assert.True(t, repo.Has("string_value"))

	assert.False(t, repo.Has("unset_key"))
}

func TestGet(t *testing.T) {
	for key, val := range testMap {
		assert.Equal(t, val, repo.Get(key))
	}

	assert.Nil(t, repo.Get("unset_key"))
}

func TestGetDefault(t *testing.T) {
	assert.Equal(t, testMap["string_value"], repo.GetDefault("string_value", "default string"))
	assert.Equal(t, "default string", repo.GetDefault("unset_key", "default string"))
}

func TestSet(t *testing.T) {
	key, value, updatedValue := "new_key", "new_value", "updated_value"

	r := repo.Set(key, value)
	assert.IsType(t, new(Repository), r)
	assert.Equal(t, value, r.Get(key))

	repo.Set(key, updatedValue)
	assert.Equal(t, updatedValue, r.Get(key))
}

func TestSetMultiple(t *testing.T) {
	newMap := map[string]interface{}{
		"new_string_key": "new_string_val",
		"new_int_key":    42,
	}

	repo.SetMultiple(newMap)

	for key, val := range newMap {
		assert.Equal(t, val, repo.Get(key))
	}
}

func TestAll(t *testing.T) {
	r := NewPopulatedRepository(testMap)

	assert.True(t, assert.ObjectsAreEqual(r.All(), testMap))
}
