package repository

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// ICacheRepository : Cache Repository
type ICacheRepository interface {
	// uses set expiration
	Set(key string, value interface{}, expirationTimeSeconds uint) bool
	// uses default expiration
	SetDefaultExpiration(key string, value interface{}) bool
	// sets items with no expiry , use this with caution please
	SetNoExpiry(key string, value interface{}) bool
	// gets key or returns not found
	Get(key string) (item interface{}, isFound bool)
}

// InMemoryCache : Memory Cache using ram
type InMemoryCache struct {
	Cacher *cache.Cache
}

// Gets the item from in memory cache
func (i *InMemoryCache) Get(key string) (item interface{}, isFound bool) {
	return i.Cacher.Get(key)
}

// Sets the item in memory cache
func (i *InMemoryCache) Set(key string, value interface{}, expirationTimeSeconds uint) bool {
	i.Cacher.Set(key,
		value,
		time.Duration(expirationTimeSeconds)*time.Second,
	)
	return true
}

// SetDefaultExpiration : sets an item to expire with server enviroment set expiration time
func (i *InMemoryCache) SetDefaultExpiration(key string, value interface{}) bool {
	i.Cacher.Set(key,
		value,
		cache.DefaultExpiration,
	)
	return true
}

// SetNoExpiry : Adds an item with no expiry
func (i *InMemoryCache) SetNoExpiry(key string, value interface{}) bool {
	i.Cacher.Set(key,
		value,
		cache.NoExpiration,
	)
	return true
}
