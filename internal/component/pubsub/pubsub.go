package pubsub

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type TopicID string
type SubscriptionID string

type PubSubClient struct {
	client *pubsub.Client
	config *PubSubConfig
}

type PubSubConfig struct {
	// common
	CredentialFilepath string
	ProjectID          string
	PrefixID           string

	// publisher config
	Publisher PublisherConfig

	// subscriber config
	Subscriber SubscriberConfig
}

type PublisherConfig struct {
	IsAsync bool
}

type SubscriberConfig struct {
	// common subscription config
	AckDeadline      int
	ExpirationPolicy int

	// deadletter config
	DeadLetterMaxRedelivery int

	// receive concurrency settings
	NumOfGoroutine        int
	MaxOutstandingMessage int
}

func NewPubSubClient(conf *PubSubConfig) (c *PubSubClient) {
	opt := option.WithCredentialsFile(conf.CredentialFilepath)
	client, err := pubsub.NewClient(context.Background(), conf.ProjectID, opt)
	if err != nil {
		log.Fatalf("failed to init pubsub client. err:%v", err)
	}

	return &PubSubClient{
		client: client,
		config: conf,
	}
}

func (c *PubSubClient) Stop() {
	c.client.Close()
}

func (c *PubSubClient) topic(ctx context.Context, id string) (t *pubsub.Topic) {
	t = c.client.Topic(id)
	exists, err := t.Exists(ctx)
	if err != nil {
		log.Fatalf("failed to get topic. err:%v", err)
	}

	if exists {
		return
	}

	log.Printf("create new topic %s", id)
	t, err = c.client.CreateTopic(ctx, id)
	if err != nil {
		log.Fatalf("failed to create topic. err:%v", err)
	}

	return
}

func (c *PubSubClient) subscription(ctx context.Context, id string, topic, dlTopic *pubsub.Topic) (s *pubsub.Subscription) {
	s = c.client.Subscription(id)
	exists, err := s.Exists(ctx)
	if err != nil {
		log.Fatalf("failed to get subscription. err:%v", err)
	}

	if exists {
		return
	}

	var dlPolicy *pubsub.DeadLetterPolicy
	if dlTopic != nil {
		log.Printf("initialize deadletter topic %s", dlTopic.String())
		dlPolicy = &pubsub.DeadLetterPolicy{
			DeadLetterTopic:     dlTopic.String(),
			MaxDeliveryAttempts: c.config.Subscriber.DeadLetterMaxRedelivery,
		}
	}

	log.Printf("create new subscription %s", id)
	s, err = c.client.CreateSubscription(ctx, id, pubsub.SubscriptionConfig{
		Topic:            topic,
		AckDeadline:      time.Duration(c.config.Subscriber.AckDeadline) * time.Second,
		ExpirationPolicy: time.Duration(c.config.Subscriber.ExpirationPolicy), // no expiration
		DeadLetterPolicy: dlPolicy,
	})
	if err != nil {
		log.Fatalf("failed to create subscription. err:%v", err)
	}

	s.ReceiveSettings = pubsub.ReceiveSettings{
		MaxOutstandingMessages: c.config.Subscriber.MaxOutstandingMessage,
		NumGoroutines:          c.config.Subscriber.NumOfGoroutine,
	}

	return
}

func (c *PubSubClient) GetTopicID(baseID string) TopicID {
	return TopicID(fmt.Sprintf("%s.%s.topic", strings.ToLower(c.config.PrefixID), baseID))
}

func (c *PubSubClient) GetSubscriptionID(baseID string) SubscriptionID {
	return SubscriptionID(fmt.Sprintf("%s.%s.subscription", strings.ToLower(c.config.PrefixID), baseID))
}

func (c *PubSubClient) GetDeadLetterTopicID(baseID string) TopicID {
	return c.GetTopicID(fmt.Sprintf("deadletter.%s", baseID))
}

func (c *PubSubClient) GetDeadLetterSubscriptionID(baseID string) SubscriptionID {
	return c.GetSubscriptionID(fmt.Sprintf("deadletter.%s", baseID))
}
