package repository

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
