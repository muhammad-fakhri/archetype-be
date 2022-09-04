package cacheutil

// BEGIN __INCLUDE_REDIS__
import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func GetJSON(key string, out interface{}, cacheFunc func(key string) (string, error)) error {
	val, err := cacheFunc(key)
	if (val != "" && val != "null") && err == nil {
		err = json.Unmarshal([]byte(val), out)
		return err
	}
	return errors.Wrap(errors.ErrRedis, fmt.Sprintf("%s", err))
}

func SetJSON(key string, in interface{}, ttl time.Duration, cacheFunc func(key, value string, ttl time.Duration) error) error {
	payload, err := json.Marshal(in)
	if err != nil {
		return errors.Wrap(errors.ErrRedis, err.Error())
	}

	return cacheFunc(key, string(payload), ttl)
}

// END __INCLUDE_REDIS__
