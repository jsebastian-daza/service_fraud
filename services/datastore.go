package services

import (
	"fmt"
	"service_fraud/utils"
	"time"
)

// RequestDataStore is a generic data structure that stores key-value pairs
// with an expiration time for each entry.
type RequestDataStore[K comparable, V any] struct {
	data   map[K]V
	expiry map[K]time.Time
}

// timeLimit defines the global expiration time for stored items.
const timeLimit = utils.TTL_IN_MINUTES * time.Minute

// NewRequestDataStore creates and initializes a new RequestDataStore.
func NewRequestDataStore[K comparable, V any]() *RequestDataStore[K, V] {
	return &RequestDataStore[K, V]{
		data:   make(map[K]V),
		expiry: make(map[K]time.Time),
	}
}

// Set stores a value with a specified key and sets its expiration time.
func (store *RequestDataStore[K, V]) Set(key K, value V) error {
	store.data[key] = value
	store.expiry[key] = time.Now().Add(timeLimit)
	return nil
}

// Get retrieves a value associated with the specified key.
// It returns an error if the key has expired or does not exist.
func (store *RequestDataStore[K, V]) Get(key K) (V, error) {
	if time.Now().After(store.expiry[key]) {
		delete(store.data, key)
		delete(store.expiry, key)
		var zeroValue V // Valor cero para el tipo V
		return zeroValue, fmt.Errorf("el dato ha expirado")
	}

	value, exists := store.data[key]
	if !exists {
		var zeroValue V // Valor cero para el tipo V
		return zeroValue, fmt.Errorf("clave no encontrada")
	}
	return value, nil
}

// Expire manually removes a key and its associated value from the store.
func (store *RequestDataStore[K, V]) Expire(key K) error {
	delete(store.data, key)
	delete(store.expiry, key)
	return nil
}
