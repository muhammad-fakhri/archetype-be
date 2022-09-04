package pubsub

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/muhammad-fakhri/go-libs/log"
)

type SubscriberHandler func(c context.Context, m *pubsub.Message) error
type SubscriberParam struct {
	TopicID                  TopicID
	SubscriptionID           SubscriptionID
	DeadletterTopicID        TopicID        // optional, leave blank if deadletter disabled
	DeadletterSubscriptionID SubscriptionID // optional, leave blank if deadletter disabled or subscription already created
	Handler                  SubscriberHandler
}

type subscriber struct {
	logger       log.SLogger
	id           SubscriptionID
	subscription *pubsub.Subscription
	handler      SubscriberHandler
	stopHandler  func()
}

func (c *PubSubClient) NewSubscriber(logger log.SLogger, param SubscriberParam) Subscriber {
	var (
		topic, dlTopic *pubsub.Topic
		sub            *pubsub.Subscription
	)
	ctx := context.Background()

	if len(param.DeadletterTopicID) > 0 { // initialize deadletter topic if enabled
		dlTopic = c.topic(ctx, string(param.DeadletterTopicID))
		if len(param.DeadletterSubscriptionID) > 0 { // initialize deadletter subscription if enabled
			c.subscription(ctx, string(param.DeadletterSubscriptionID), dlTopic, nil)
		}
	}

	topic = c.topic(ctx, string(param.TopicID))
	sub = c.subscription(ctx, string(param.SubscriptionID), topic, dlTopic)

	return &subscriber{
		subscription: sub,
		id:           param.SubscriptionID,
		handler:      param.Handler,
		logger:       logger,
		stopHandler: func() {
			topic.Stop()
			if dlTopic != nil {
				dlTopic.Stop()
			}
		},
	}
}

func (s *subscriber) Start() (stopFunc func()) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		if err := s.subscription.Receive(ctx, s.receiver); err != nil {
			s.logger.Fatalf(ctx, "failed to receive messages. err:%v", err)
		}
	}(ctx)

	return func() {
		cancel()
	}
}

const (
	valueLogTypeSubscriber = "subscriber"
	statusOK               = 0
	statusError            = 1
)

func (s *subscriber) receiver(ctx context.Context, msg *pubsub.Message) {
	var (
		err       error
		duration  int64 //duration in milliseconds
		startTime time.Time
	)

	defer func() {
		s.log(ctx, startTime.Unix(), duration, err)
	}()

	ctx = context.WithValue(ctx, log.ContextDataMapKey, map[string]string{log.ContextIdKey: msg.ID})
	startTime = time.Now()

	err = s.handler(ctx, msg)
	if err != nil {
		msg.Nack()
		return
	}
	msg.Ack()

	duration = time.Since(startTime).Milliseconds()
}

func (s *subscriber) log(ctx context.Context, startTime, duration int64, err error) {
	status := statusOK
	if err != nil {
		status = statusError
	}

	dataMap := map[string]interface{}{
		log.FieldType:         valueLogTypeSubscriber,
		log.FieldURL:          s.id,
		log.FieldReqTimestamp: startTime,
		log.FieldDurationMs:   duration,
		log.FieldStatus:       status,
	}

	if err == nil {
		s.logger.InfoMap(ctx, dataMap)
	} else {
		s.logger.InfoMap(ctx, dataMap, err)
	}
}

func (s *subscriber) Stop() {
	s.stopHandler()
}
