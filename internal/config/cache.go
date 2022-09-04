package config

import (
	"time"

	"github.com/muhammad-fakhri/archetype-be/internal/util/stdutil"
	"github.com/muhammad-fakhri/go-libs/cache"
)

var (
	// BEGIN __INCLUDE_REDIS__
	redisDefaultTTLSec = 300

	redisEventDetailTTLSec = 900
	redisUserAuthTTLSec    = 600
	// END __INCLUDE_REDIS__

	localCacheDefaultTTLSec = 300
)

// BEGIN __INCLUDE_REDIS__
type RedisTTL struct {
	Default     time.Duration
	UserAuth    time.Duration
	EventDetail time.Duration
	// add more below
}

// END __INCLUDE_REDIS__
type LocalCacheTTL struct {
	Default time.Duration
	// add more below
}

func (c *Config) initCacheConfig(cfg *configIni) {
	// BEGIN __INCLUDE_REDIS__
	appConfig.RedisMaster = &cache.Config{
		ServerAddr:      cfg.RedisMasterAddr,
		IdleTimeout:     time.Duration(cfg.RedisIdleTimeout) * time.Second,
		MaxActive:       cfg.RedisMaxActive,
		MaxConnLifetime: time.Duration(cfg.RedisMaxConnLifetime) * time.Second,
		MaxIdle:         cfg.RedisMaxIdle,
		Wait:            cfg.RedisWait,
		UseCommonErr:    true,
	}
	appConfig.RedisSlave = &cache.Config{
		ServerAddr:      cfg.RedisSlaveAddr,
		IdleTimeout:     time.Second * time.Duration(cfg.RedisIdleTimeout),
		MaxActive:       cfg.RedisMaxActive,
		MaxConnLifetime: time.Second * time.Duration(cfg.RedisMaxConnLifetime),
		MaxIdle:         cfg.RedisMaxIdle,
		Wait:            cfg.RedisWait,
		UseCommonErr:    true,
	}

	appConfig.RedisTTL = &RedisTTL{
		Default:     stdutil.GetTimeSecondOrDefault(cfg.RedisDefaultTTL, redisDefaultTTLSec),
		UserAuth:    stdutil.GetTimeSecondOrDefault(cfg.RedisUserAuthTTL, redisUserAuthTTLSec),
		EventDetail: stdutil.GetTimeSecondOrDefault(cfg.RedisEventDetailTTL, redisEventDetailTTLSec),
	}
	// END __INCLUDE_REDIS__

	appConfig.LocalCacheTTL = &LocalCacheTTL{
		Default: stdutil.GetTimeSecondOrDefault(cfg.LocalCacheDefaultTTL, localCacheDefaultTTLSec),
	}
}
