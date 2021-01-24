package uint16s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/uint16s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := uint16s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint16(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint16s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint16s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint16(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint16s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint16s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint16(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint16(0), uint16s.Val(nil))
	require.Equal(t, uint16(0), uint16s.Val(uint16s.Ptr(0)))
	require.Equal(t, uint16(1), uint16s.Val(uint16s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint16(1), uint16s.ValDef(nil, 1))
	require.Equal(t, uint16(0), uint16s.ValDef(uint16s.Ptr(0), 1))
	require.Equal(t, uint16(1), uint16s.ValDef(uint16s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := uint16s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, uint16(10), v)
	_, err = uint16s.Parse("")
	require.Error(t, err)
	_, err = uint16s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []uint16{2, 0, 3, 1, 4}
	require.False(t, uint16s.Slice(s).IsSorted())
	uint16s.Slice(s).Sort()
	require.Equal(t, []uint16{0, 1, 2, 3, 4}, s)
	require.True(t, uint16s.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[uint16]struct{}{}, uint16s.SliceToMap(nil))
	require.Equal(t, map[uint16]struct{}{}, uint16s.SliceToMap([]uint16{}))
	require.Equal(t, map[uint16]struct{}{1: {}}, uint16s.SliceToMap([]uint16{1}))
	require.Equal(t, map[uint16]struct{}{1: {}, 2: {}}, uint16s.SliceToMap([]uint16{1, 2}))
	require.Equal(t, map[uint16]struct{}{1: {}, 2: {}}, uint16s.SliceToMap([]uint16{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []uint16{}, uint16s.MapToSlice(nil))
	require.Equal(t, []uint16{}, uint16s.MapToSlice(map[uint16]struct{}{}))
	require.Equal(t, []uint16{1}, uint16s.MapToSlice(map[uint16]struct{}{1: {}}))
	require.Equal(t, map[uint16]struct{}{1: {}, 2: {}}, uint16s.SliceToMap(uint16s.MapToSlice(map[uint16]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, uint16(0), uint16s.SafeIndex(nil, 0))
	require.Equal(t, uint16(0), uint16s.SafeIndex(nil, 1))
	require.Equal(t, uint16(0), uint16s.SafeIndex(nil, -1))
	require.Equal(t, uint16(0), uint16s.SafeIndex([]uint16{}, 0))
	require.Equal(t, uint16(0), uint16s.SafeIndex([]uint16{}, 1))
	require.Equal(t, uint16(0), uint16s.SafeIndex([]uint16{}, -1))
	require.Equal(t, uint16(1), uint16s.SafeIndex([]uint16{1}, 0))
	require.Equal(t, uint16(0), uint16s.SafeIndex([]uint16{1}, 1))
	require.Equal(t, uint16(0), uint16s.SafeIndex([]uint16{1}, -1))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, uint16s.SafeIndexPtr(nil, 0))
	require.Nil(t, uint16s.SafeIndexPtr(nil, 1))
	require.Nil(t, uint16s.SafeIndexPtr(nil, -1))
	require.Nil(t, uint16s.SafeIndexPtr([]uint16{}, 0))
	require.Nil(t, uint16s.SafeIndexPtr([]uint16{}, 1))
	require.Nil(t, uint16s.SafeIndexPtr([]uint16{}, -1))
	require.Equal(t, uint16s.Ptr(1), uint16s.SafeIndexPtr([]uint16{1}, 0))
	require.Nil(t, uint16s.SafeIndexPtr([]uint16{1}, 1))
	require.Nil(t, uint16s.SafeIndexPtr([]uint16{1}, -1))
}
