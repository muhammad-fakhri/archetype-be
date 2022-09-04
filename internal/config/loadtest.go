package config

type LoadTestConfig struct {
	Enabled bool
}

func (c *Config) IsLoadTest() bool {
	if c == nil || c.LoadTest == nil {
		return false
	}

	return c.LoadTest.Enabled
}

func (c *Config) initLoadTestConfig(cfg *configIni) {
	appConfig.LoadTest = &LoadTestConfig{
		Enabled: cfg.LoadTestEnabled,
	}
}
