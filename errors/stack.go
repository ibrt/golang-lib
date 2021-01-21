package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// GetCallers returns the raw stack trace from the error, or the current raw stack trace if not found.
func GetCallers(err error) []uintptr {
	if e, ok := err.(*wrappedError); ok {
		if e.callers != nil {
			return e.callers
		}
	}

	callers := make([]uintptr, 1024)
	return callers[:runtime.Callers(2, callers[:])]
}

// Skip skips the given number of leading frames in the stack trace (only on first wrap).
func Skip(skip int) OptionFunc {
	return func(firstWrap bool, err error) {
		if e, ok := err.(*wrappedError); ok && firstWrap && e.callers != nil {
			if skip > len(e.callers) {
				skip = len(e.callers)
			}
			e.callers = e.callers[skip:]
		}
	}
}

// FormatStackTrace formats the given raw stack trace.
func FormatStackTrace(callers []uintptr) []string {
	frames := runtime.CallersFrames(callers)
	stackTrace := make([]string, 0, len(callers))

	for {
		frame, more := frames.Next()

		stackTrace = append(
			stackTrace,
			fmt.Sprintf("%v (%v:%v)", filepath.Base(frame.Function), frame.File, frame.Line))

		if !more {
			break
		}
	}

	return stackTrace
}
