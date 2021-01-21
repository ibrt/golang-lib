package errors

import "fmt"

// Prefix adds a prefix to the error message.
func Prefix(format string, a ...interface{}) OptionFunc {
	return func(_ bool, err error) {
		if e, ok := err.(*wrappedError); ok {
			e.prefix = fmt.Sprintf(format, a...) + ": " + e.prefix
		}
	}
}
