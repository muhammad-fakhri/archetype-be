package component

import (
	"github.com/muhammad-fakhri/archetype-be/internal/component/ratelimiter"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/go-libs/cache"
)

func InitRateLimiter(cache cache.Cache, key string) ratelimiter.RateLimiter {
	conf := config.Get().RateLimiter

	c, ok := conf.Config[key]
	if !ok {
		c = ratelimiter.RateLimiterConfig{} // use default config
	}

	return ratelimiter.NewRateLimiter(cache, key, c)
}
