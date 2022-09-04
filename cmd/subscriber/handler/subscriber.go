package handler

import (
	"log"

	"github.com/muhammad-fakhri/archetype-be/internal/component/pubsub"
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/usecase"
)

type SubscriberTask int

const (
	EmptyTask SubscriberTask = iota
	// BEGIN __INCLUDE_EXAMPLE__
	UpdateSystemConfigTask
	UpdateSystemConfigWithDeadLetterTask
	// END __INCLUDE_EXAMPLE__
)

func Init(client *pubsub.PubSubClient, usecase usecase.Usecase, task SubscriberTask) (param pubsub.SubscriberParam) {
	switch task {
	// BEGIN __INCLUDE_EXAMPLE__
	case UpdateSystemConfigTask:
		// warning: make sure to filter error case specifically to avoid infinite retry loop
		param = pubsub.SubscriberParam{
			TopicID:        client.GetTopicID(constant.EventExampleUpdateSystemConfig),
			SubscriptionID: client.GetSubscriptionID(constant.EventExampleUpdateSystemConfig),
			Handler:        UpdateSystemConfig(usecase.UpdateSystemConfigs),
		}
	case UpdateSystemConfigWithDeadLetterTask:
		param = pubsub.SubscriberParam{
			TopicID:                  client.GetTopicID(constant.EventExampleUpdateSystemConfig),
			SubscriptionID:           client.GetSubscriptionID(constant.EventExampleDeadletterUpdateSystemConfig),
			Handler:                  UpdateSystemConfig(usecase.UpdateSystemConfigs),
			DeadletterTopicID:        client.GetDeadLetterTopicID(constant.EventExampleUpdateSystemConfig),
			DeadletterSubscriptionID: client.GetDeadLetterSubscriptionID(constant.EventExampleDeadletterUpdateSystemConfig),
		}
		// END __INCLUDE_EXAMPLE__
	default:
		log.Fatalf("failed init subscriber. err:invalid task %d", task)
	}

	return param
}
