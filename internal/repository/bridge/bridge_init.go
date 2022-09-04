package bridge

import (
	"github.com/muhammad-fakhri/archetype-be/internal/component/email"
)

type bridge struct {
	config *bridgeConfig
	// BEGIN __INCLUDE_EMAIL__
	emailClient email.Clienter
	// END __INCLUDE_EMAIL__
}

type bridgeConfig struct {
}

func NewBridge(
	// BEGIN __INCLUDE_EMAIL__
	emailClient email.Clienter,
	// END __INCLUDE_EMAIL__
) Bridge {
	return &bridge{
		// BEGIN __INCLUDE_EMAIL__
		emailClient: emailClient,
		// END __INCLUDE_EMAIL__
		config: &bridgeConfig{},
	}
}
