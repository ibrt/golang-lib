package clockfixtures

import (
	"context"
	"testing"
	"time"

	clocklib "github.com/benbjohnson/clock"
	"github.com/ibrt/golang-lib/errors"
	"github.com/ibrt/golang-lib/inject"
	"github.com/ibrt/golang-lib/modules/clock"
)

// Fixtures provides test fixtures for Clock using a real clock.
type Fixtures struct {
	releaser inject.Releaser
}

// BeforeSuite implements fixtures.BeforeSuite.
func (f *Fixtures) BeforeSuite(ctx context.Context) context.Context {
	injector, releaser, err := clock.Initializer(ctx)
	errors.MaybeMustWrap(err)
	f.releaser = releaser
	return inject.MustInject(ctx, injector)
}

// AfterSuite implements fixtures.AfterSuite.
func (f *Fixtures) AfterSuite(ctx context.Context) {
	inject.SafeRelease(f.releaser)
}

// MockFixtures provides test fixtures for Clock using a mock.
type MockFixtures struct {
	Clock *clocklib.Mock
}

// BeforeTest implements fixtures.BeforeTest.
func (f *MockFixtures) BeforeTest(ctx context.Context, t *testing.T) context.Context {
	f.Clock = clocklib.NewMock()
	f.Clock.Set(time.Now().UTC())
	return inject.MustInject(ctx, clock.SingletonInjectorFactory(f.Clock))
}

// AfterTest implements fixtures.AfterTest.
func (f *MockFixtures) AfterTest(ctx context.Context) {
	f.Clock = nil
}
