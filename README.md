<!-- BEGIN __DO_NOT_INCLUDE__ -->

# Welcome to archetype-be

This project will jumpstart your Golang project and provide a set of templates you may use to keep your code tidy and consistent. The user provides the required input and thereafter the tool transforms the blueprint project into your own personal project.

## How to use

See https://github.com/muhammad-fakhri/archetype-generator/-/blob/master/README.md

## What's in the box?

The generated project results in:

- `README.md` (which you have to edit to add contents)
- `cmd` as application entry point
- Go modules files (`go.mod` and a generated `go.sum`)
- `internal` package which contains common logic
- gitlab ci.
- manifest and deployment script
- Dockerfile
- and many more...

When creating the project you are asked whether you'd like to include xxx feature support. This will results in different code.

---

## The following part is going to be part of the actual generated project

<!-- This section will be generated for the new project -->
<!-- END __DO_NOT_INCLUDE__ -->

# archetype-be

## Testing

### Create Mock

```
./script/generate_mock.sh
```

### Unit Tests

```
./script/coverage.sh
./script/coverage.sh html
```

## Database

### Create Database

```
CREATE DATABASE {dbname} CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### Migration

```
migrate create -ext sql -dir migrations create_events_table
```

```
migrate -database "mysql://{dbuser}:{dbpass}@tcp(localhost:3306)/{dbname}" -path migrations up
```

To fix dirty migration, run this command

```
migrate -database "mysql://{dbuser}:{dbpass}@tcp(localhost:3306)/{dbname}" -path migrations force {latest_success_id}
```

## Profiling

### Pprof

To load test with wrk, run this command

```
wrk -t3 -c100 -d30s --latency http://localhost:8080/public/v1/ping -H "x-user-id: 5002" -H "x-tenant: id"
```

## How to run

### Command

```
go run main.go
go run main.go --mode=cron-event-example-updater
```

### Config

```
main: /conf/app.ini
pubsub: /conf/credential.json
```
