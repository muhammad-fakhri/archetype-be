package cron

import (
	"github.com/muhammad-fakhri/archetype-be/cmd/cron/handler"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/archetype-be/internal/usecase"
	logger "github.com/muhammad-fakhri/go-libs/log"
	"github.com/robfig/cron/v3"
)

func Start(task handler.CronTask) (stopFunc func()) {
	// init component
	conf := config.Get()
	logger := logger.NewSLogger(conf.AppName)
	usecase := usecase.InitDependencies(logger)
	client := cron.New(cron.WithChain(
		cron.DelayIfStillRunning(cron.DefaultLogger),
	))

	// init handler
	param := handler.Init(usecase, conf.Scheduler, task)
	client.Schedule(param.Schedule, cron.FuncJob(param.Handler))

	// start cron
	go client.Start()

	return func() {
		client.Stop()
	}
}
