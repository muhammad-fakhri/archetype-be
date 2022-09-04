package config

import "github.com/muhammad-fakhri/archetype-be/pkg/errors"

type Log struct {
	Debug bool
}

func (c *Config) LogDebug() bool {
	if c == nil || c.Log == nil {
		return false
	}

	return c.Log.Debug
}

func (c *Config) initLogConfig(cfg *configIni) {
	appConfig.Log = &Log{
		Debug: cfg.LogDebug,
	}

	errors.EnableDebug(cfg.LogDebug)
}
