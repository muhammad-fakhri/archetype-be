package component

import (
	"time"

	"github.com/muhammad-fakhri/archetype-be/internal/component/localcache"
)

type lcacheParam struct {
	ttl time.Duration
}

type LCacheOption func(*lcacheParam)

func WithLCacheTTL(ttl time.Duration) LCacheOption {
	return func(c *lcacheParam) {
		c.ttl = ttl
	}
}

func InitLocalCache(opts ...LCacheOption) localcache.LocalCache {
	opt := &lcacheParam{}
	for _, o := range opts {
		o(opt)
	}

	return localcache.NewMemoryCache(opt.ttl)
}
