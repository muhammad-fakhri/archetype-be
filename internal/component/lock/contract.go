package lock

import (
	"time"
)

type Locker interface {
	WaitForLock(key string, maxRetryCount int, sleepDuration time.Duration) (err error)
	SetLockWithWait(key string, ttl time.Duration, maxRetryCount int, sleepDuration time.Duration) (err error)
	ReleaseLock(key string) (err error)
}
