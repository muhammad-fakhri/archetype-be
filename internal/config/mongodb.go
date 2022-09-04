package config

// BEGIN __INCLUDE_DB_MONGO__
import (
	"fmt"
	"strings"
	"time"
)

const separator = "|"

// DBConfig represents the configuration map for MongoDB client.
type MongoDBConfig struct {
	DBAppName string

	MongoDBUser          string
	MongoDBPass          string
	MongoDBHost          string
	MongoDBPort          string
	MongoDBName          string
	MongoDBOptions       string
	MongoDBMinConnection uint64
	MongoDBMaxConnection uint64

	MongoDBConnectionURI string

	MongoDBTimeout         time.Duration
	MongoDBMaxConnIdleTime time.Duration
	MongoDBMaxStaleness    time.Duration
}

func (c *Config) initMongoDB(cfg *configIni) {
	appConfig.MongoDB = &MongoDBConfig{
		DBAppName: cfg.AppName,

		MongoDBUser:          cfg.MongoDBUser,
		MongoDBPass:          cfg.MongoDBPass,
		MongoDBHost:          cfg.MongoDBHost,
		MongoDBPort:          cfg.MongoDBPort,
		MongoDBName:          cfg.MongoDBName,
		MongoDBOptions:       cfg.MongoDBOptions,
		MongoDBMinConnection: cfg.MongoDBMinConnection,
		MongoDBMaxConnection: cfg.MongoDBMaxConnection,

		MongoDBConnectionURI: CreateMongoConnURI(
			cfg.MongoDBUser,
			cfg.MongoDBPass,
			strings.Split(cfg.MongoDBHost, separator),
			strings.Split(cfg.MongoDBPort, separator),
			cfg.MongoDBOptions,
		),

		MongoDBTimeout:         time.Duration(cfg.MongoDBTimeoutMS) * time.Millisecond,
		MongoDBMaxConnIdleTime: time.Duration(cfg.MongoDBMaxConnIdleTimeMS) * time.Millisecond,
		MongoDBMaxStaleness:    time.Duration(cfg.MongoDBMaxStalenessMS) * time.Millisecond,
	}
}

// CreateMongoConnURI constructs mongo connection URI from params.
func CreateMongoConnURI(user, pass string, host, port []string, opt string) string {
	connURIFmt := "mongodb://"
	var params []interface{}

	if user != "" {
		connURIFmt += "%s"
		params = append(params, user)

		if pass != "" {
			connURIFmt += ":%s"
			params = append(params, pass)
		}
		connURIFmt += "@"
	}

	if len(host) == len(port) {
		var hostPort []string
		for i := range host {
			hostPort = append(hostPort, fmt.Sprintf("%s:%s", host[i], port[i]))
		}
		connURIFmt += strings.Join(hostPort, ",")
	}

	connURIFmt += "/"

	if opt != "" {
		connURIFmt += "?%s"
		params = append(params, opt)
	}

	return fmt.Sprintf(connURIFmt, params...)
}

// END __INCLUDE_DB_MONGO__
