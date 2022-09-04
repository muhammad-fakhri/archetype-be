package component

import (
	"github.com/muhammad-fakhri/archetype-be/internal/component/email"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
)

func InitEmailClient() email.Clienter {
	conf := config.Get()
	return email.NewGomailClient(conf.Email.Credential)
}
