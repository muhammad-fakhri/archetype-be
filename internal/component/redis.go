package component

import (
	"log"

	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/go-libs/cache"
)

// Redis component
type Redis struct {
	Master cache.Cache
	Slave  cache.Cache
}

// InitRedis to initialize redis
func InitRedis() *Redis {
	conf := config.Get()

	if conf.RedisMaster == nil || conf.RedisSlave == nil {
		log.Fatalf("failed to get redis config")
	}
	redis := &Redis{}

	redisMaster, err := cache.New(cache.Redis, conf.RedisMaster)
	if err != nil {
		log.Fatalf("failed to open redis master connection. %+v", err)
	}
	redisSlave, err := cache.New(cache.Redis, conf.RedisSlave)
	if err != nil {
		log.Fatalf("failed to open redis slave connection. %+v", err)
	}

	redis.Master = redisMaster
	redis.Slave = redisSlave

	return redis
}
