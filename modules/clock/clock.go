package clock

import (
	"context"

	clocklib "github.com/benbjohnson/clock"
	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
)

type contextKey int

const (
	clockContextKey contextKey = iota
)

// Clock describes a clock.
type Clock clocklib.Clock

// Initializer is a Clock initializer.
func Initializer(_ context.Context) (inject.Injector, inject.Releaser, error) {
	return inject.SingletonInjectorFactory(clockContextKey, clocklib.New()), nil, nil
}

// SingletonInjectorFactory always injects the given Clock.
func SingletonInjectorFactory(clk Clock) inject.Injector {
	return inject.SingletonInjectorFactory(clockContextKey, clk)
}

// Get returns the Clock, or nil if not found.
func Get(ctx context.Context) Clock {
	if clk, ok := ctx.Value(clockContextKey).(Clock); ok {
		return clk
	}
	return nil
}

// MustGet returns the Clock, panics if not found.
func MustGet(ctx context.Context) Clock {
	clk := Get(ctx)
	errors.Assertf(clk != nil, "clock unexpectedly nil", errors.Skip())
	return clk
}
