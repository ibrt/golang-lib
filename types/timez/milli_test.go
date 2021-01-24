package timez_test

import (
	"testing"
	"time"

	"github.com/ibrt/golang-lib/types/timez"
	"github.com/stretchr/testify/require"
)

func TestMilli(t *testing.T) {
	v, err := time.Parse(time.RFC3339Nano, "2021-01-24T03:30:48.861422291Z")
	require.NoError(t, err)
	require.Equal(t, int64(1611459048861), timez.UnixMilli(v))
	require.Equal(t, "2021-01-24T03:30:48.861Z", timez.FromUnixMilli(0, 1611459048861).UTC().Format(time.RFC3339Nano))
}
