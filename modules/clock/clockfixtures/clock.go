package clockfixtures

import (
	"context"
	"testing"
	"time"

	mockclock "github.com/benbjohnson/clock"
	"github.com/ibrt/golang-lib/modules/clock"
)

// MockFixtures provides test fixtures for Mail using a mock.
type MockFixtures struct {
	Clock *mockclock.Mock
}

// BeforeTest implements fixtures.BeforeTest.
func (f *MockFixtures) BeforeTest(ctx context.Context, t *testing.T) context.Context {
	f.Clock = mockclock.NewMock()
	f.Clock.Set(time.Now().UTC())
	return clock.Provide(ctx, f.Clock)
}

// AfterTest implements fixtures.AfterTest.
func (f *MockFixtures) AfterTest(ctx context.Context) {
	f.Clock = nil
}
