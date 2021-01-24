package int32s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/int32s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := int32s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int32(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int32s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int32s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int32(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int32s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int32s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int32(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int32(0), int32s.Val(nil))
	require.Equal(t, int32(0), int32s.Val(int32s.Ptr(0)))
	require.Equal(t, int32(1), int32s.Val(int32s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int32(1), int32s.ValDef(nil, 1))
	require.Equal(t, int32(0), int32s.ValDef(int32s.Ptr(0), 1))
	require.Equal(t, int32(1), int32s.ValDef(int32s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := int32s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int32(10), v)
	_, err = int32s.Parse("")
	require.Error(t, err)
	_, err = int32s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []int32{2, 0, 3, 1, 4}
	require.False(t, int32s.Slice(s).IsSorted())
	int32s.Slice(s).Sort()
	require.Equal(t, []int32{0, 1, 2, 3, 4}, s)
	require.True(t, int32s.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[int32]struct{}{}, int32s.SliceToMap(nil))
	require.Equal(t, map[int32]struct{}{}, int32s.SliceToMap([]int32{}))
	require.Equal(t, map[int32]struct{}{1: {}}, int32s.SliceToMap([]int32{1}))
	require.Equal(t, map[int32]struct{}{1: {}, 2: {}}, int32s.SliceToMap([]int32{1, 2}))
	require.Equal(t, map[int32]struct{}{1: {}, 2: {}}, int32s.SliceToMap([]int32{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []int32{}, int32s.MapToSlice(nil))
	require.Equal(t, []int32{}, int32s.MapToSlice(map[int32]struct{}{}))
	require.Equal(t, []int32{1}, int32s.MapToSlice(map[int32]struct{}{1: {}}))
	require.Equal(t, map[int32]struct{}{1: {}, 2: {}}, int32s.SliceToMap(int32s.MapToSlice(map[int32]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, int32(0), int32s.SafeIndex(nil, 0))
	require.Equal(t, int32(0), int32s.SafeIndex(nil, 1))
	require.Equal(t, int32(0), int32s.SafeIndex(nil, -1))
	require.Equal(t, int32(0), int32s.SafeIndex([]int32{}, 0))
	require.Equal(t, int32(0), int32s.SafeIndex([]int32{}, 1))
	require.Equal(t, int32(0), int32s.SafeIndex([]int32{}, -1))
	require.Equal(t, int32(1), int32s.SafeIndex([]int32{1}, 0))
	require.Equal(t, int32(0), int32s.SafeIndex([]int32{1}, 1))
	require.Equal(t, int32(0), int32s.SafeIndex([]int32{1}, -1))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, int32s.SafeIndexPtr(nil, 0))
	require.Nil(t, int32s.SafeIndexPtr(nil, 1))
	require.Nil(t, int32s.SafeIndexPtr(nil, -1))
	require.Nil(t, int32s.SafeIndexPtr([]int32{}, 0))
	require.Nil(t, int32s.SafeIndexPtr([]int32{}, 1))
	require.Nil(t, int32s.SafeIndexPtr([]int32{}, -1))
	require.Equal(t, int32s.Ptr(1), int32s.SafeIndexPtr([]int32{1}, 0))
	require.Nil(t, int32s.SafeIndexPtr([]int32{1}, 1))
	require.Nil(t, int32s.SafeIndexPtr([]int32{1}, -1))
}
