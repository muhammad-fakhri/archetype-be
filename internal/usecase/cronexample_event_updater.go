package usecase

// BEGIN __INCLUDE_EXAMPLE_CRON__
import (
	"context"
	"fmt"
)

func (u *usecase) UpdateEventExample(ctx context.Context) (err error) {
	fmt.Println("cron example run")
	return nil
}

// END __INCLUDE_EXAMPLE_CRON__
