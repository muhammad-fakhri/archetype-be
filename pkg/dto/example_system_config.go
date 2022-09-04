package dto

// BEGIN __INCLUDE_EXAMPLE__
import (
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	commonconstant "github.com/muhammad-fakhri/go-libs/constant"
)

type PublicSystemConfig struct {
	Configs map[constant.SystemConfig]*SystemConfig `json:"configs"`
}

type PublicSystemConfigRequest struct {
	Configs map[constant.SystemConfig]interface{} `json:"configs"`
	IsAsync bool                                  `json:"is_async"` //to simulate update using pubsub
}

type CountryConfig struct {
	Whitelist []commonconstant.Country `json:"whitelist"` // common config
	// might add more feature specific config here, e.g: maintenance mode
}

type SystemConfig struct {
	Config    interface{} `json:"config"`
	CreatedAt int64       `json:"created_at"`
	UpdatedAt int64       `json:"updated_at"`
}

// END __INCLUDE_EXAMPLE__
