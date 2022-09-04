package config

// AuthConfig represents configurations for this service's authentication.
type AuthConfig struct {
	// play admin cms authentication key
	JWTSecret string
}

func (c *Config) initAuthConfig(cfg *configIni) {
	appConfig.AuthConfig = &AuthConfig{
		JWTSecret: cfg.JWTSecret,
	}
}
