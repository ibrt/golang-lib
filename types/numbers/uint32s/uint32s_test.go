package uint32s_test

import (
	"testing"

	"github.com/ibrt/golang-lib/types/numbers/uint32s"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := uint32s.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint32(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint32s.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint32s.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint32(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint32s.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint32s.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint32(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint32(0), uint32s.Val(nil))
	require.Equal(t, uint32(0), uint32s.Val(uint32s.Ptr(0)))
	require.Equal(t, uint32(1), uint32s.Val(uint32s.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint32(1), uint32s.ValDef(nil, 1))
	require.Equal(t, uint32(0), uint32s.ValDef(uint32s.Ptr(0), 1))
	require.Equal(t, uint32(1), uint32s.ValDef(uint32s.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := uint32s.Parse("10")
	require.NoError(t, err)
	require.Equal(t, uint32(10), v)
	_, err = uint32s.Parse("")
	require.Error(t, err)
	_, err = uint32s.Parse("A")
	require.Error(t, err)
}

func TestSlice(t *testing.T) {
	s := []uint32{2, 0, 3, 1, 4}
	require.False(t, uint32s.Slice(s).IsSorted())
	uint32s.Slice(s).Sort()
	require.Equal(t, []uint32{0, 1, 2, 3, 4}, s)
	require.True(t, uint32s.Slice(s).IsSorted())
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[uint32]struct{}{}, uint32s.SliceToMap(nil))
	require.Equal(t, map[uint32]struct{}{}, uint32s.SliceToMap([]uint32{}))
	require.Equal(t, map[uint32]struct{}{1: {}}, uint32s.SliceToMap([]uint32{1}))
	require.Equal(t, map[uint32]struct{}{1: {}, 2: {}}, uint32s.SliceToMap([]uint32{1, 2}))
	require.Equal(t, map[uint32]struct{}{1: {}, 2: {}}, uint32s.SliceToMap([]uint32{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []uint32{}, uint32s.MapToSlice(nil))
	require.Equal(t, []uint32{}, uint32s.MapToSlice(map[uint32]struct{}{}))
	require.Equal(t, []uint32{1}, uint32s.MapToSlice(map[uint32]struct{}{1: {}}))
	require.Equal(t, map[uint32]struct{}{1: {}, 2: {}}, uint32s.SliceToMap(uint32s.MapToSlice(map[uint32]struct{}{1: {}, 2: {}})))
}
