package uint64s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/uint64s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := uint64s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint64(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint64s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint64s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint64(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint64s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint64s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint64(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint64(0), uint64s.Val(nil))
	require.Equal(t, uint64(0), uint64s.Val(uint64s.Ptr(0)))
	require.Equal(t, uint64(1), uint64s.Val(uint64s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint64(1), uint64s.ValDef(nil, 1))
	require.Equal(t, uint64(0), uint64s.ValDef(uint64s.Ptr(0), 1))
	require.Equal(t, uint64(1), uint64s.ValDef(uint64s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := uint64s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, uint64(10), v)
	_, err = uint64s.Parse("")
	require.Error(t, err)
	_, err = uint64s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []uint64{2, 0, 3, 1, 4}
	require.False(t, uint64s.Slice(s).IsSorted())
	uint64s.Slice(s).Sort()
	require.Equal(t, []uint64{0, 1, 2, 3, 4}, s)
	require.True(t, uint64s.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[uint64]struct{}{}, uint64s.SliceToMap(nil))
	require.Equal(t, map[uint64]struct{}{}, uint64s.SliceToMap([]uint64{}))
	require.Equal(t, map[uint64]struct{}{1: {}}, uint64s.SliceToMap([]uint64{1}))
	require.Equal(t, map[uint64]struct{}{1: {}, 2: {}}, uint64s.SliceToMap([]uint64{1, 2}))
	require.Equal(t, map[uint64]struct{}{1: {}, 2: {}}, uint64s.SliceToMap([]uint64{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []uint64{}, uint64s.MapToSlice(nil))
	require.Equal(t, []uint64{}, uint64s.MapToSlice(map[uint64]struct{}{}))
	require.Equal(t, []uint64{1}, uint64s.MapToSlice(map[uint64]struct{}{1: {}}))
	require.Equal(t, map[uint64]struct{}{1: {}, 2: {}}, uint64s.SliceToMap(uint64s.MapToSlice(map[uint64]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, uint64(0), uint64s.SafeIndex(nil, 0))
	require.Equal(t, uint64(0), uint64s.SafeIndex(nil, 1))
	require.Equal(t, uint64(0), uint64s.SafeIndex(nil, -1))
	require.Equal(t, uint64(0), uint64s.SafeIndex([]uint64{}, 0))
	require.Equal(t, uint64(0), uint64s.SafeIndex([]uint64{}, 1))
	require.Equal(t, uint64(0), uint64s.SafeIndex([]uint64{}, -1))
	require.Equal(t, uint64(1), uint64s.SafeIndex([]uint64{1}, 0))
	require.Equal(t, uint64(0), uint64s.SafeIndex([]uint64{1}, 1))
	require.Equal(t, uint64(0), uint64s.SafeIndex([]uint64{1}, -1))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, uint64s.SafeIndexPtr(nil, 0))
	require.Nil(t, uint64s.SafeIndexPtr(nil, 1))
	require.Nil(t, uint64s.SafeIndexPtr(nil, -1))
	require.Nil(t, uint64s.SafeIndexPtr([]uint64{}, 0))
	require.Nil(t, uint64s.SafeIndexPtr([]uint64{}, 1))
	require.Nil(t, uint64s.SafeIndexPtr([]uint64{}, -1))
	require.Equal(t, uint64s.Ptr(1), uint64s.SafeIndexPtr([]uint64{1}, 0))
	require.Nil(t, uint64s.SafeIndexPtr([]uint64{1}, 1))
	require.Nil(t, uint64s.SafeIndexPtr([]uint64{1}, -1))
}
