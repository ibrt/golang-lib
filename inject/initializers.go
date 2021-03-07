package inject

import (
	"context"

	"github.com/ibrt/golang-lib/errors"
)

// Initializer initializes a value, returning a corresponding Injector and Releaser.
type Initializer func(ctx context.Context) (Injector, Releaser, error)

// MustInitialize calls the Initializer, panicking on error.
func MustInitialize(ctx context.Context, initializer Initializer) (Injector, Releaser) {
	injector, releaser, err := initializer(ctx)
	errors.MaybeMustWrap(err, errors.Skip())
	return injector, releaser
}

// Bootstrap calls all initializers and returns a compound Injector and Releaser.
// In case of error, the Releasers obtained until then are returned to allow for a clean exit.
func Bootstrap(baseCtx context.Context, initializers ...Initializer) (Injector, Releaser, error) {
	injectors := make([]Injector, 0, len(initializers))
	releasers := make([]Releaser, 0, len(initializers))

	for _, initializer := range initializers {
		injector, releaser, err := initializer(baseCtx)
		if err != nil {
			return nil, CompoundReleaserFactory(releasers...), errors.Wrap(err)
		}

		injectors = append(injectors, injector)
		releasers = append(releasers, releaser)
	}

	return CompoundInjectorFactory(injectors...), CompoundReleaserFactory(releasers...), nil
}
