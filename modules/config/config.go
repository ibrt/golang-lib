package config

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"reflect"

	"github.com/caarlos0/env/v6"
	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
	"github.com/ibrt/golang-lib/validation"
	"gopkg.in/yaml.v2"
)

type contextKey int

const (
	configContextKey contextKey = iota
)

// Stage describes a configuration stage.
type Stage string

// Base configuration stages.
const (
	Test    Stage = "test"
	Staging Stage = "staging"
	Prod    Stage = "prod"
)

// Config describes the common features of a config object.
type Config interface {
	GetStage() Stage
}

// AppliedOptions describes the result of applying options.
type AppliedOptions struct {
	RequireValidation bool
}

// Option applies an option.
type Option func(*AppliedOptions)

// RequireValidation is an Option that requires validation.
func RequireValidation(appliedOptions *AppliedOptions) {
	appliedOptions.RequireValidation = true
}

// EnvInitializerFactory returns an Initializer for the given config type which reads values from the environment.
func EnvInitializerFactory(ctx context.Context, configType reflect.Type, options ...Option) inject.Initializer {
	return func(ctx context.Context) (inject.Injector, inject.Releaser, error) {
		if err := validateConfigType(configType); err != nil {
			return nil, nil, errors.Wrap(err, errors.Skip())
		}

		appliedOptions := applyOptions(options...)

		configValue := reflect.New(configType).Interface()
		if err := env.Parse(&configValue); err != nil {
			return nil, nil, errors.Wrap(err, errors.Prefix("parsing env config"), errors.Skip())
		}

		if err := validateConfigValue(appliedOptions, configValue); err != nil {
			return nil, nil, errors.Wrap(err, errors.Skip())
		}

		return inject.SingletonInjectorFactory(configContextKey, &configValue), nil, nil
	}
}

// JSONFileInitializerFactory returns an Initializer for the given config type which parses the given file as JSON.
func JSONFileInitializerFactory(ctx context.Context, configType reflect.Type, configPath string, options ...Option) inject.Initializer {
	return func(ctx context.Context) (inject.Injector, inject.Releaser, error) {
		if err := validateConfigType(configType); err != nil {
			return nil, nil, errors.Wrap(err, errors.Skip())
		}

		appliedOptions := applyOptions(options...)

		buf, err := ioutil.ReadFile(configPath)
		if err != nil {
			return nil, nil, errors.Wrap(err, errors.Prefix("reading JSON config"), errors.Skip())
		}

		configValue := reflect.New(configType).Interface()
		if err := json.Unmarshal(buf, &configValue); err != nil {
			return nil, nil, errors.Wrap(err, errors.Prefix("parsing JSON config"), errors.Skip())
		}

		if err := validateConfigValue(appliedOptions, configValue); err != nil {
			return nil, nil, errors.Wrap(err, errors.Skip())
		}

		return inject.SingletonInjectorFactory(configContextKey, &configValue), nil, nil
	}
}

// YAMLFileInitializerFactory returns an Initializer for the given config type which parses the given file as Yaml.
func YAMLFileInitializerFactory(ctx context.Context, configType reflect.Type, configPath string, options ...Option) inject.Initializer {
	return func(ctx context.Context) (inject.Injector, inject.Releaser, error) {
		if err := validateConfigType(configType); err != nil {
			return nil, nil, errors.Wrap(err, errors.Skip())
		}

		appliedOptions := applyOptions(options...)

		buf, err := ioutil.ReadFile(configPath)
		if err != nil {
			return nil, nil, errors.Wrap(err, errors.Prefix("reading YAML config"), errors.Skip())
		}

		configValue := reflect.New(configType).Interface()
		if err := yaml.Unmarshal(buf, &configValue); err != nil {
			return nil, nil, errors.Wrap(err, errors.Prefix("parsing YAML config"), errors.Skip())
		}

		if err := validateConfigValue(appliedOptions, configValue); err != nil {
			return nil, nil, errors.Wrap(err, errors.Skip())
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
func MustGet(ctx context.Context) interface{} {
	configValue := Get(ctx)
	errors.Assertf(configValue != nil, "configValue unexpectedly nil", errors.Skip())
	return configValue
}

func validateConfigType(configType reflect.Type) error {
	if configType.Kind() != reflect.Struct {
		return errors.Errorf("configType must be of kind reflect.Struct", errors.Skip())
	}
	if !configType.AssignableTo(reflect.TypeOf((Config)(nil))) {
		return errors.Errorf("configType must implement Config", errors.Skip())
	}
	return nil
}

func applyOptions(options ...Option) *AppliedOptions {
	appliedOptions := &AppliedOptions{}
	for _, option := range options {
		option(appliedOptions)
	}
	return appliedOptions
}

func validateConfigValue(appliedOptions *AppliedOptions, configValue interface{}) error {
	if simpleValidator, ok := configValue.(validation.SimpleValidator); ok {
		if !simpleValidator.Valid() {
			return errors.Errorf("invalid config", errors.Skip())
		}
	}

	if validator, ok := configValue.(validation.Validator); ok {
		if err := validator.Validate(); err != nil {
			return errors.Wrap(err, errors.Prefix("invalid config"), errors.Skip())
		}
	}

	if appliedOptions.RequireValidation {
		return errors.Errorf("validation is required but configType does not implement SimpleValidator or Validator", errors.Skip())
	}

	return nil
}
