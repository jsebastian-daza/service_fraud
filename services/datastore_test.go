package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestDataStore_SetAndGet(t *testing.T) {
	store := NewRequestDataStore[string, string]()
	key := "testKey"
	value := "testValue"
	err := store.Set(key, value)
	assert.NoError(t, err)

	retrievedValue, err := store.Get(key)

	assert.NoError(t, err)
	assert.Equal(t, value, retrievedValue)
}

func TestRequestDataStore_Get_NotFound(t *testing.T) {
	store := NewRequestDataStore[string, string]()
	key := "nonExistentKey"

	retrievedValue, err := store.Get(key)

	assert.Error(t, err)
	assert.Equal(t, "el dato ha expirado", err.Error())
	assert.Equal(t, "", retrievedValue)
}

func TestRequestDataStore_Expire(t *testing.T) {

	store := NewRequestDataStore[string, string]()
	key := "testKey"
	value := "testValue"

	err := store.Set(key, value)
	assert.NoError(t, err)

	err = store.Expire(key)
	assert.NoError(t, err)

	retrievedValue, err := store.Get(key)

	assert.Error(t, err)
	assert.Equal(t, "el dato ha expirado", err.Error())
	assert.Equal(t, "", retrievedValue)
}
