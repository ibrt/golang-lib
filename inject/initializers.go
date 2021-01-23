package inject

import (
	"context"

	"github.com/ibrt/golang-lib/errors"
)

// Initializer initializes a value, returning a corresponding Injector and Releaser.
type Initializer func(ctx context.Context) (Injector, Releaser, error)

// Initialize calls all initializers and returns a compound Injector and Releaser.
// In case of error, the Releasers obtained until then are returned to allow for a clean exit.
func Initialize(baseCtx context.Context, initializers ...Initializer) (Injector, Releaser, error) {
	injectors := make([]Injector, 0, len(initializers))
	releasers := make([]Releaser, 0, len(initializers))

	for _, initializer := range initializers {
		injector, releaser, err := initializer(baseCtx)
		if err != nil {
			return nil, CompoundReleaser(releasers...), errors.Wrap(err)
		}

		injectors = append(injectors, injector)
		releasers = append(releasers, releaser)

		if baseCtx, err = injector(baseCtx); err != nil {
			return nil, CompoundReleaser(releasers...), errors.Wrap(err)
		}
	}

	return CompoundInjector(injectors...), CompoundReleaser(releasers...), nil
}
