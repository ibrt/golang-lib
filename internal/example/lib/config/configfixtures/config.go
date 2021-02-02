package configfixtures

import (
	"context"
	"time"

	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
	"github.com/ibrt/golang-lib/internal/example/lib/config"
)

// Fixtures provides test fixtures for config.
type Fixtures struct {
	// intentionally empty
}

// BeforeSuite implements the fixtures.BeforeSuite interface.
func (f *Fixtures) BeforeSuite(ctx context.Context) context.Context {
	cfgInjector, cfgReleaser := config.TestProvider(&config.Config{
		ClientTimeout: 30 * time.Second,
	})
	errors.Assertf(cfgReleaser == nil, "config releaser unexpectedly not nil")
	ctx = inject.MustInject(ctx, cfgInjector)

	reqIDInjector, reqIDReleaser := inject.MustInitialize(ctx, config.RequestIDInitializer)
	errors.Assertf(reqIDReleaser == nil, "request ID releaser unexpectedly not nil")
	ctx = inject.MustInject(ctx, reqIDInjector)

	return ctx
}
