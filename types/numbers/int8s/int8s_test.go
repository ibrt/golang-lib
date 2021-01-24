package int8s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/int8s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := int8s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int8(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int8s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int8s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int8(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int8s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int8s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int8(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int8(0), int8s.Val(nil))
	require.Equal(t, int8(0), int8s.Val(int8s.Ptr(0)))
	require.Equal(t, int8(1), int8s.Val(int8s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int8(1), int8s.ValDef(nil, 1))
	require.Equal(t, int8(0), int8s.ValDef(int8s.Ptr(0), 1))
	require.Equal(t, int8(1), int8s.ValDef(int8s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := int8s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int8(10), v)
	_, err = int8s.Parse("")
	require.Error(t, err)
	_, err = int8s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []int8{2, 0, 3, 1, 4}
	require.False(t, int8s.Slice(s).IsSorted())
	int8s.Slice(s).Sort()
	require.Equal(t, []int8{0, 1, 2, 3, 4}, s)
	require.True(t, int8s.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[int8]struct{}{}, int8s.SliceToMap(nil))
	require.Equal(t, map[int8]struct{}{}, int8s.SliceToMap([]int8{}))
	require.Equal(t, map[int8]struct{}{1: {}}, int8s.SliceToMap([]int8{1}))
	require.Equal(t, map[int8]struct{}{1: {}, 2: {}}, int8s.SliceToMap([]int8{1, 2}))
	require.Equal(t, map[int8]struct{}{1: {}, 2: {}}, int8s.SliceToMap([]int8{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []int8{}, int8s.MapToSlice(nil))
	require.Equal(t, []int8{}, int8s.MapToSlice(map[int8]struct{}{}))
	require.Equal(t, []int8{1}, int8s.MapToSlice(map[int8]struct{}{1: {}}))
	require.Equal(t, map[int8]struct{}{1: {}, 2: {}}, int8s.SliceToMap(int8s.MapToSlice(map[int8]struct{}{1: {}, 2: {}})))
}
