package testutil

import (
	"testing"

	"github.com/muhammad-fakhri/archetype-be/internal/component"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/archetype-be/internal/repository"
	"github.com/muhammad-fakhri/archetype-be/internal/test/mockbridge"
	"github.com/muhammad-fakhri/archetype-be/internal/test/mockcomponent"
	"github.com/muhammad-fakhri/archetype-be/internal/test/mockrepository"
	"github.com/muhammad-fakhri/archetype-be/internal/test/mockusecase"
	"github.com/muhammad-fakhri/go-libs/cache/mock_cache"
	"github.com/muhammad-fakhri/go-libs/log"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
)

type MockComponent struct {
	Controller *gomock.Controller
	Logger     log.SLogger
	Config     *config.Config
	Usecase    *mockusecase.MockUsecase
	Repository *mockrepository.MockRepository
	Bridge     *mockbridge.MockBridge
	// BEGIN __INCLUDE_DB_SQL__
	DB     *component.DB
	DBMock sqlmock.Sqlmock
	// END __INCLUDE_DB_SQL__
	// BEGIN __INCLUDE_DB_MONGO__
	MongoTest *mtest.T
	// END __INCLUDE_DB_MONGO__
	// BEGIN __INCLUDE_REDIS__
	Redis  *component.Redis
	Cache  *mock_cache.MockCache
	Locker *mockcomponent.MockLocker
	// END __INCLUDE_REDIS__
	LocalCache *mockcomponent.MockLocalCache
	// BEGIN __INCLUDE_PUBSUB__
	Publisher *mockcomponent.MockPublisher
	// END __INCLUDE_PUBSUB__
	// BEGIN __INCLUDE_GCS__
	Storage *mockcomponent.MockStorage
	// END __INCLUDE_GCS__
}

func InitMock(t *testing.T) (mock *MockComponent) {
	config.InitMockConfig()
	conf := config.Get()
	mockCtrl := gomock.NewController(t)
	// BEGIN __INCLUDE_DB_SQL__
	db, sqlmock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	// END __INCLUDE_DB_SQL__

	// BEGIN __INCLUDE_REDIS__
	cachemock := mock_cache.NewMockCache(mockCtrl)
	// END __INCLUDE_REDIS__

	return &MockComponent{
		Controller: mockCtrl,
		Logger:     log.NewSLogger(conf.AppName),
		Config:     conf,
		Usecase:    mockusecase.NewMockUsecase(mockCtrl),
		Repository: mockrepository.NewMockRepository(mockCtrl),
		Bridge:     mockbridge.NewMockBridge(mockCtrl),
		// BEGIN __INCLUDE_DB_SQL__
		DB: &component.DB{
			Master: db,
			Slave:  db,
		},
		DBMock: sqlmock,
		// END __INCLUDE_DB_SQL__

		// BEGIN __INCLUDE_DB_MONGO__
		MongoTest: mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock)),
		// END __INCLUDE_DB_MONGO__
		// BEGIN __INCLUDE_REDIS__
		Redis: &component.Redis{
			Master: cachemock,
			Slave:  cachemock,
		},
		Cache:  cachemock,
		Locker: mockcomponent.NewMockLocker(mockCtrl),
		// END __INCLUDE_REDIS__
		LocalCache: mockcomponent.NewMockLocalCache(mockCtrl),
		// BEGIN __INCLUDE_PUBSUB__
		Publisher: mockcomponent.NewMockPublisher(mockCtrl),
		// END __INCLUDE_PUBSUB__
		// BEGIN __INCLUDE_GCS__
		Storage: mockcomponent.NewMockStorage(mockCtrl),
		// END __INCLUDE_GCS__
	}
}

type repositoryParam struct {
	// BEGIN __INCLUDE_PUBSUB__
	publishers *repository.Publishers
	// END __INCLUDE_PUBSUB__
}

type RepositoryOption func(*repositoryParam)

// BEGIN __INCLUDE_PUBSUB__
func WithPublishers(param *repository.Publishers) RepositoryOption {
	return func(r *repositoryParam) {
		r.publishers = param
	}
}

// END __INCLUDE_PUBSUB__
func InitMockRepository(m *MockComponent, opts ...RepositoryOption) repository.Repository {
	opt := &repositoryParam{}

	for _, o := range opts {
		o(opt)
	}

	return repository.NewRepository(m.Logger, m.Config, m.Bridge, m.LocalCache,
		// BEGIN __INCLUDE_DB_SQL__
		m.DB,
		// END __INCLUDE_DB_SQL__
		// BEGIN __INCLUDE_DB_MONGO__
		&component.MongoDB{
			Master: m.MongoTest.DB,
		},
		// END __INCLUDE_DB_MONGO__
		// BEGIN __INCLUDE_REDIS__
		m.Redis,
		m.Locker,
		// END __INCLUDE_REDIS__
		// BEGIN __INCLUDE_PUBSUB__
		opt.publishers,
		// END __INCLUDE_PUBSUB__
		// BEGIN __INCLUDE_GCS__
		m.Storage,
		// END __INCLUDE_GCS__
	)
}
