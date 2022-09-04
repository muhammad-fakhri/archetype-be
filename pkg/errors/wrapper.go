package errors

import "fmt"

func Wrap(err error, message string) error {
	return fmt.Errorf("%w. %s", err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return fmt.Errorf("%w. %s", err, fmt.Sprintf(format, args...))
}
