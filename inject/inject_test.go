package inject_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
	"github.com/stretchr/testify/require"
)

func TestInject(t *testing.T) {
	injector, releaser, err := inject.Initialize(context.Background(),
		ConfigInitializer,
		HTTPClientInitializer,
		ConnectionInitializer,
		RequestIDInitializer)
	require.Nil(t, injector)
	require.NotNil(t, releaser)
	require.NotPanics(t, func() { releaser() })
	require.EqualError(t, err, "strconv.ParseUint: parsing \"\": invalid syntax")

	require.NoError(t, os.Setenv("GOLANG_LIB_HTTP_CLIENT_TIMEOUT_SECONDS", "30"))
	defer func() {
		_ = os.Setenv("GOLANG_LIB_HTTP_CLIENT_TIMEOUT_SECONDS", "")
	}()

	injector, releaser, err = inject.Initialize(context.Background(),
		ConfigInitializer,
		HTTPClientInitializer,
		ConnectionInitializer,
		RequestIDInitializer)
	require.NoError(t, err)

	ctx, err := injector(context.Background())
	require.NoError(t, err)
	require.NotNil(t, GetConfig(ctx))
	require.NotNil(t, GetHTTPClient(ctx))
	require.NotNil(t, GetConnection(ctx))
	require.True(t, GetConnection(ctx).connected)
	require.NotNil(t, GetRequestID(ctx))

	cfg := GetConfig(ctx)
	httpClient := GetHTTPClient(ctx)
	c := GetConnection(ctx)
	reqID := GetRequestID(ctx)

	ctx, err = injector(context.Background())
	require.NoError(t, err)
	require.Equal(t, cfg, GetConfig(ctx))
	require.Equal(t, httpClient, GetHTTPClient(ctx))
	require.Equal(t, c, GetConnection(ctx))
	require.NotEqual(t, reqID, GetRequestID(ctx))

	releaser()
	require.False(t, GetConnection(ctx).connected)

	injector, _, err = inject.Initialize(context.Background(),
		func(ctx context.Context) (inject.Injector, inject.Releaser, error) {
			fail := false

			return func(ctx context.Context) (context.Context, error) {
				if !fail {
					fail = true
					return ctx, nil
				}
				return nil, errors.Errorf("test error")
			}, inject.NoOpReleaser, nil
		})
	require.NoError(t, err)
	_, err = injector(context.Background())
	require.EqualError(t, err, "test error")

	injector, _, err = inject.Initialize(context.Background(),
		func(ctx context.Context) (inject.Injector, inject.Releaser, error) {
			return func(ctx context.Context) (context.Context, error) {
				return nil, errors.Errorf("test error")
			}, inject.NoOpReleaser, nil
		})
	require.EqualError(t, err, "test error")
}

func TestMiddleware(t *testing.T) {
	injector, _, err := inject.Initialize(context.Background(), ConnectionInitializer)
	require.NoError(t, err)

	called := false
	inject.InjectorMiddlewareFactory(injector)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.NotNil(t, GetConnection(r.Context()))
		called = true
	})).ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "http://localhost", nil))
	require.True(t, called)
}
