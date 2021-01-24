package ints_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/ints"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := ints.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := ints.PtrZeroToNil(0)
	require.Nil(t, p)
	p = ints.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := ints.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = ints.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int(0), ints.Val(nil))
	require.Equal(t, int(0), ints.Val(ints.Ptr(0)))
	require.Equal(t, int(1), ints.Val(ints.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int(1), ints.ValDef(nil, 1))
	require.Equal(t, int(0), ints.ValDef(ints.Ptr(0), 1))
	require.Equal(t, int(1), ints.ValDef(ints.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := ints.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int(10), v)
	_, err = ints.Parse("")
	require.Error(t, err)
	_, err = ints.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []int{2, 0, 3, 1, 4}
	require.False(t, ints.Slice(s).IsSorted())
	ints.Slice(s).Sort()
	require.Equal(t, []int{0, 1, 2, 3, 4}, s)
	require.True(t, ints.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[int]struct{}{}, ints.SliceToMap(nil))
	require.Equal(t, map[int]struct{}{}, ints.SliceToMap([]int{}))
	require.Equal(t, map[int]struct{}{1: {}}, ints.SliceToMap([]int{1}))
	require.Equal(t, map[int]struct{}{1: {}, 2: {}}, ints.SliceToMap([]int{1, 2}))
	require.Equal(t, map[int]struct{}{1: {}, 2: {}}, ints.SliceToMap([]int{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []int{}, ints.MapToSlice(nil))
	require.Equal(t, []int{}, ints.MapToSlice(map[int]struct{}{}))
	require.Equal(t, []int{1}, ints.MapToSlice(map[int]struct{}{1: {}}))
	require.Equal(t, map[int]struct{}{1: {}, 2: {}}, ints.SliceToMap(ints.MapToSlice(map[int]struct{}{1: {}, 2: {}})))
}

func TestSwapMap(t *testing.T) {
	swap, err := ints.SwapMap(map[int]int{
		1: 2,
		3: 4,
	})
	require.NoError(t, err)
	require.Equal(t,
		map[int]int{
			2: 1,
			4: 3,
		}, swap)

	swap, err = ints.SwapMap(map[int]int{})
	require.NoError(t, err)
	require.Equal(t, map[int]int{}, swap)

	swap, err = ints.SwapMap(nil)
	require.NoError(t, err)
	require.Equal(t, map[int]int{}, swap)

	_, err = ints.SwapMap(map[int]int{
		1: 3,
		2: 3,
	})
	require.EqualError(t, err, "duplicate value: 3")
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, int(0), ints.SafeIndex(nil, 0))
	require.Equal(t, int(0), ints.SafeIndex(nil, 1))
	require.Equal(t, int(0), ints.SafeIndex(nil, -1))
	require.Equal(t, int(0), ints.SafeIndex([]int{}, 0))
	require.Equal(t, int(0), ints.SafeIndex([]int{}, 1))
	require.Equal(t, int(0), ints.SafeIndex([]int{}, -1))
	require.Equal(t, int(1), ints.SafeIndex([]int{1}, 0))
	require.Equal(t, int(0), ints.SafeIndex([]int{1}, 1))
	require.Equal(t, int(0), ints.SafeIndex([]int{1}, -1))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, ints.SafeIndexPtr(nil, 0))
	require.Nil(t, ints.SafeIndexPtr(nil, 1))
	require.Nil(t, ints.SafeIndexPtr(nil, -1))
	require.Nil(t, ints.SafeIndexPtr([]int{}, 0))
	require.Nil(t, ints.SafeIndexPtr([]int{}, 1))
	require.Nil(t, ints.SafeIndexPtr([]int{}, -1))
	require.Equal(t, ints.Ptr(1), ints.SafeIndexPtr([]int{1}, 0))
	require.Nil(t, ints.SafeIndexPtr([]int{1}, 1))
	require.Nil(t, ints.SafeIndexPtr([]int{1}, -1))
}
