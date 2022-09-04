#bin/sh

cd ../

# repository
mockgen -source=internal/repository/contract.go -package mockrepository > internal/test/mockrepository/contract.go

# bridge
mockgen -source=internal/repository/bridge/contract.go -package mockbridge > internal/test/mockbridge/contract.go

# usecase
mockgen -source=internal/usecase/contract.go -package mockusecase > internal/test/mockusecase/contract.go

# component
mockgen -source=internal/component/concurrencylimiter/contract.go -package mockcomponent > internal/test/mockcomponent/concurrencylimiter.go
mockgen -source=internal/component/localcache/contract.go -package mockcomponent > internal/test/mockcomponent/localcache.go
# BEGIN __INCLUDE_REDIS__
mockgen -source=internal/component/lock/contract.go -package mockcomponent > internal/test/mockcomponent/lock.go
mockgen -source=internal/component/ratelimiter/contract.go -package mockcomponent > internal/test/mockcomponent/ratelimiter.go
# END __INCLUDE_REDIS__
# BEGIN __INCLUDE_GCS__
mockgen -source=internal/component/storage/contract.go -package mockcomponent > internal/test/mockcomponent/storage.go
# END __INCLUDE_GCS__
# BEGIN __INCLUDE_PUBSUB__
mockgen -source=internal/component/pubsub/contract.go -package mockcomponent > internal/test/mockcomponent/pubsub.go
# END __INCLUDE_PUBSUB__
# BEGIN __INCLUDE_EMAIL__
mockgen -source=internal/component/email/contract.go -package mockcomponent > internal/test/mockcomponent/email.go
# END __INCLUDE_EMAIL__