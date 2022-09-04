package constant

import (
	"time"

	commonconstant "github.com/muhammad-fakhri/go-libs/constant"
)

const (
	WebGracefulStopTimeout = 15 * time.Second
)

// custom context key
type ContextKey string

const (
	JwtClaimsKey         ContextKey = "jwt_claims"
	AuthorizedAdminIDKey ContextKey = "admin_id" //authorized admin identifier
)

var DomainEnv = map[commonconstant.Environment]string{
	commonconstant.EnvProduction: "",
	commonconstant.EnvStaging:    "uat.",
	commonconstant.EnvTest:       "test.",
}

// BEGIN __INCLUDE_EXAMPLE__
type ExampleMode int

const (
	ExampleSendEmail ExampleMode = iota + 1
	ExamplePublish
)

// END __INCLUDE_EXAMPLE__
