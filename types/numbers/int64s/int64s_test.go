package int64s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/int64s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := int64s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int64(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int64s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int64s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int64(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int64s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int64s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int64(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int64(0), int64s.Val(nil))
	require.Equal(t, int64(0), int64s.Val(int64s.Ptr(0)))
	require.Equal(t, int64(1), int64s.Val(int64s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int64(1), int64s.ValDef(nil, 1))
	require.Equal(t, int64(0), int64s.ValDef(int64s.Ptr(0), 1))
	require.Equal(t, int64(1), int64s.ValDef(int64s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := int64s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int64(10), v)
	_, err = int64s.Parse("")
	require.Error(t, err)
	_, err = int64s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []int64{2, 0, 3, 1, 4}
	require.False(t, int64s.Slice(s).IsSorted())
	int64s.Slice(s).Sort()
	require.Equal(t, []int64{0, 1, 2, 3, 4}, s)
	require.True(t, int64s.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[int64]struct{}{}, int64s.SliceToMap(nil))
	require.Equal(t, map[int64]struct{}{}, int64s.SliceToMap([]int64{}))
	require.Equal(t, map[int64]struct{}{1: {}}, int64s.SliceToMap([]int64{1}))
	require.Equal(t, map[int64]struct{}{1: {}, 2: {}}, int64s.SliceToMap([]int64{1, 2}))
	require.Equal(t, map[int64]struct{}{1: {}, 2: {}}, int64s.SliceToMap([]int64{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []int64{}, int64s.MapToSlice(nil))
	require.Equal(t, []int64{}, int64s.MapToSlice(map[int64]struct{}{}))
	require.Equal(t, []int64{1}, int64s.MapToSlice(map[int64]struct{}{1: {}}))
	require.Equal(t, map[int64]struct{}{1: {}, 2: {}}, int64s.SliceToMap(int64s.MapToSlice(map[int64]struct{}{1: {}, 2: {}})))
}

func TestSwapMap(t *testing.T) {
	swap, err := int64s.SwapMap(map[int64]int64{
		1: 2,
		3: 4,
	})
	require.NoError(t, err)
	require.Equal(t,
		map[int64]int64{
			2: 1,
			4: 3,
		}, swap)

	swap, err = int64s.SwapMap(map[int64]int64{})
	require.NoError(t, err)
	require.Equal(t, map[int64]int64{}, swap)

	swap, err = int64s.SwapMap(nil)
	require.NoError(t, err)
	require.Equal(t, map[int64]int64{}, swap)

	_, err = int64s.SwapMap(map[int64]int64{
		1: 3,
		2: 3,
	})
	require.EqualError(t, err, "duplicate value: 3")
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, int64(0), int64s.SafeIndex(nil, 0))
	require.Equal(t, int64(0), int64s.SafeIndex(nil, 1))
	require.Equal(t, int64(0), int64s.SafeIndex(nil, -1))
	require.Equal(t, int64(0), int64s.SafeIndex([]int64{}, 0))
	require.Equal(t, int64(0), int64s.SafeIndex([]int64{}, 1))
	require.Equal(t, int64(0), int64s.SafeIndex([]int64{}, -1))
	require.Equal(t, int64(1), int64s.SafeIndex([]int64{1}, 0))
	require.Equal(t, int64(0), int64s.SafeIndex([]int64{1}, 1))
	require.Equal(t, int64(0), int64s.SafeIndex([]int64{1}, -1))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, int64s.SafeIndexPtr(nil, 0))
	require.Nil(t, int64s.SafeIndexPtr(nil, 1))
	require.Nil(t, int64s.SafeIndexPtr(nil, -1))
	require.Nil(t, int64s.SafeIndexPtr([]int64{}, 0))
	require.Nil(t, int64s.SafeIndexPtr([]int64{}, 1))
	require.Nil(t, int64s.SafeIndexPtr([]int64{}, -1))
	require.Equal(t, int64s.Ptr(1), int64s.SafeIndexPtr([]int64{1}, 0))
	require.Nil(t, int64s.SafeIndexPtr([]int64{1}, 1))
	require.Nil(t, int64s.SafeIndexPtr([]int64{1}, -1))
}
