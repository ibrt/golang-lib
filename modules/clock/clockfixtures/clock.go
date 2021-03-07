package clockfixtures

import (
	"context"
	"testing"
	"time"

	clocklib "github.com/benbjohnson/clock"
	"github.com/ibrt/golang-lib/inject"
	"github.com/ibrt/golang-lib/modules/clock"
)

// Fixtures provides test fixtures for Clock using a real clock.
type Fixtures struct {
	// intentionally empty
}

// BeforeSuite implements fixtures.BeforeSuite.
func (f *Fixtures) BeforeSuite(ctx context.Context) context.Context {
	return inject.MustInject(ctx, clock.SingletonInjectorFactory(clocklib.New()))
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
