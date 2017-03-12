package config

// Repository can be used as a simple container for storing key => value
// configuration pairs.
type Repository struct {
	items map[string]interface{}
}

// NewRepository will create a new empty configuration repository.
func NewRepository() *Repository {
	return &Repository{make(map[string]interface{})}
}

// NewPopulatedRepository will create a new configuration repository
// containing the supplied key => value pairs.
func NewPopulatedRepository(items map[string]interface{}) *Repository {
	return &Repository{items}
}

// Has can be used to check if a key exists inside the repository.
func (r *Repository) Has(key string) bool {
	_, ok := r.items[key]

	return ok
}

// Get can be used to receive a value for the given key from the repository
// and returns nil if not found.
func (r *Repository) Get(key string) interface{} {
	return r.items[key]
}

// GetDefault can be used to receive a value for the given key from the repository
// and returns the supplied default value if not found.
func (r *Repository) GetDefault(key string, defaultVal interface{}) interface{} {
	if !r.Has(key) {
		return defaultVal
	}

	return r.items[key]
}

// Set adds a key => value pair to the repository.
func (r *Repository) Set(key string, value interface{}) *Repository {
	r.items[key] = value

	return r
}

// SetMultiple can be used to add multiple key => value pairs to the repository.
func (r *Repository) SetMultiple(values map[string]interface{}) *Repository {
	for key, value := range values {
		r.items[key] = value
	}

	return r
}

// All returns all items bound in the repository.
func (r *Repository) All() map[string]interface{} {
	return r.items
}
