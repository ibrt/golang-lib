package stringz_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/stringz"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := stringz.Ptr("")
	require.NotNil(t, p)
	require.Equal(t, "", *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := stringz.PtrEmptyToNil("")
	require.Nil(t, p)
	p = stringz.PtrEmptyToNil("s")
	require.NotNil(t, p)
	require.Equal(t, "s", *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := stringz.PtrDefToNil("s", "s")
	require.Nil(t, p)
	p = stringz.PtrDefToNil("s", "x")
	require.NotNil(t, p)
	require.Equal(t, "s", *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, "", stringz.Val(nil))
	require.Equal(t, "s", stringz.Val(stringz.Ptr("s")))
}

func TestValDef(t *testing.T) {
	require.Equal(t, "s", stringz.ValDef(nil, "s"))
	require.Equal(t, "x", stringz.ValDef(stringz.Ptr("x"), "s"))
	require.Equal(t, "s", stringz.ValDef(stringz.Ptr("s"), "s"))
}

func TestSlice(t *testing.T) {
	s := []string{"c", "a", "d", "b", "e"}
	require.False(t, stringz.Slice(s).IsSorted())
	stringz.Slice(s).Sort()
	require.Equal(t, []string{"a", "b", "c", "d", "e"}, s)
	require.True(t, stringz.Slice(s).IsSorted())
}
func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[string]struct{}{}, stringz.SliceToMap(nil))
	require.Equal(t, map[string]struct{}{}, stringz.SliceToMap([]string{}))
	require.Equal(t, map[string]struct{}{"1": {}}, stringz.SliceToMap([]string{"1"}))
	require.Equal(t, map[string]struct{}{"1": {}, "2": {}}, stringz.SliceToMap([]string{"1", "2"}))
	require.Equal(t, map[string]struct{}{"1": {}, "2": {}}, stringz.SliceToMap([]string{"1", "1", "2", "2"}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []string{}, stringz.MapToSlice(nil))
	require.Equal(t, []string{}, stringz.MapToSlice(map[string]struct{}{}))
	require.Equal(t, []string{"1"}, stringz.MapToSlice(map[string]struct{}{"1": {}}))
	require.Equal(t, map[string]struct{}{"1": {}, "2": {}}, stringz.SliceToMap(stringz.MapToSlice(map[string]struct{}{"1": {}, "2": {}})))
}
