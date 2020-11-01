package repository

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// ICacheRepository : Cache Repository
type ICacheRepository interface {
	Set(key string, value interface{}, expirationTimeSeconds uint) bool
	Get(key string) (item interface{}, isFound bool)
}

// RedisCacheRepository : caches via redis cache , TODO
type RedisCacheRepository struct {
}

// Set  : set in redis the cache required
func (r *RedisCacheRepository) Set(key string, value interface{}, expirationTimeSeconds uint) bool {
	return false
}

// Get : Get item from redis via key
func (r *RedisCacheRepository) Get(key string) (item interface{}, isFound bool) {
	return nil, false
}

// InMemoryCache : Memory Cache using ram
type InMemoryCache struct {
	Cacher *cache.Cache
}

// Gets the item from in memory cache
func (i *InMemoryCache) Get(key string) (item interface{}, isFound bool) {
	return i.Cacher.Get("key")
}

// Sets the item in memory cache
func (i *InMemoryCache) Set(key string, value interface{}, expirationTimeSeconds uint) bool {
	i.Cacher.Set(key,
		value,
		time.
			Duration(expirationTimeSeconds)*time.Second,
	)
	return true
}
