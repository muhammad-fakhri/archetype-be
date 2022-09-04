package email

import "github.com/muhammad-fakhri/go-libs/email"

type Clienter interface {
	Send(mail *email.MailDetail) error
}
