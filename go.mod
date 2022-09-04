module github.com/muhammad-fakhri/archetype-be

go 1.13

require (
	cloud.google.com/go/pubsub v1.18.0
	cloud.google.com/go/storage v1.18.2
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0 // indirect
	github.com/json-iterator/go v1.1.12
	github.com/julienschmidt/httprouter v1.3.0
	github.com/klauspost/compress v1.13.5 // indirect
	github.com/koding/cache v0.0.0-20161222233015-e8a81b0b3f20
	github.com/mitchellh/mapstructure v1.4.3
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/muhammad-fakhri/go-libs/authz v1.0.0
	github.com/muhammad-fakhri/go-libs/cache v1.0.0
	github.com/muhammad-fakhri/go-libs/constant v1.0.0
	github.com/muhammad-fakhri/go-libs/email v1.0.0
	github.com/muhammad-fakhri/go-libs/httpmiddleware v1.0.0
	github.com/muhammad-fakhri/go-libs/log v1.0.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/cors v1.8.2
	github.com/stretchr/testify v1.7.1
	go.mongodb.org/mongo-driver v1.3.3
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/api v0.67.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.66.4
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace golang.org/x/net => golang.org/x/net v0.0.0-20210825183410-e898025ed96a
