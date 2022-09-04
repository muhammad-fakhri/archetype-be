package bridge

import (
	"context"

	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
	"github.com/muhammad-fakhri/go-libs/email"
)

func (b *bridge) SendEmail(ctx context.Context, mail *email.MailDetail) (err error) {
	err = b.emailClient.Send(mail)
	if err != nil {
		return errors.Wrap(errors.ErrHostEmail, err.Error())
	}

	return
}
