package bridge

import (
	"context"

	// BEGIN __INCLUDE_EMAIL__
	"github.com/muhammad-fakhri/go-libs/email"
	// END __INCLUDE_EMAIL__
)

type Bridge interface {
	// BEGIN __INCLUDE_EMAIL__
	SendEmail(ctx context.Context, mail *email.MailDetail) (err error)
	// END __INCLUDE_EMAIL__
}
