package inject

import (
	"context"
	"net/http"

	"github.com/ibrt/golang-lib/errors"
)

// Injector injects values into a Context.
type Injector func(ctx context.Context) (context.Context, error)

// SingletonInjector always injects the given (contextKey, value) pair.
func SingletonInjector(contextKey, value interface{}) Injector {
	return func(ctx context.Context) (context.Context, error) {
		return context.WithValue(ctx, contextKey, value), nil
	}
}

// CompoundInjector combines multiple injectors into one.
func CompoundInjector(injectors ...Injector) Injector {
	return func(ctx context.Context) (context.Context, error) {
		for _, injector := range injectors {
			var err error
			if ctx, err = injector(ctx); err != nil {
				return nil, errors.Wrap(err)
			}
		}
		return ctx, nil
	}
}

// InjectorMiddlewareFactory returns a simple HTTP middleware which populates Context using the injector.
// Note: this simple implementation panics if Inject returns error.
func InjectorMiddlewareFactory(injector Injector) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, err := injector(r.Context())
			errors.MaybeMustWrap(err)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
