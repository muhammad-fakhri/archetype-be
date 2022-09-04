package handler

// BEGIN __INCLUDE_EXAMPLE_CRON__
import "context"

type UpdateEventExampleHandler func(ctx context.Context) (err error)

func UpdateEventExample(f UpdateEventExampleHandler) CronHandler {
	return func() {
		ctx := context.Background()
		f(ctx)
	}
}

// END __INCLUDE_EXAMPLE_CRON__
