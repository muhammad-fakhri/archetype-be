package config

import "github.com/muhammad-fakhri/go-libs/email"

// BEGIN __INCLUDE_EMAIL__
type EmailConfig struct {
	Credential email.Config
}

func (c *Config) initEmailConfig(cfg *configIni) {
	c.Email = &EmailConfig{
		Credential: email.Config{
			Email:    cfg.SenderEmailAddress,
			Password: cfg.SenderPassword,
		},
	}
}

// END __INCLUDE_EMAIL__
