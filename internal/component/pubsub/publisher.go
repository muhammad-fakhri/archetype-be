package pubsub

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
)

type publisher struct {
	topic  *pubsub.Topic
	config *publisherConfig
}

type publisherConfig struct {
	async bool
}

func (c *PubSubClient) NewPublisher(topicID TopicID) Publisher {
	topic := c.topic(context.Background(), string(topicID))
	return &publisher{
		topic: topic,
		config: &publisherConfig{
			async: c.config.Publisher.IsAsync,
		},
	}
}

func (p *publisher) Publish(ctx context.Context, payload interface{}, attributes map[string]string) (err error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	result := p.topic.Publish(ctx, &pubsub.Message{
		Data:       data,
		Attributes: attributes,
	})
	if !p.config.async {
		_, err = result.Get(ctx)
	}

	return
}
