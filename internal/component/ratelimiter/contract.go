package ratelimiter

import (
	"time"
)

// RateLimiter represents the interface of rate limiter.
type RateLimiter interface {
	// Allow checks for rate limit and increases the rate counter. Allow returns ErrTooManyRequest
	// if rate limit is still exceeded after max attempt count has been reached.
	Allow(maxAttempt int, attemptWaitDuration time.Duration) error
	// Finish clears up the allocated rate limit (1) by decreasing the counter.
	// Caller must call Finish() after Allow() calls.
	Finish() error
}
