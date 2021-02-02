package config

import (
	"context"
	"os"
	"time"

	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
	"github.com/ibrt/golang-lib/random"
	"github.com/ibrt/golang-lib/types/numbers/int64s"
)

type contextKey int

const (
	configContextKey contextKey = iota
	requestIDContextKey
)

// Config is an example "config" dependency.
type Config struct {
	ClientTimeout time.Duration
}

// Initializer is the inject.Initializer for Config, returning a "singleton" injector.
func Initializer(ctx context.Context) (inject.Injector, inject.Releaser, error) {
	clientTimeoutSeconds, err := int64s.Parse(os.Getenv("GOLANG_LIB_CLIENT_TIMEOUT_SECONDS"))
	if err != nil {
		return nil, nil, errors.Wrap(err)
	}

	injector := inject.SingletonInjector(configContextKey, &Config{
		ClientTimeout: time.Duration(clientTimeoutSeconds) * time.Second,
	})

	return injector, nil, nil
}

// TestProvider returns an Injector and Releaser for use in tests.
func TestProvider(cfg *Config) (inject.Injector, inject.Releaser) {
	return inject.SingletonInjector(configContextKey, cfg), nil
}

// GetConfig gets the Config from Context.
func GetConfig(ctx context.Context) *Config {
	return ctx.Value(configContextKey).(*Config)
}

// RequestIDInitializer is the initializer for request ids, returning a "per-call" injector.
func RequestIDInitializer(ctx context.Context) (inject.Injector, inject.Releaser, error) {
	injector := func(ctx context.Context) (context.Context, error) {
		requestID, err := random.GetHex(16)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		return context.WithValue(ctx, requestIDContextKey, requestID), nil
	}

	return injector, nil, nil
}

// GetRequestID gets the request ID from Context.
func GetRequestID(ctx context.Context) string {
	return ctx.Value(requestIDContextKey).(string)
}
