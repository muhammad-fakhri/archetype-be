package repository

import (
	"github.com/muhammad-fakhri/archetype-be/internal/component"
	"github.com/muhammad-fakhri/archetype-be/internal/component/localcache"
	"github.com/muhammad-fakhri/archetype-be/internal/component/lock"
	"github.com/muhammad-fakhri/archetype-be/internal/component/pubsub"
	"github.com/muhammad-fakhri/archetype-be/internal/component/storage"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/repository/bridge"
	"github.com/muhammad-fakhri/go-libs/log"
	"golang.org/x/sync/singleflight"
)

type repository struct {
	logger     log.SLogger
	bridge     bridge.Bridge
	localCache localcache.LocalCache
	// BEGIN __INCLUDE_DB_SQL__
	db *component.DB
	// END __INCLUDE_DB_SQL__
	// BEGIN __INCLUDE_DB_MONGO__
	mongodb *component.MongoDB
	// END __INCLUDE_DB_MONGO__
	// BEGIN __INCLUDE_REDIS__
	redis  *component.Redis
	locker lock.Locker
	// END __INCLUDE_REDIS__
	// BEGIN __INCLUDE_GCS__
	storage storage.Storage
	// END __INCLUDE_GCS__
	// BEGIN __INCLUDE_PUBSUB__
	publisher *Publishers
	// END __INCLUDE_PUBSUB__
	requestGroup *singleflight.Group
	config       *repositoryConfig
}

type repositoryConfig struct {
	// BEGIN __INCLUDE_REDIS__
	redisTTL *config.RedisTTL
	// END __INCLUDE_REDIS__
	// BEGIN __INCLUDE_EXAMPLE__
	appName       string
	cmsAuditEmail []string
	// END __INCLUDE_EXAMPLE__
}

func NewRepository(
	logger log.SLogger,
	conf *config.Config,
	bridge bridge.Bridge,
	localCache localcache.LocalCache,
	// BEGIN __INCLUDE_DB_SQL__
	db *component.DB,
	// END __INCLUDE_DB_SQL__
	// BEGIN __INCLUDE_DB_MONGO__
	mongodb *component.MongoDB,
	// END __INCLUDE_DB_MONGO__
	// BEGIN __INCLUDE_REDIS__
	redis *component.Redis,
	locker lock.Locker,
	// END __INCLUDE_REDIS__
	// BEGIN __INCLUDE_PUBSUB__
	publisher *Publishers,
	// END __INCLUDE_PUBSUB__
	// BEGIN __INCLUDE_GCS__
	storage storage.Storage,
	// END __INCLUDE_GCS__
) Repository {
	return &repository{
		logger:     logger,
		bridge:     bridge,
		localCache: localCache,
		// BEGIN __INCLUDE_DB_SQL__
		db: db,
		// END __INCLUDE_DB_SQL__
		// BEGIN __INCLUDE_DB_MONGO__
		mongodb: mongodb,
		// END __INCLUDE_DB_MONGO__
		// BEGIN __INCLUDE_REDIS__
		redis:  redis,
		locker: locker,
		// END __INCLUDE_REDIS__
		requestGroup: &singleflight.Group{},
		// BEGIN __INCLUDE_PUBSUB__
		publisher: publisher,
		// END __INCLUDE_PUBSUB__
		// BEGIN __INCLUDE_GCS__
		storage: storage,
		// END __INCLUDE_GCS__
		config: &repositoryConfig{
			// BEGIN __INCLUDE_REDIS__
			redisTTL: conf.RedisTTL,
			// END __INCLUDE_REDIS__
			// BEGIN __INCLUDE_EXAMPLE__
			appName: conf.AppName,
			//cmsAuditEmail: []string{"some_email"},
			// END __INCLUDE_EXAMPLE__
		},
	}
}

func InitDependencies(log log.SLogger) Repository {
	conf := config.Get()

	lcache := component.InitLocalCache(component.WithLCacheTTL(conf.LocalCacheTTL.Default))
	// BEGIN __INCLUDE_REDIS__
	redis := component.InitRedis()
	// END __INCLUDE_REDIS__

	bridge := bridge.NewBridge(
		// BEGIN __INCLUDE_EMAIL__
		component.InitEmailClient(),
		// END __INCLUDE_EMAIL__
	)

	return NewRepository(log, conf, bridge, lcache,
		// BEGIN __INCLUDE_DB_SQL__
		component.InitDatabase(),
		// END __INCLUDE_DB_SQL__
		// BEGIN __INCLUDE_DB_MONGO__
		component.InitMongoDB(),
		// END __INCLUDE_DB_MONGO__
		// BEGIN __INCLUDE_REDIS__
		redis,
		component.InitLocker(redis.Slave),
		// END __INCLUDE_REDIS__
		// BEGIN __INCLUDE_PUBSUB__
		InitPublishers(),
		// END __INCLUDE_PUBSUB__
		// BEGIN __INCLUDE_GCS__
		component.InitStorageClient(),
		// END __INCLUDE_GCS__
	)
}

// BEGIN __INCLUDE_PUBSUB__
type Publishers struct {
	// BEGIN __INCLUDE_EXAMPLE__
	UpdateSystemConfig pubsub.Publisher
	// END __INCLUDE_EXAMPLE__

	// add more publishers here
}

func InitPublishers() *Publishers {
	// init pubsub client
	// BEGIN __INCLUDE_EXAMPLE__
	c := component.InitPubSubClient()
	// END __INCLUDE_EXAMPLE__
	return &Publishers{
		// BEGIN __INCLUDE_EXAMPLE__
		UpdateSystemConfig: c.NewPublisher(c.GetTopicID(constant.EventExampleUpdateSystemConfig)),
		// END __INCLUDE_EXAMPLE__
	}
}

// END __INCLUDE_PUBSUB__
