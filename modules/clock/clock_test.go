package clock_test

import (
	"context"
	"testing"
	"time"

	"github.com/ibrt/golang-lib/fixtures"
	"github.com/ibrt/golang-lib/modules/clock"
	"github.com/ibrt/golang-lib/modules/clock/clockfixtures"
	"github.com/stretchr/testify/require"
)

func TestClock(t *testing.T) {
	fixtures.RunSuite(t, &Suite{})
}

type Suite struct {
	Clock *clockfixtures.Fixtures
}

func (s *Suite) TestGet(ctx context.Context, t *testing.T) {
	require.Nil(t, clock.Get(context.Background()))
}

func (s *Suite) TestFixtures(ctx context.Context, t *testing.T) {
	first := clock.MustGet(ctx).Now()
	time.Sleep(100 * time.Millisecond)
	require.NotZero(t, clock.MustGet(ctx).Now().Sub(first))
}

type MockSuite struct {
	Clock *clockfixtures.MockFixtures
}

func (s *MockSuite) TestMockFixtures(ctx context.Context, t *testing.T) {
	v := time.Now().Add(-time.Minute)
	s.Clock.Clock.Set(v)
	require.Equal(t, v, clock.MustGet(ctx).Now())
}
