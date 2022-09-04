package config

// BEGIN __INCLUDE_PUBSUB__
import (
	"fmt"

	"github.com/muhammad-fakhri/archetype-be/internal/component/pubsub"
)

type PubSubConfig struct {
	Config *pubsub.PubSubConfig
}

func (c *Config) initPubSubConfig(cfg *configIni) {
	c.PubSub = &PubSubConfig{
		Config: &pubsub.PubSubConfig{
			CredentialFilepath: DefaultConfigPath + cfg.PubSubCredentialFilename,
			ProjectID:          cfg.PubSubProjectID,
			PrefixID:           fmt.Sprintf("%s.%s", cfg.PubSubPrefixID, cfg.AppName),
			Publisher: pubsub.PublisherConfig{
				IsAsync: cfg.PubSubPublishAsync,
			},
			Subscriber: pubsub.SubscriberConfig{
				NumOfGoroutine:          cfg.PubSubNumOfGoroutine,
				MaxOutstandingMessage:   cfg.PubSubMaxOutstandingMessage,
				DeadLetterMaxRedelivery: cfg.PubSubDeadLetterMaxRedelivery,
				AckDeadline:             cfg.PubSubAckDeadline,
				ExpirationPolicy:        cfg.PubSubExpirationPolicy,
			},
		},
	}
}

// END __INCLUDE_PUBSUB__
