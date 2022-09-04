package ratelimiter

import (
	"log"
	"time"

	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	"github.com/muhammad-fakhri/go-libs/cache"
)

const (
	DefaultRateLimit  = 100
	DefaultExpiryTime = 5 * time.Second
)

type rateLimiter struct {
	cache      cache.Cache
	rateLimit  int64
	expiryTime time.Duration
	key        string
}

type RateLimiterConfig struct {
	RateLimit  int64
	ExpiryTime time.Duration
}

func NewRateLimiter(cache cache.Cache, key string, config RateLimiterConfig) RateLimiter {
	if len(key) == 0 {
		log.Fatalf("failed to init rate limiter. err:invalid key")
	}

	limiter := &rateLimiter{
		cache:      cache,
		key:        key,
		rateLimit:  config.RateLimit,
		expiryTime: config.ExpiryTime,
	}
	if limiter.rateLimit == 0 {
		limiter.rateLimit = DefaultRateLimit
	}
	if limiter.expiryTime == 0 {
		limiter.expiryTime = DefaultExpiryTime
	}
	return limiter
}

// Allow checks for rate limit, returns ErrTooManyRequest if rate limit is exceeded after max attempts.
func (r *rateLimiter) Allow(maxAttempt int, attemptWaitDuration time.Duration) error {

	if maxAttempt <= 0 {
		maxAttempt = 1
	}

	for attempt := 1; attempt <= maxAttempt; attempt++ {
		// increase rate counter
		currentRate, err := r.cache.IncrBy(r.key, 1)
		if err != nil {
			return err
		}
		// if current rate exceeds rate limit, revert back the incr and retry
		if currentRate > r.rateLimit {
			_, err = r.cache.DecrWithLimit(r.key, 1, 0)
			if err != nil {
				return err
			}
			// skip sleep on last attempt
			if attempt < maxAttempt {
				time.Sleep(attemptWaitDuration)
			}
			continue
		}

		// expire this key in case caller was down before
		// caller could call rateLimiter.Finish()
		go r.cache.Expire(r.key, r.expiryTime)
		return nil
	}

	// return too many request error when max attempt is reached
	return errors.ErrTooManyRequest
}

// Finish clears up the allocated rate limit (1) by decreasing the counter.
// Caller must call Finish() after Allow() calls.
func (r *rateLimiter) Finish() error {
	_, err := r.cache.DecrWithLimit(r.key, 1, 0)
	if err == cache.ErrLimitExceeded {
		err = nil
	}
	return err
}
