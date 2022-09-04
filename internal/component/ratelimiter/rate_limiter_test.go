package ratelimiter_test

import (
	"sync"
	"testing"
	"time"

	"github.com/muhammad-fakhri/archetype-be/internal/component/ratelimiter"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	"github.com/muhammad-fakhri/go-libs/cache"
)

func TestRateLimiter_Allow(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	rateLimit := int64(ratelimiter.DefaultRateLimit)
	key := "key"

	testLimiter := ratelimiter.NewRateLimiter(m.Redis.Slave, key, ratelimiter.RateLimiterConfig{})

	tests := []struct {
		testName     string
		maxAttempt   int
		expectations func(wg *sync.WaitGroup)
		err          error
	}{
		{
			testName: "Allow() returns nil on counter below rate limit",
			expectations: func(wg *sync.WaitGroup) {
				m.Cache.EXPECT().IncrBy(key, int64(1)).Return(int64(2), nil)

				wg.Add(1)
				m.Cache.EXPECT().Expire(key, ratelimiter.DefaultExpiryTime).DoAndReturn(
					func(key string, ttl time.Duration) (int64, error) {
						wg.Done()
						return 0, nil
					},
				)
			},
		},
		{
			testName:   "Allow() retries until rate is no longer exceeding limit (if possible)",
			maxAttempt: 3,
			expectations: func(wg *sync.WaitGroup) {
				// first attempt, reaches rate limit
				m.Cache.EXPECT().IncrBy(key, int64(1)).Return(rateLimit+1, nil).Times(1)
				m.Cache.EXPECT().DecrWithLimit(key, int64(1), int64(0)).
					Return(rateLimit, nil).Times(1)

				// sleep and do second attempt, assume the counter is decreased by another process
				// time.Sleep(...)
				m.Cache.EXPECT().IncrBy(key, int64(1)).Return(rateLimit, nil)

				wg.Add(1)
				m.Cache.EXPECT().Expire(key, ratelimiter.DefaultExpiryTime).DoAndReturn(
					func(key string, ttl time.Duration) (int64, error) {
						wg.Done()
						return 0, nil
					},
				)
			},
		},
		{
			testName:   "Allow() returns ErrTooManyRequest when max attempt is reached",
			maxAttempt: 3,
			expectations: func(wg *sync.WaitGroup) {
				m.Cache.EXPECT().IncrBy(key, int64(1)).Return(rateLimit+1, nil).Times(3)
				m.Cache.EXPECT().DecrWithLimit(key, int64(1), int64(0)).
					Return(rateLimit, nil).Times(3)
			},
			err: errors.ErrTooManyRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			wg := &sync.WaitGroup{}
			defer wg.Wait()

			test.expectations(wg)

			err := testLimiter.Allow(test.maxAttempt, 0)
			if err != test.err {
				t.Errorf("Allow() fails. expected err: %+v. got: %+v", test.err, err)
			}
		})
	}
}

func TestRateLimiter_Finish(t *testing.T) {
	m := testutil.InitMock(t)
	defer m.Controller.Finish()

	key := "key"

	testLimiter := ratelimiter.NewRateLimiter(m.Redis.Slave, key, ratelimiter.RateLimiterConfig{})

	tests := []struct {
		testName     string
		expectations func()
		err          error
	}{
		{
			testName: "Finish() returns nil",
			expectations: func() {
				m.Cache.EXPECT().DecrWithLimit(key, int64(1), int64(0)).
					Return(int64(0), nil)
			},
		},
		{
			testName: "Finish() returns nil on cache.ErrLimitExceeded",
			expectations: func() {
				m.Cache.EXPECT().DecrWithLimit(key, int64(1), int64(0)).
					Return(int64(-1), cache.ErrLimitExceeded)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			test.expectations()
			err := testLimiter.Finish()
			if err != test.err {
				t.Errorf("Finish() fails. expected err: %+v. got: %+v", test.err, err)
			}
		})
	}
}
