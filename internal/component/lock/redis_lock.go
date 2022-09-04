package lock

import (
	"strconv"
	"time"

	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	"github.com/muhammad-fakhri/go-libs/cache"
)

type redisLock struct {
	cache  cache.Cache
	config *redisLockConfig
}

type redisLockConfig struct {
	enabled bool
}

func NewRedisLock(cache cache.Cache, enabled bool) Locker {
	return &redisLock{
		config: &redisLockConfig{
			enabled: enabled,
		},
		cache: cache,
	}
}

// WaitForLock waits and retries until value pointed by redis key is empty. Default MaxRetryCount = 1.
func (c *redisLock) WaitForLock(key string, maxRetryCount int, sleepDuration time.Duration) (err error) {
	if !c.config.enabled {
		return
	}

	retryThreshold := 1
	if maxRetryCount > 0 {
		retryThreshold = maxRetryCount
	}

	for retryCount := 0; retryCount < retryThreshold; retryCount++ {
		val, err := c.cache.Get(key)
		// stop and return on ErrNil (redis key-value does not exists)
		if err == cache.ErrNil {
			return nil
		}
		// return non nil / non err nil errors
		if err != nil {
			return err
		}
		// stop and return on nil error and empty value
		if val == "" || val == "null" {
			return nil
		}
		// skip wait on last attempt
		if retryCount == retryThreshold-1 {
			continue
		}
		time.Sleep(sleepDuration)
	}

	return errors.ErrTooManyRequest
}

// SetLockWithWait sets lock in place and wait if the lock is not available. Default MaxRetryCount = 1.
func (c *redisLock) SetLockWithWait(key string, ttl time.Duration, maxRetryCount int, sleepDuration time.Duration) (err error) {
	if !c.config.enabled {
		return
	}

	retryThreshold := 1
	if maxRetryCount > 0 {
		retryThreshold = maxRetryCount
	}

	for retryCount := 0; retryCount < retryThreshold; retryCount++ {
		err := c.cache.SetNX(key, strconv.FormatBool(true), ttl)
		// return any non NX error
		if err != cache.ErrNX {
			return err
		}
		// skip wait on last attempt
		if retryCount == retryThreshold-1 {
			continue
		}
		time.Sleep(sleepDuration)
	}

	return errors.ErrTooManyRequest
}

// ReleaseLock releases the lock associated with the given key
func (c *redisLock) ReleaseLock(key string) (err error) {
	if !c.config.enabled {
		return
	}

	return c.cache.Del(key)
}
