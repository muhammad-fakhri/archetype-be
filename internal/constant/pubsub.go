package constant

const (
	/*
		declare topic and subscription name here
		prefix: {{environment}}.{{app_name}}.{{deadletter-optional}}
		format: {{event}}.topic/subscription
		event e.g:
			claimreward
			submitplay

		custom prefix to differentiate between environment will be added on topic/subscription initialization
		e.g:
			production.archetype-be.claimreward.topic
			production.archetype-be.claimreward.subscription
			production.archetype-be.deadletter.claimreward.topic
			production.archetype-be.deadletter.claimreward.subscription

			test.archetype-be.example.topic
			test.archetype-be.example.subscription
	*/

	/*
		option 1 define event name and then call GetTopic/SubscriptionID when initialize publisher/subscriber
			e.g: (in) example.systemconfig.update --> (out) test.archetype-be.example.systemconfig.update.topic
		option 2 use full topic/subscriptionID from the start, skip GetTopic/SubscriptionID, this option might be needed when subscribing message from external service
	*/
	// BEGIN __INCLUDE_EXAMPLE__
	EventExampleUpdateSystemConfig           string = "example.systemconfig.update" // option 1
	EventExampleDeadletterUpdateSystemConfig string = "example.systemconfig-dl.update"
	// END __INCLUDE_EXAMPLE__
)
