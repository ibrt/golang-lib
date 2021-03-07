package internal

import (
	"context"

	"github.com/ibrt/golang-lib/inject"
)

// TestInitializer is a helper interface to mock Initializer.
type TestInitializer interface {
	Initialize(ctx context.Context) (inject.Injector, inject.Releaser, error)
}

// TestInjector is a helper interface to mock Injector.
type TestInjector interface {
	Inject(ctx context.Context) (context.Context, error)
}

// TestReleaser is a helper interface to mock Releaser.
type TestReleaser interface {
	Release()
}
