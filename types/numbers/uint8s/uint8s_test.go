package uint8s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/uint8s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := uint8s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint8(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint8s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint8s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint8(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint8s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint8s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint8(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint8(0), uint8s.Val(nil))
	require.Equal(t, uint8(0), uint8s.Val(uint8s.Ptr(0)))
	require.Equal(t, uint8(1), uint8s.Val(uint8s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint8(1), uint8s.ValDef(nil, 1))
	require.Equal(t, uint8(0), uint8s.ValDef(uint8s.Ptr(0), 1))
	require.Equal(t, uint8(1), uint8s.ValDef(uint8s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := uint8s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, uint8(10), v)
	_, err = uint8s.Parse("")
	require.Error(t, err)
	_, err = uint8s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []uint8{2, 0, 3, 1, 4}
	require.False(t, uint8s.Slice(s).IsSorted())
	uint8s.Slice(s).Sort()
	require.Equal(t, []uint8{0, 1, 2, 3, 4}, s)
	require.True(t, uint8s.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[uint8]struct{}{}, uint8s.SliceToMap(nil))
	require.Equal(t, map[uint8]struct{}{}, uint8s.SliceToMap([]uint8{}))
	require.Equal(t, map[uint8]struct{}{1: {}}, uint8s.SliceToMap([]uint8{1}))
	require.Equal(t, map[uint8]struct{}{1: {}, 2: {}}, uint8s.SliceToMap([]uint8{1, 2}))
	require.Equal(t, map[uint8]struct{}{1: {}, 2: {}}, uint8s.SliceToMap([]uint8{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []uint8{}, uint8s.MapToSlice(nil))
	require.Equal(t, []uint8{}, uint8s.MapToSlice(map[uint8]struct{}{}))
	require.Equal(t, []uint8{1}, uint8s.MapToSlice(map[uint8]struct{}{1: {}}))
	require.Equal(t, map[uint8]struct{}{1: {}, 2: {}}, uint8s.SliceToMap(uint8s.MapToSlice(map[uint8]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, uint8(0), uint8s.SafeIndex(nil, 0))
	require.Equal(t, uint8(0), uint8s.SafeIndex(nil, 1))
	require.Equal(t, uint8(0), uint8s.SafeIndex(nil, -1))
	require.Equal(t, uint8(0), uint8s.SafeIndex([]uint8{}, 0))
	require.Equal(t, uint8(0), uint8s.SafeIndex([]uint8{}, 1))
	require.Equal(t, uint8(0), uint8s.SafeIndex([]uint8{}, -1))
	require.Equal(t, uint8(1), uint8s.SafeIndex([]uint8{1}, 0))
	require.Equal(t, uint8(0), uint8s.SafeIndex([]uint8{1}, 1))
	require.Equal(t, uint8(0), uint8s.SafeIndex([]uint8{1}, -1))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, uint8s.SafeIndexPtr(nil, 0))
	require.Nil(t, uint8s.SafeIndexPtr(nil, 1))
	require.Nil(t, uint8s.SafeIndexPtr(nil, -1))
	require.Nil(t, uint8s.SafeIndexPtr([]uint8{}, 0))
	require.Nil(t, uint8s.SafeIndexPtr([]uint8{}, 1))
	require.Nil(t, uint8s.SafeIndexPtr([]uint8{}, -1))
	require.Equal(t, uint8s.Ptr(1), uint8s.SafeIndexPtr([]uint8{1}, 0))
	require.Nil(t, uint8s.SafeIndexPtr([]uint8{1}, 1))
	require.Nil(t, uint8s.SafeIndexPtr([]uint8{1}, -1))
}
