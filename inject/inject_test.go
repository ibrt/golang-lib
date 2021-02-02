package inject_test

/*
import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ibrt/golang-lib/internal/example"
	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
	"github.com/stretchr/testify/require"
)

func TestInject(t *testing.T) {
	injector, releaser, err := inject.Bootstrap(context.Background(),
		example.ConfigInitializer,
		example.HTTPClientInitializer,
		example.ConnectionInitializer,
		example.RequestIDInitializer)
	require.Nil(t, injector)
	require.NotNil(t, releaser)
	require.NotPanics(t, func() { releaser() })
	require.EqualError(t, err, "strconv.ParseUint: parsing \"\": invalid syntax")

	require.NoError(t, os.Setenv("GOLANG_LIB_HTTP_CLIENT_TIMEOUT_SECONDS", "30"))
	defer func() {
		_ = os.Setenv("GOLANG_LIB_HTTP_CLIENT_TIMEOUT_SECONDS", "")
	}()

	injector, releaser, err = inject.Bootstrap(context.Background(),
		example.ConfigInitializer,
		example.HTTPClientInitializer,
		example.ConnectionInitializer,
		example.RequestIDInitializer)
	require.NoError(t, err)

	ctx, err := injector(context.Background())
	require.NoError(t, err)
	require.NotNil(t, example.GetConfig(ctx))
	require.NotNil(t, example.GetHTTPClient(ctx))
	require.NotNil(t, example.GetConnection(ctx))
	require.True(t, example.GetConnection(ctx).connected)
	require.NotNil(t, example.GetRequestID(ctx))

	cfg := example.GetConfig(ctx)
	httpClient := example.GetHTTPClient(ctx)
	c := example.GetConnection(ctx)
	reqID := example.GetRequestID(ctx)

	ctx, err = injector(context.Background())
	require.NoError(t, err)
	require.Equal(t, cfg, example.GetConfig(ctx))
	require.Equal(t, httpClient, example.GetHTTPClient(ctx))
	require.Equal(t, c, example.GetConnection(ctx))
	require.NotEqual(t, reqID, example.GetRequestID(ctx))

	releaser()
	require.False(t, example.GetConnection(ctx).connected)

	injector, _, err = inject.Bootstrap(context.Background(),
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

	injector, _, err = inject.Bootstrap(context.Background(),
		func(ctx context.Context) (inject.Injector, inject.Releaser, error) {
			return func(ctx context.Context) (context.Context, error) {
				return nil, errors.Errorf("test error")
			}, inject.NoOpReleaser, nil
		})
	require.EqualError(t, err, "test error")
}

func TestMiddleware(t *testing.T) {
	injector, _, err := inject.Bootstrap(context.Background(), example.ConnectionInitializer)
	require.NoError(t, err)

	called := false
	inject.InjectorMiddlewareFactory(injector)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.NotNil(t, example.GetConnection(r.Context()))
		called = true
	})).ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "http://localhost", nil))
	require.True(t, called)
}
*/
