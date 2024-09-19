package interfaces

// DataStore defines a generic interface for storing and retrieving key-value pairs.
// The key type K must be comparable, and the value type V can be any type.
type DataStore[K comparable, V any] interface {
	// Set stores the value associated with the given key.
	// Returns an error if the operation fails.
	Set(key K, value V) error
	// Get retrieves the value associated with the given key.
	// Returns the value and an error if the key is not found or another error occurs.
	Get(key K) (V, error)
	// Expire removes the value associated with the given key.
	// Returns an error if the operation fails.
	Expire(key K) error
}
