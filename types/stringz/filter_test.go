package stringz_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/stringz"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	eq := stringz.Equal("a")
	require.True(t, eq("a"))
	require.False(t, eq("b"))

	not := stringz.Not(eq)
	require.False(t, not("a"))
	require.True(t, not("b"))

	require.True(t, stringz.And(eq, eq)("a"))
	require.False(t, stringz.And(eq, stringz.Equal("b"))("a"))

	require.True(t, stringz.Or(eq, eq)("a"))
	require.True(t, stringz.Or(eq, stringz.Equal("b"))("a"))
	require.False(t, stringz.Or(eq, eq)("b"))

	require.Equal(t, []string{}, stringz.Filter(nil, eq))
	require.Equal(t, []string{}, stringz.Filter([]string{}, eq))
	require.Equal(t, []string{"a", "a"}, stringz.Filter([]string{"a", "b", "a"}, eq))
}
