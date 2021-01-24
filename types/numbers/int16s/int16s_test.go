package int16s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/int16s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := int16s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int16(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int16s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int16s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int16(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int16s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int16s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int16(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int16(0), int16s.Val(nil))
	require.Equal(t, int16(0), int16s.Val(int16s.Ptr(0)))
	require.Equal(t, int16(1), int16s.Val(int16s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int16(1), int16s.ValDef(nil, 1))
	require.Equal(t, int16(0), int16s.ValDef(int16s.Ptr(0), 1))
	require.Equal(t, int16(1), int16s.ValDef(int16s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := int16s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int16(10), v)
	_, err = int16s.Parse("")
	require.Error(t, err)
	_, err = int16s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []int16{2, 0, 3, 1, 4}
	require.False(t, int16s.Slice(s).IsSorted())
	int16s.Slice(s).Sort()
	require.Equal(t, []int16{0, 1, 2, 3, 4}, s)
	require.True(t, int16s.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[int16]struct{}{}, int16s.SliceToMap(nil))
	require.Equal(t, map[int16]struct{}{}, int16s.SliceToMap([]int16{}))
	require.Equal(t, map[int16]struct{}{1: {}}, int16s.SliceToMap([]int16{1}))
	require.Equal(t, map[int16]struct{}{1: {}, 2: {}}, int16s.SliceToMap([]int16{1, 2}))
	require.Equal(t, map[int16]struct{}{1: {}, 2: {}}, int16s.SliceToMap([]int16{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []int16{}, int16s.MapToSlice(nil))
	require.Equal(t, []int16{}, int16s.MapToSlice(map[int16]struct{}{}))
	require.Equal(t, []int16{1}, int16s.MapToSlice(map[int16]struct{}{1: {}}))
	require.Equal(t, map[int16]struct{}{1: {}, 2: {}}, int16s.SliceToMap(int16s.MapToSlice(map[int16]struct{}{1: {}, 2: {}})))
}

func TestSwapMap(t *testing.T) {
	swap, err := int16s.SwapMap(map[int16]int16{
		1: 2,
		3: 4,
	})
	require.NoError(t, err)
	require.Equal(t,
		map[int16]int16{
			2: 1,
			4: 3,
		}, swap)

	swap, err = int16s.SwapMap(map[int16]int16{})
	require.NoError(t, err)
	require.Equal(t, map[int16]int16{}, swap)

	swap, err = int16s.SwapMap(nil)
	require.NoError(t, err)
	require.Equal(t, map[int16]int16{}, swap)

	_, err = int16s.SwapMap(map[int16]int16{
		1: 3,
		2: 3,
	})
	require.EqualError(t, err, "duplicate value: 3")
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, int16(0), int16s.SafeIndex(nil, 0))
	require.Equal(t, int16(0), int16s.SafeIndex(nil, 1))
	require.Equal(t, int16(0), int16s.SafeIndex(nil, -1))
	require.Equal(t, int16(0), int16s.SafeIndex([]int16{}, 0))
	require.Equal(t, int16(0), int16s.SafeIndex([]int16{}, 1))
	require.Equal(t, int16(0), int16s.SafeIndex([]int16{}, -1))
	require.Equal(t, int16(1), int16s.SafeIndex([]int16{1}, 0))
	require.Equal(t, int16(0), int16s.SafeIndex([]int16{1}, 1))
	require.Equal(t, int16(0), int16s.SafeIndex([]int16{1}, -1))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, int16s.SafeIndexPtr(nil, 0))
	require.Nil(t, int16s.SafeIndexPtr(nil, 1))
	require.Nil(t, int16s.SafeIndexPtr(nil, -1))
	require.Nil(t, int16s.SafeIndexPtr([]int16{}, 0))
	require.Nil(t, int16s.SafeIndexPtr([]int16{}, 1))
	require.Nil(t, int16s.SafeIndexPtr([]int16{}, -1))
	require.Equal(t, int16s.Ptr(1), int16s.SafeIndexPtr([]int16{1}, 0))
	require.Nil(t, int16s.SafeIndexPtr([]int16{1}, 1))
	require.Nil(t, int16s.SafeIndexPtr([]int16{1}, -1))
}
