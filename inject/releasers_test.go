package inject_test

import (
	"testing"

	"github.com/ibrt/golang-lib/inject"
	"github.com/stretchr/testify/require"
)

type testCloser struct {
	closed bool
}

func (t *testCloser) Close() error {
	t.closed = true
	return nil
}

func TestCloseReleaserFactory(t *testing.T) {
	require.NotPanics(t, func() { inject.CloseReleaserFactory(nil)() })

	c := &testCloser{}
	inject.CloseReleaserFactory(c)()
	require.True(t, c.closed)
}
