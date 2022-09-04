package localcache

import (
	"time"

	"github.com/koding/cache"
)

type memoryCache struct {
	memory cache.Cache
}

func NewMemoryCache(ttl time.Duration) LocalCache {
	var memory cache.Cache
	if ttl > 0 {
		memory = cache.NewMemoryWithTTL(ttl)
	} else {
		memory = cache.NewMemory()
	}

	return &memoryCache{
		memory: memory,
	}
}

// implementation
func (c *memoryCache) Get(key string) (interface{}, error) {
	return c.memory.Get(key)
}

func (c *memoryCache) Set(key string, value interface{}) error {
	return c.memory.Set(key, value)
}

func (c *memoryCache) Delete(key string) error {
	return c.memory.Delete(key)
}
