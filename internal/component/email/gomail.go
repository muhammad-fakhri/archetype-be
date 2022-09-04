package email

import (
	"github.com/muhammad-fakhri/go-libs/email"
)

type GomailClient struct {
	credential email.Config
}

func NewGomailClient(credential email.Config) Clienter {
	return &GomailClient{
		credential: credential,
	}
}

// send email using open/close method, not suitable for high traffic
func (c *GomailClient) Send(mail *email.MailDetail) error {
	client, err := email.NewEmail(email.Gomail, c.credential)
	if err != nil {
		return err
	}

	defer client.Close()

	return client.Send(mail)
}
