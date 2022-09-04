package component

import (
	"github.com/muhammad-fakhri/archetype-be/internal/component/pubsub"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
)

func InitPubSubClient() (c *pubsub.PubSubClient) {
	conf := config.Get()
	return pubsub.NewPubSubClient(conf.PubSub.Config)
}
