package component

import (
	"github.com/muhammad-fakhri/archetype-be/internal/component/lock"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/go-libs/cache"
)

func InitLocker(cache cache.Cache) lock.Locker {
	conf := config.Get()
	return lock.NewRedisLock(cache, !conf.IsLoadTest())
}
