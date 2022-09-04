package testutil

import "database/sql/driver"

/* Prepare database mock */
type AnyNumber struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyNumber) Match(v driver.Value) bool {
	_, ok := v.(int64)
	return ok
}
