package testutil

import (
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
)

// Common
var (
	DefaultErr = errors.ErrUnknown
)

const (
	EventID        = "1c020ce37bf9756e"
	Country        = "ID"
	AdminAuthToken = "admin_auth_token"
	UserAuthToken  = "user_auth_token"
	AdminCMSToken  = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	UserID         = int64(9876543210)
	AdminID        = "admin@fakhri.com"
)
