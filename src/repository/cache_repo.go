package repository

// ICacheRepository : Cache Repository
type ICacheRepository interface {
	Set(key string, value interface{}, expirationTimeSeconds uint) bool
	Get(key string) (item interface{}, isFound bool)
}

// RedisCacheRepository : caches via redis cache
type RedisCacheRepository struct {
}

func (r *RedisCacheRepository) set(key string, value interface{}, expirationTimeSeconds uint) bool {
	return false
}

func (r *RedisCacheRepository) get(key string) (item interface{}, isFound bool) {
	return nil, false
}
