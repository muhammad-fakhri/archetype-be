package config

import (
	"strings"

	"github.com/muhammad-fakhri/archetype-be/internal/util/stdutil"
	"github.com/muhammad-fakhri/go-libs/constant"
)

func (c *Config) initCommonConfig(cfg *configIni) {
	c.Environment = constant.Environment(cfg.Environment)
	c.AppName = cfg.AppName
	c.AppSecret = cfg.AppSecret
	c.HttpPort = stdutil.GetStringOrDefault(cfg.HttpPort, "8080")
	countries := strings.Split(cfg.CountryFilter, ",")
	for _, country := range countries {
		c.CountryFilter = append(c.CountryFilter, constant.Country(country))
	}
}
