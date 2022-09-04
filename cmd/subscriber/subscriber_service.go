package subscriber

import (
	"github.com/muhammad-fakhri/archetype-be/cmd/subscriber/handler"
	"github.com/muhammad-fakhri/archetype-be/internal/component"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/archetype-be/internal/usecase"
	logger "github.com/muhammad-fakhri/go-libs/log"
)

func Start(task handler.SubscriberTask) (stopFunc func()) {
	// init component
	conf := config.Get()
	logger := logger.NewSLogger(conf.AppName)
	usecase := usecase.InitDependencies(logger)
	client := component.InitPubSubClient()

	// init handler
	param := handler.Init(client, usecase, task)
	subscriber := client.NewSubscriber(logger, param)

	// start subscriber
	go subscriber.Start()

	return func() {
		subscriber.Stop()
		client.Stop()
	}
}
