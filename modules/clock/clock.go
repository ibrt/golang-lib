package clock

import (
	"context"

	"github.com/benbjohnson/clock"
)

type contextKey int

const (
	clockContextKey contextKey = iota
)

// Clock describes a clock.
type Clock clock.Clock

// NewClock initializes a new Clock.
func NewClock() Clock {
	return clock.New()
}

// Provide adds to context.
func Provide(ctx context.Context, clk Clock) context.Context {
	return context.WithValue(ctx, clockContextKey, clk)
}

// Get gets from context.
func Get(ctx context.Context) Clock {
	return ctx.Value(clockContextKey).(Clock)
}
