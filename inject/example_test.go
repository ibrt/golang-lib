package inject_test

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
)

var (
	_ inject.Initializer = ConfigInitializer
	_ inject.Initializer = HTTPClientInitializer
	_ inject.Initializer = RequestIDInitializer
)

type contextKey int

const (
	configContextKey contextKey = iota
	httpClientContextKey
	connectionContextKey
	requestIDContextKey
)

// Config is an imaginary "config" dependency for test purposes.
type Config struct {
	HTTPClientTimeout time.Duration
}

// ConfigInitializer is the initializer for Config.
func ConfigInitializer(ctx context.Context) (inject.Injector, inject.Releaser, error) {
	httpClientTimeout, err := strconv.ParseUint(os.Getenv("GOLANG_LIB_HTTP_CLIENT_TIMEOUT_SECONDS"), 10, 64)
	if err != nil {
		return nil, nil, errors.Wrap(err)
	}

	return inject.SingletonInjector(configContextKey, &Config{
		HTTPClientTimeout: time.Duration(httpClientTimeout) * time.Second,
	}), inject.NoOpReleaser, nil
}

// GetConfig gets the Config from Context.
func GetConfig(ctx context.Context) *Config {
	return ctx.Value(configContextKey).(*Config)
}

// HTTPClientInitializer is the initializer for http.Client.
func HTTPClientInitializer(ctx context.Context) (inject.Injector, inject.Releaser, error) {
	return inject.SingletonInjector(httpClientContextKey, &http.Client{
		Timeout: GetConfig(ctx).HTTPClientTimeout,
	}), inject.NoOpReleaser, nil
}

// GetHTTPClient gets the HTTP client from Context.
func GetHTTPClient(ctx context.Context) *http.Client {
	return ctx.Value(httpClientContextKey).(*http.Client)
}

// Connection is an imaginary dependency that needs to be initialized and eleased.
type Connection struct {
	connected bool
}

// ConnectionInitializer is the initializer for Connection.
func ConnectionInitializer(ctx context.Context) (inject.Injector, inject.Releaser, error) {
	c := &Connection{connected: true}
	return inject.SingletonInjector(connectionContextKey, c), func() { c.connected = false }, nil
}

// GetConnection gets the Connection from Context.
func GetConnection(ctx context.Context) *Connection {
	return ctx.Value(connectionContextKey).(*Connection)
}

// RequestIDInitializer is the initializer for request ids.
func RequestIDInitializer(ctx context.Context) (inject.Injector, inject.Releaser, error) {
	randSrc := rand.New(rand.NewSource(time.Now().UnixNano()))

	return func(ctx context.Context) (context.Context, error) {
		buf := make([]byte, 16)
		if _, err := randSrc.Read(buf); err != nil {
			return nil, errors.Wrap(err)
		}

		return context.WithValue(ctx, requestIDContextKey, fmt.Sprintf("%x", buf)), nil
	}, inject.NoOpReleaser, nil
}

// GetRequestID gets the request ID from Context.
func GetRequestID(ctx context.Context) string {
	return ctx.Value(requestIDContextKey).(string)
}

// RunServer runs an imaginary server.
func RunServer() {
	injector, releaser, err := inject.Initialize(context.Background(),
		ConfigInitializer,
		HTTPClientInitializer,
		ConnectionInitializer,
		RequestIDInitializer)
	defer inject.SafeRelease(releaser)
	errors.MaybeMustWrap(err)

	err = http.ListenAndServe(":3000",
		inject.InjectorMiddlewareFactory(injector)(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(GetRequestID(r.Context())))
				w.WriteHeader(http.StatusOK)
			})))
	if err != http.ErrServerClosed {
		panic(err)
	}
}
