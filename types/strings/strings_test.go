package strings_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/strings"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := strings.Ptr("")
	require.NotNil(t, p)
	require.Equal(t, "", *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := strings.PtrEmptyToNil("")
	require.Nil(t, p)
	p = strings.PtrEmptyToNil("s")
	require.NotNil(t, p)
	require.Equal(t, "s", *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := strings.PtrDefToNil("s", "s")
	require.Nil(t, p)
	p = strings.PtrDefToNil("s", "x")
	require.NotNil(t, p)
	require.Equal(t, "s", *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, "", strings.Val(nil))
	require.Equal(t, "s", strings.Val(strings.Ptr("s")))
}

func TestValDef(t *testing.T) {
	require.Equal(t, "s", strings.ValDef(nil, "s"))
	require.Equal(t, "x", strings.ValDef(strings.Ptr("x"), "s"))
	require.Equal(t, "s", strings.ValDef(strings.Ptr("s"), "s"))
}

func TestSlice(t *testing.T) {
	s := []string{"c", "a", "d", "b", "e"}
	require.False(t, strings.Slice(s).IsSorted())
	strings.Slice(s).Sort()
	require.Equal(t, []string{"a", "b", "c", "d", "e"}, s)
	require.True(t, strings.Slice(s).IsSorted())
}
