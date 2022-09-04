package config

import (
	"log"

	"github.com/muhammad-fakhri/go-libs/cache"
	"github.com/muhammad-fakhri/go-libs/constant"
	"gopkg.in/ini.v1"
)

const (
	DefaultConfigPath = "conf/"
)

type Config struct {
	// app
	Environment   constant.Environment
	AppName       string
	AppSecret     string
	HttpPort      string
	CountryFilter []constant.Country

	AuthConfig *AuthConfig
	// BEGIN __INCLUDE_DB_SQL__
	DBMaster *sqlDB
	DBSlave  *sqlDB
	// END __INCLUDE_DB_SQL__
	// BEGIN __INCLUDE_DB_MONGO__
	MongoDB *MongoDBConfig
	// END __INCLUDE_DB_MONGO__
	// BEGIN __INCLUDE_REDIS__
	RedisMaster *cache.Config
	RedisSlave  *cache.Config
	RedisTTL    *RedisTTL
	RateLimiter *RateLimiterConfig
	// END __INCLUDE_REDIS__
	LocalCacheTTL *LocalCacheTTL
	Log           *Log
	Scheduler     *Scheduler
	// BEGIN __INCLUDE_EMAIL__
	Email *EmailConfig
	// END __INCLUDE_EMAIL__
	// BEGIN __INCLUDE_GCS__
	GCS *GCSConfig
	// END __INCLUDE_GCS__
	LoadTest *LoadTestConfig
	// BEGIN __INCLUDE_PUBSUB__
	PubSub *PubSubConfig
	// END __INCLUDE_PUBSUB__
}

type configIni struct {
	// App config
	Environment   string `ini:"environment"`
	AppName       string `ini:"appname"`
	AppSecret     string `ini:"appsecret"`
	HttpPort      string `ini:"httpport"`
	CountryFilter string `ini:"countryfilter"`

	// Play admin config
	JWTSecret string `ini:"jwtsecret"`
	// BEGIN __INCLUDE_DB_SQL__
	// Database config
	DBMasterMaxIdle int    `ini:"dbmaster_maxidle"`
	DBMasterMaxOpen int    `ini:"dbmaster_maxopen"`
	DBSlaveMaxIdle  int    `ini:"dbslave_maxidle"`
	DBSlaveMaxOpen  int    `ini:"dbslave_maxopen"`
	DBMasterUser    string `ini:"dbmaster_user"`
	DBMasterPass    string `ini:"dbmaster_pass"`
	DBMasterHost    string `ini:"dbmaster_host"`
	DBMasterPort    string `ini:"dbmaster_port"`
	DBMasterName    string `ini:"dbmaster_name"`
	DBSlaveUser     string `ini:"dbslave_user"`
	DBSlavePass     string `ini:"dbslave_pass"`
	DBSlaveHost     string `ini:"dbslave_host"`
	DBSlavePort     string `ini:"dbslave_port"`
	DBSlaveName     string `ini:"dbslave_name"`
	// END __INCLUDE_DB_SQL__
	// BEGIN __INCLUDE_DB_MONGO__
	// MongoDB Config
	MongoDBUser              string `ini:"mongodb_user"`
	MongoDBPass              string `ini:"mongodb_pass"`
	MongoDBHost              string `ini:"mongodb_host"`
	MongoDBPort              string `ini:"mongodb_port"`
	MongoDBName              string `ini:"mongodb_name"`
	MongoDBOptions           string `ini:"mongodb_options"`
	MongoDBMaxConnection     uint64 `ini:"mongodb_maxconnection"`
	MongoDBMinConnection     uint64 `ini:"mongodb_minconnection"`
	MongoDBTimeoutMS         int    `ini:"mongodb_timeoutms"`
	MongoDBMaxConnIdleTimeMS int    `ini:"mongodb_maxconnidletimems"`
	MongoDBMaxStalenessMS    int    `ini:"mongodb_maxstalenessms"`
	// END __INCLUDE_DB_MONGO__
	// BEGIN __INCLUDE_REDIS__
	// Redis config
	RedisMasterAddr      string `ini:"redismaster_addr"`
	RedisSlaveAddr       string `ini:"redisslave_addr"`
	RedisMaxIdle         int    `ini:"redis_maxidle"`
	RedisMaxActive       int    `ini:"redis_maxactive"`
	RedisIdleTimeout     int64  `ini:"redis_idletimeoutsec"`
	RedisMaxConnLifetime int64  `ini:"redis_maxconnlifetimesec"`
	RedisWait            bool   `ini:"redis_wait"`

	// Cache TTL
	RedisDefaultTTL     int `ini:"redisttl_defaultsec"`
	RedisUserAuthTTL    int `ini:"redisttl_userauthsec"`
	RedisEventDetailTTL int `ini:"redisttl_eventdetailsec"`
	// END __INCLUDE_REDIS__
	LocalCacheDefaultTTL int `ini:"localcachettl_defaultsec"`

	// Log Config
	LogDebug bool `ini:"log_debug"`

	// Scheduler
	// BEGIN __INCLUDE_EXAMPLE_CRON__
	CronSpecEventExampleUpdater string `ini:"cronspec_eventupdater"`
	// END __INCLUDE_EXAMPLE_CRON__

	// BEGIN __INCLUDE_EMAIL__
	// MailConfig
	SenderEmailAddress string `ini:"mail_senderaddress"`
	SenderPassword     string `ini:"mail_senderpassword"`
	// END __INCLUDE_EMAIL__

	// BEGIN __INCLUDE_GCS__
	// GCSConfig
	GCSCredentialFileName       string `ini:"gcs_credentialfilename"`
	GCSBucketName               string `ini:"gcs_bucketname"`
	GCSDownloadURLExpiryTimeSec int    `ini:"gcs_downloadurlexpiryttlsec"`
	GCSUploadURLExpiryTimeSec   int    `ini:"gcs_uploadurlexpiryttlsec"`
	GCSReadBatchSize            int64  `ini:"gcs_readbatchsize"`
	// END __INCLUDE_GCS__

	// LoadTestConfig
	LoadTestEnabled bool `ini:"loadtest_enabled"`

	// BEGIN __INCLUDE_PUBSUB__
	// PubsubConfig
	// common
	PubSubCredentialFilename string `ini:"pubsub_credentialfilename"`
	PubSubProjectID          string `ini:"pubsub_projectid"`
	PubSubPrefixID           string `ini:"pubsub_prefixid"`
	PubSubPublishAsync       bool   `ini:"pubsub_publishasync"`
	// common subscription config
	PubSubAckDeadline      int `ini:"pubsub_subscriptionackdeadline"`
	PubSubExpirationPolicy int `ini:"pubsub_subscriptionexpiration"`
	// deadletter config
	PubSubDeadLetterMaxRedelivery int `ini:"pubsub_deadletterredelivery"`
	// receive concurrency settings
	PubSubNumOfGoroutine        int `ini:"pubsub_subscriptionthread"`
	PubSubMaxOutstandingMessage int `ini:"pubsub_subscriptionmaxmessage"`
	// END __INCLUDE_PUBSUB__

}

var appConfig *Config

func Init() {
	var err error

	c := &configIni{}
	cIni, err := ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("[Init] failed to read config, %+v\n", err)
	}
	err = cIni.MapTo(c)
	if err != nil {
		log.Fatalf("[Init] failed to map config, %+v\n", err)
	}

	// App config
	appConfig = &Config{}
	appConfig.initCommonConfig(c)
	appConfig.initAuthConfig(c)
	appConfig.initLogConfig(c)
	appConfig.initSchedulerConfig(c)
	// BEGIN __INCLUDE_DB_SQL__
	appConfig.initSqlDBConfig(c)
	// END __INCLUDE_DB_SQL__
	appConfig.initCacheConfig(c)
	// BEGIN __INCLUDE_REDIS__
	appConfig.initRateLimiterConfig(c)
	// END __INCLUDE_REDIS__
	// BEGIN __INCLUDE_EMAIL__
	appConfig.initEmailConfig(c)
	// END __INCLUDE_EMAIL__
	// BEGIN __INCLUDE_GCS__
	appConfig.initGCSConfig(c)
	// END __INCLUDE_GCS__
	appConfig.initLoadTestConfig(c)
	// BEGIN __INCLUDE_PUBSUB__
	appConfig.initPubSubConfig(c)
	// END __INCLUDE_PUBSUB__
	// BEGIN __INCLUDE_DB_MONGO__
	appConfig.initMongoDB(c)
	// END __INCLUDE_DB_MONGO__
}

// for unit test purpose
func InitMockConfig() {
	appConfig = &Config{}
	appConfig.CountryFilter = []constant.Country{constant.ID, constant.MY}
	appConfig.AuthConfig = &AuthConfig{
		JWTSecret: "fakhri1234",
	}
	appConfig.Environment = constant.EnvTest
	appConfig.AppName = "unit-test"
	// BEGIN __INCLUDE_REDIS__
	appConfig.RedisTTL = &RedisTTL{}
	// END __INCLUDE_REDIS__
}

func Get() *Config {
	return appConfig
}
