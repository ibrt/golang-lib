package random_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ibrt/golang-lib/fixtures"
	"github.com/ibrt/golang-lib/modules/random"
	"github.com/ibrt/golang-lib/modules/random/randomfixtures"
	"github.com/stretchr/testify/require"
)

func TestRandom(t *testing.T) {
	fixtures.RunSuite(t, &Suite{})
	fixtures.RunSuite(t, &DeterministicSuite{})
	fixtures.RunSuite(t, &ConfigurableSuite{})
}

type Suite struct {
	Random *randomfixtures.Fixtures
}

func (s *DeterministicSuite) TestSuite(ctx context.Context, t *testing.T) {
	buf := make([]byte, 1024)
	_, err := random.MustGet(ctx).Read(buf)
	fixtures.RequireNoError(t, err)
}

type DeterministicSuite struct {
	Random *randomfixtures.DeterministicFixtures
}

func (s *DeterministicSuite) TestDeterministicSuite(ctx context.Context, t *testing.T) {
	readBuf := randomfixtures.MakeSequence(0, 1024)
	buf := make([]byte, 1024)
	_, err := random.MustGet(ctx).Read(buf)
	fixtures.RequireNoError(t, err)
	require.Equal(t, readBuf, buf)
}

type ConfigurableSuite struct {
	Random *randomfixtures.ConfigurableFixtures
}

func (s *ConfigurableSuite) TestRead(ctx context.Context, t *testing.T) {
	readBuf := randomfixtures.MakeSequence(0, 16)
	ctx = s.Random.Set(ctx, readBuf)

	buf := make([]byte, 16)
	_, err := random.MustGet(ctx).Read(buf)
	fixtures.RequireNoError(t, err)
	require.Equal(t, readBuf, buf)
}

func (s *ConfigurableSuite) TestGetSecureBytes(ctx context.Context, t *testing.T) {
	readBuf := randomfixtures.MakeSequence(0, 16)
	ctx = s.Random.Set(ctx, readBuf)

	require.Empty(t, random.MustGet(ctx).MustGetSecureBytes(0))
	buf := random.MustGet(ctx).MustGetSecureBytes(16)
	require.Len(t, buf, 16)
	require.Equal(t, readBuf[:16], buf)

	buf, err := random.MustGet(ctx).GetSecureBytes(16)
	require.EqualError(t, err, "EOF")
	require.Nil(t, buf)
	require.Panics(t, func() { _ = random.MustGet(ctx).MustGetSecureBytes(16) })
}

func (s *ConfigurableSuite) TestGetHex(ctx context.Context, t *testing.T) {
	readBuf := randomfixtures.MakeSequence(0, 8)
	ctx = s.Random.Set(ctx, readBuf)

	require.Empty(t, random.MustGet(ctx).MustGetHex(0))
	str := random.MustGet(ctx).MustGetHex(16)
	require.Len(t, str, 16)
	require.Equal(t, fmt.Sprintf("%x", readBuf[:8]), str)

	str, err := random.MustGet(ctx).GetHex(16)
	require.EqualError(t, err, "EOF")
	require.Empty(t, str)
	require.Panics(t, func() { _ = random.MustGet(ctx).MustGetHex(16) })
}

func (s *ConfigurableSuite) TestGetAlphabet(ctx context.Context, t *testing.T) {
	readBuf := randomfixtures.MakeSequence(0, len(random.AlphaNum)+len(random.AlphaNum)/4)
	ctx = s.Random.Set(ctx, readBuf)

	require.Empty(t, random.MustGet(ctx).MustGetAlphabet(0, random.AlphaNum))
	str := random.MustGet(ctx).MustGetAlphabet(len(random.AlphaNum), random.AlphaNum)
	require.Len(t, str, len(random.AlphaNum))
	require.Equal(t, string(random.AlphaNum), str)

	str, err := random.MustGet(ctx).GetAlphabet(16, random.AlphaNum)
	require.EqualError(t, err, "EOF")
	require.Empty(t, str)
	require.Panics(t, func() { _ = random.MustGet(ctx).MustGetAlphabet(16, random.AlphaNum) })
	require.Panics(t, func() { random.MustGet(ctx).MustGetAlphabet(16, []rune{}) })
}

func (s *ConfigurableSuite) TestGetID(ctx context.Context, t *testing.T) {
	readBuf := randomfixtures.MakeSequence(0, random.IDLen+random.IDLen/4)
	ctx = s.Random.Set(ctx, readBuf)

	str := random.MustGet(ctx).MustGetID()
	require.Len(t, str, random.IDLen)
	require.Equal(t, string(random.LowerAlphaNum)[:random.IDLen], str)

	str, err := random.MustGet(ctx).GetID()
	require.EqualError(t, err, "EOF")
	require.Empty(t, str)
	require.Panics(t, func() { _ = random.MustGet(ctx).MustGetID() })
}

func (s *ConfigurableSuite) TestGet_Nil(ctx context.Context, t *testing.T) {
	require.Nil(t, random.Get(ctx))
}
