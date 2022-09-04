package config

import "fmt"

// BEGIN __INCLUDE_DB_SQL__
const (
	dbStringConnection = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4"
)

type sqlDB struct {
	ConnectionString string
	MaxIdle          int
	MaxOpen          int
}

func (c *Config) initSqlDBConfig(cfg *configIni) {
	appConfig.DBMaster = &sqlDB{
		ConnectionString: fmt.Sprintf(
			dbStringConnection,
			cfg.DBMasterUser,
			cfg.DBMasterPass,
			cfg.DBMasterHost,
			cfg.DBMasterPort,
			cfg.DBMasterName,
		),
		MaxIdle: cfg.DBMasterMaxIdle,
		MaxOpen: cfg.DBMasterMaxOpen,
	}

	appConfig.DBSlave = &sqlDB{
		ConnectionString: fmt.Sprintf(
			dbStringConnection,
			cfg.DBSlaveUser,
			cfg.DBSlavePass,
			cfg.DBSlaveHost,
			cfg.DBSlavePort,
			cfg.DBSlaveName,
		),
		MaxIdle: cfg.DBMasterMaxIdle,
		MaxOpen: cfg.DBMasterMaxOpen,
	}
}

// END __INCLUDE_DB_SQL__
