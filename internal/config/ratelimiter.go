package config

import (
	"github.com/muhammad-fakhri/archetype-be/internal/component/ratelimiter"
)

type RateLimiterConfig struct {
	Config map[string]ratelimiter.RateLimiterConfig
}

func (c *Config) initRateLimiterConfig(cfg *configIni) {
	c.RateLimiter = &RateLimiterConfig{
		Config: map[string]ratelimiter.RateLimiterConfig{},
	}
}
