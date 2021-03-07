package config

import (
	"context"
	"reflect"

	"github.com/caarlos0/env/v6"
	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
	"github.com/ibrt/golang-lib/validation"
)

type contextKey int

const (
	configContextKey contextKey = iota
)

// Config describes a struct that can be used as config.
type Config interface {
	IsConfig()
}

// ConfigMixin can be embedded to implement the Config interface.
type ConfigMixin struct {
	// intentionally empty
}

// IsConfig implements the Config interface.
func (ConfigMixin) IsConfig() {
	// intentionally empty
}

// InitializerFactory returns an Initializer for the given config type which reads values from the environment.
// Refer to https://github.com/caarlos0/env for documentation on environment parsing.
func InitializerFactory(ctx context.Context, configType reflect.Type, parserFuncs map[reflect.Type]env.ParserFunc) inject.Initializer {
	mustValidateConfigType(configType)

	return func(ctx context.Context) (inject.Injector, inject.Releaser, error) {
		configValue := reflect.New(configType).Interface()
		if err := env.ParseWithFuncs(&configValue, parserFuncs); err != nil {
			return nil, nil, errors.Wrap(err, errors.Prefix("parsing env config"), errors.Skip())
		}

		if err := validation.Validate(configValue); err != nil {
			return nil, nil, errors.Wrap(err, errors.Prefix("config"), errors.Skip())
		}

		return inject.SingletonInjectorFactory(configContextKey, &configValue), nil, nil
	}
}

// Get returns the config, or nil if not found.
func Get(ctx context.Context) Config {
	if configValue, ok := ctx.Value(configContextKey).(Config); ok {
		return configValue
	}
	return nil
}

// Get returns the config, panics if not found.
func MustGet(ctx context.Context) Config {
	configValue := Get(ctx)
	errors.Assertf(configValue != nil, "config unexpectedly nil", errors.Skip())
	return configValue
}

func mustValidateConfigType(configType reflect.Type) {
	if configType.Kind() != reflect.Struct {
		errors.MustErrorf("configType must be of kind reflect.Struct", errors.Skip())
	}
	if !configType.AssignableTo(reflect.TypeOf((Config)(nil))) {
		errors.MustErrorf("configType must implement Config", errors.Skip())
	}
}
